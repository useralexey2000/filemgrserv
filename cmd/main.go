package main

import (
	"filemgrserv/filemanager"
	"filemgrserv/pb"
	"filemgrserv/service"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	maxSaveCon = 10
	maxListCon = 100
	host       = "localhost"
	port       = "5555"
)

func main() {
	fmt.Println("serv started")

	lis, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	mgr := filemanager.New(&filemanager.DefaultNameCreator{})
	serv := service.NewFileServiceClient(mgr, maxSaveCon, maxListCon)

	grpcServer := grpc.NewServer()
	pb.RegisterFileServiceServer(grpcServer, serv)
	reflection.Register(grpcServer)
	log.Fatalf("server stopped with err: %v\n", grpcServer.Serve(lis))
}
