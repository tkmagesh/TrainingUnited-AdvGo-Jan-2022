package main

import (
	"context"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {

	// Create a new request
	req := &proto.AddRequest{
		X: 10,
		Y: 20,
	}

	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect to the server : %v\n", err)
	}
	// Create a new client
	client := proto.NewAppServiceClient(clientConn)

	// Call the server
	res, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}

	// Print the result
	log.Printf("10 + 20 = %d", res.GetResult())

}
