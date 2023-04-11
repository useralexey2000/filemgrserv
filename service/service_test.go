package service

import (
	"context"
	"errors"
	"filemgrserv/domain"
	"filemgrserv/filemanager"
	"filemgrserv/pb"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

type fakeNamer struct{}

func (n *fakeNamer) CreateName() string {
	return "fake"
}

type fakeManager struct {
	counter int
	wg      *sync.WaitGroup
}

func (m *fakeManager) SaveFile(dir string, bs []byte) error {
	m.wg.Add(1)
	defer m.wg.Done()

	m.counter++
	time.Sleep(1 * time.Millisecond)
	m.counter--
	return nil
}

func (m *fakeManager) ListFilesInfo(dir string, limit, offset int) ([]*domain.FileInfo, error) {
	m.wg.Add(1)
	defer m.wg.Done()

	m.counter++
	time.Sleep(1 * time.Millisecond)
	m.counter--
	return []*domain.FileInfo{}, nil
}

func TestSaveFile(t *testing.T) {
	bs, err := os.ReadFile("../hamster.png")
	if err != nil {
		panic(err)
	}

	s := NewFileServiceClient(filemanager.New(&fakeNamer{}), 1, 1)

	_, err = s.SaveFile(context.Background(), &pb.SaveFileRequest{
		File: bs,
	})

	if err != nil {
		t.Errorf("can't save file error %v", err)
	}

	if _, err := os.Stat(dir + "/fake.png"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("no file found error %v", err)
	}

	os.Remove(dir + "/fake.png")
}

func TestListFiles(t *testing.T) {
	for i := 0; i < 3; i++ {
		_, err := os.Create(fmt.Sprint(dir, "/", i, ".png"))
		if err != nil {
			panic(err)
		}
	}
	defer func() {
		for i := 0; i < 3; i++ {
			err := os.Remove(fmt.Sprint(dir, "/", i, ".png"))
			if err != nil {
				panic(err)
			}
		}
	}()

	s := NewFileServiceClient(filemanager.New(&fakeNamer{}), 1, 1)

	res, err := s.ListFiles(context.Background(), &pb.ListFilesRequest{
		Offset: 0,
		Limit:  3,
	})

	if err != nil {
		t.Errorf("can't list files %v", err)
	}

	if len(res.Files) != 3 {
		t.Errorf("wrong list len, want %v, got %v", 3, len(res.Files))
	}
}

func TestSaveFile_limit(t *testing.T) {
	done := make(chan struct{})
	maxconn := 10

	wg := sync.WaitGroup{}
	fkmanager := &fakeManager{wg: &wg}
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				if fkmanager.counter > maxconn {
					t.Errorf("exceeded number of parallel connection")
				}
			}
		}
	}()

	s := NewFileServiceClient(fkmanager, maxconn, 0)

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()

			s.SaveFile(context.Background(), &pb.SaveFileRequest{
				File: []byte{},
			})
		})
	}

	wg.Wait()
	done <- struct{}{}
}

func TestListFiles_limit(t *testing.T) {
	done := make(chan struct{})
	maxconn := 100

	wg := sync.WaitGroup{}
	fkmanager := &fakeManager{wg: &wg}
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				if fkmanager.counter > maxconn {
					t.Errorf("exceeded number of parallel connection")
				}
			}
		}
	}()

	s := NewFileServiceClient(fkmanager, 0, maxconn)

	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()

			s.ListFiles(context.Background(), &pb.ListFilesRequest{Offset: 0, Limit: 0})
		})
	}

	wg.Wait()
	done <- struct{}{}
}
