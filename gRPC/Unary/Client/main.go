package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	
	"fmt"
	//"io/ioutil"
	"os"
	"time"

	pb "example/unary/proto"
)

// define the port
const (
	ip_adress = "10.10.12.226"
	port = ":8080"
)

func main() {
	// conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Did not connect: %v", err)
	// }
	// defer conn.Close()
	// //object of the client
	// client := pb.NewUnaryClient(conn)

	// // Get the current time
	// currentTime := time.Now()
	// // Format the current time as a string
	// formattedTime := currentTime.Format("2006-01-02 15:04:05.000")
	// // Print the formatted time
	// fmt.Println("Formatted time before sending :", formattedTime)
	// // Create or open the log file
	// logFilePath := "logfile.txt"
	// file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Println("Error opening or creating log file:", err)
	// 	return
	// }
	// defer file.Close()
	// // Write the formatted time to the log file
	// if _, err := file.WriteString(formattedTime before sending + "\n"); err != nil {
	// 	fmt.Println("Error writing to log file:", err)
	// 	return
	// }
	// fmt.Println("Formatted time written to", logFilePath)
	// response, err := client.Hello(context.Background(), &pb.HelloRequest{Msg: "Rohan"})
	// if err != nil {
	// 	log.Fatalf("Error %v", err)
	// }
	// log.Printf("%s", response.Msg)


		address := ip_adress + port

	// Establish a connection to the server
	conn, err := grpc.Dial(address , grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// Create the gRPC client
	client := pb.NewUnaryClient(conn)

	// Get the current time before calling the 'Hello' method
	beforeTime := time.Now()
	// Format the current time as a string
	formattedBeforeTime := beforeTime.Format("2006-01-02 15:04:05.000")
	// Print the formatted time
	fmt.Println("Formatted time before sending:", formattedBeforeTime)

	// Create or open the log file
	logFilePath := "logfile.txt"
	responseFilePath := "response.txt"

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file2,err2 := os.OpenFile(responseFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening or creating log file:", err)
		return
	}
	if err2 != nil {
		fmt.Println("Error opening or creating response file:", err2)
		return
	}


	defer file.Close()
	defer file2.Close()
	// Write the formatted time to the log file
	if _, err := file.WriteString(formattedBeforeTime + " (before sending)\n"); err != nil {
		fmt.Println("Error writing to log file:", err)
		return
	}
	fmt.Println("Formatted time (before sending) written to", logFilePath)

	// Call the 'Hello' method on the server
	response, err := client.Hello(context.Background(), &pb.HelloRequest{Msg: "Rohan"})
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	// log.Printf("%s", response.Msg)
	log.Printf("Hello msg recieved ")
	//print size of response
	// log.Printf("%s", response.sizeof(Msg))

	// Get the current time after calling the 'Hello' method
	afterTime := time.Now()
	// Format the current time as a string
	formattedAfterTime := afterTime.Format("2006-01-02 15:04:05.000")
	// Print the formatted time
	fmt.Println("Formatted time after receiving:", formattedAfterTime)

	// Write the formatted time to the log file
	if _, err := file.WriteString(formattedAfterTime + " (after receiving)\n"); err != nil {
		fmt.Println("Error writing to log file:", err)
		return
	}

	
	fmt.Println("Formatted time (after receiving) written to", responseFilePath)

		// if _, err := file.WriteString(formattedBeforeTime + " (before sending)\n"); err != nil {
		// fmt.Println("Error writing to log file:", err)
		// return
		if _, err2 := file2.WriteString(response.Msg + " (response)\n"); err2 != nil {
		fmt.Println("Error writing to response file:", err2)
		return

	}
}
