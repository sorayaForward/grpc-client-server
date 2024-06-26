package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	"tp2-grpc-devoir/services"

	"google.golang.org/grpc"
)

type server struct {
	services.UnimplementedCalculeServer
}

func fibbonacci(num int32) int32 {
	if num <= 1 {
		return num
	}
	return fibbonacci(num-1) + fibbonacci(num-2)
}

func prime(num int32) int32 {
	if num <= 1 {
		return 0
	}
	for i := int32(2); i*i <= num; i++ {
		if num%i == 0 {
			return 0
		}
	}
	return 1
}

var (
	clients        = make(map[string]int)
	maxConnections int // Example maximum number of connections per client
	mutex          sync.Mutex
)

func (s *server) Send(ctx context.Context, in *services.Number) (*services.Result, error) {
	mutex.Lock()
	defer mutex.Unlock()
	clientName := in.Nom

	// Check if the client is already in the map
	if client, ok := clients[clientName]; ok {
		// Client already exists, decrement the number of connections
		client--
		clients[clientName] = client
		if client == 0 {
			return &services.Result{Num: -2}, nil
		}
	} else {
		// Client doesn't exist, add it to the map with connections initialized to max
		clients[clientName] = maxConnections
	}
	log.Println(clients)

	// Process the request
	service := in.Service
	num := in.Num
	var result int32
	if service == "F" {
		result = fibbonacci(num)
	} else if service == "P" {
		result = prime(num)
	}

	return &services.Result{Num: result}, nil
}

func main() {
	// Create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	// Ask the user to enter a number
	fmt.Print("Please enter the maximum of connections for each client : ")
	if !scanner.Scan() {
		log.Fatal("Failed to read user input")
	}
	maxConnStr := scanner.Text()

	// Parse the user input as an integer
	maxConn, _ := strconv.Atoi(maxConnStr)
	maxConnections = maxConn

	// Listen for incoming connections
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server is listening on :50051")

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the server with the generated protobuf service
	services.RegisterCalculeServer(s, &server{})

	// Serve incoming requests in a separate Goroutine
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
