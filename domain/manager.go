package domain

type FileManager interface {
	SaveFile(dir string, bs []byte) error
	ListFilesInfo(dir string, limit, offset int) ([]*FileInfo, error)
}
