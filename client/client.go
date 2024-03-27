package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"tp2-grpc-devoir/services"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()
	c := services.NewCalculeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	// Ask the user to enter a number
	fmt.Print("Please enter your name: ")
	if !scanner.Scan() {
		log.Fatal("Failed to read user input")
	}
	nom := scanner.Text()

	service := ""
	for service != "F" && service != "P" {
		// Ask the user to choose a service (Fibonacci or Prime)
		fmt.Print("Please choose a service (F for Fibonacci, P for Prime): ")
		if !scanner.Scan() {
			log.Fatal("Failed to read user input")
		}
		service = scanner.Text()
	}

	// Ask the user to enter a number
	fmt.Print("Please enter a number: ")
	if !scanner.Scan() {
		log.Fatal("Failed to read user input")
	}
	numStr := scanner.Text()

	// Convert the input number to int32
	num, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		log.Fatalf("Failed to parse number: %v", err)
	}

	// Create the gRPC request
	req := &services.Number{
		Nom:     nom,
		Num:     int32(num),
		Service: service,
	}

	// Send the gRPC request
	r, err := c.Send(ctx, req)
	if err != nil {
		log.Fatalf("Error calling gRPC function: %v", err)
	}

	s := r.GetNum()
	if s != -2 {
		// Print the response from the gRPC server
		if service == "P" {
			if s == 0 {
				log.Printf("Response from gRPC server: %d is prime", num)
			} else {
				log.Printf("Response from gRPC server: %d is not prime", num)
			}
		} else {
			log.Printf("Response from gRPC server: fibbonaci of %d is %d", num, r.GetNum())

		}
	} else {
		log.Printf("The maximum number of connections is exceeded!")
	}

}
