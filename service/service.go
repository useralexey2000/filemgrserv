package service

import (
	"context"
	"filemgrserv/domain"
	"filemgrserv/pb"
	"filemgrserv/semaphore"
)

const dir = "../img"

type FileService struct {
	saveControl semaphore.Semaphore
	listControl semaphore.Semaphore
	mgr         domain.FileManager
	pb.UnimplementedFileServiceServer
}

func NewFileServiceClient(mgr domain.FileManager, maxSaveCon, maxListCon int) *FileService {
	return &FileService{
		saveControl: semaphore.New(maxSaveCon),
		listControl: semaphore.New(maxListCon),
		mgr:         mgr,
	}
}

func (s *FileService) SaveFile(
	ctx context.Context, req *pb.SaveFileRequest) (*pb.SaveFileResponse, error) {
	s.saveControl.Acquire()
	defer s.saveControl.Release()

	err := s.mgr.SaveFile(dir, req.File)
	return &pb.SaveFileResponse{}, err
}

func (s *FileService) ListFiles(
	ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	s.listControl.Acquire()
	defer s.listControl.Release()

	fis, err := s.mgr.ListFilesInfo(dir, int(req.Limit), int(req.Offset))
	if err != nil {
		return &pb.ListFilesResponse{}, err
	}

	return &pb.ListFilesResponse{
		Files: fileInfoListToProto(fis),
	}, nil
}

var _ pb.FileServiceServer = (*FileService)(nil)
