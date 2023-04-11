package service

import (
	"filemgrserv/domain"
	"filemgrserv/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func fileInfoToProto(fi *domain.FileInfo) *pb.FileInfo {
	return &pb.FileInfo{
		Name:      fi.Name,
		CreatedAt: timestamppb.New(fi.CreatedAt),
		UpdatedAt: timestamppb.New(fi.UpdatedAt),
	}
}

func fileInfoListToProto(fis []*domain.FileInfo) []*pb.FileInfo {
	pbfis := make([]*pb.FileInfo, 0, len(fis))
	for _, v := range fis {
		fi := fileInfoToProto(v)
		pbfis = append(pbfis, fi)
	}

	return pbfis
}
