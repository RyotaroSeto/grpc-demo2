package main

import (
	"context"
	pb "grpc-demo2/voicer"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myVoicerServer struct {
	pb.UnimplementedInvoicerServer
}

func (s myVoicerServer) Create(ctx context.Context, voicer *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func (s myVoicerServer) Update(ctx context.Context, voicer *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{
		Pdf:  []byte("test2"),
		Docx: []byte("test2"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &myVoicerServer{}
	pb.RegisterInvoicerServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
