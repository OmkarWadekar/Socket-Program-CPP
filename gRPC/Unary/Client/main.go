package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "example/unary/proto"
)

// define the port
const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUnaryClient(conn)

	response, err := client.Hello(context.Background(), &pb.HelloRequest{Msg: "Rohan"})

	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("%s", response.Msg)
}
