package filemanager

import (
	"crypto/rand"
	"encoding/hex"
	"filemgrserv/domain"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

type Namer interface {
	CreateName() string
}

type DefaultNameCreator struct{}

func (d *DefaultNameCreator) CreateName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)

	return hex.EncodeToString(randBytes)
}

type DefaultFileManager struct {
	namer Namer
}

func New(namer Namer) *DefaultFileManager {
	return &DefaultFileManager{
		namer: namer,
	}
}

func (d *DefaultFileManager) SaveFile(dir string, bs []byte) error {
	name := d.namer.CreateName()
	ext, err := d.fileExtention(bs)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(dir, fmt.Sprint(name, ".", ext)), bs, 0666)
}

func (d *DefaultFileManager) ListFilesInfo(
	dir string, limit, offset int) ([]*domain.FileInfo, error) {
	fis := make([]*domain.FileInfo, 0)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for i, v := range entries {
		// stop at limit number
		if i == limit {
			break
		}

		// skip offset
		if i < offset {
			continue
		}

		if v.Type().IsDir() {
			continue
		}

		info, err := v.Info()
		if err != nil {
			return nil, err
		}

		fi := parseFileInfo(info)
		fis = append(fis, &fi)
	}

	return fis, nil
}

func parseFileInfo(i fs.FileInfo) domain.FileInfo {
	stat := i.Sys().(*syscall.Stat_t)
	createdAt := timespecToTime(stat.Ctim)
	updatedAt := timespecToTime(stat.Mtim)

	return domain.FileInfo{
		Name:      i.Name(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func (d *DefaultFileManager) fileExtention(bs []byte) (string, error) {
	mime := http.DetectContentType(bs)

	sl := strings.Split(mime, "/")
	if len(sl) != 2 {
		return "", fmt.Errorf("can't parse mime")
	}

	if sl[0] != "image" {
		return "", fmt.Errorf("incorrect type")
	}

	return sl[1], nil
}
