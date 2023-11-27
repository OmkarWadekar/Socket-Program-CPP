package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "example/unary/proto"
)

// define the port
const (
	port = ":8080"
)

type helloServer struct {
	pb.UnaryServer
}

func (s *helloServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	result := "Hello " + req.Msg

	return &pb.HelloResponse{Msg: result}, nil
}

func main() {

	//listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	// create a new gRPC server
	grpcServer := grpc.NewServer()
	// register the greet service
	pb.RegisterUnaryServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())
	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

}
