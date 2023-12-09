package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"io/ioutil"

	"google.golang.org/grpc"

	pb "example/unary/proto"
)

// define the port
const (
	ip_adress = "10.10.12.226"
	port      = ":8080"
)

type helloServer struct {
	pb.UnaryServer
}

func readFromFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (s *helloServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	filename := "data.txt"

	// Read the content of the file into a string
	content, err := os.ReadFile(filename)

	if err != nil {
		fmt.Printf("Error reading from file: %v\n", err)
	}
	//result := "Hello " + req.Msg
	fmt.Printf("Request")

	// log.Printf("Hello " + req.Msg)
	return &pb.HelloResponse{Msg: string(content)}, nil
}

func main() {

	//listen on the port
	address := ip_adress + port
	lis, err := net.Listen("tcp", address)
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
