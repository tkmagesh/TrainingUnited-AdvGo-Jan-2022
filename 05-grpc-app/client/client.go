package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {

	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect to the server : %v\n", err)
	}
	// Create a new client
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	/* Request & Response */
	//doRequestResponse(ctx, client)

	/* Server Streaming */
	//doServerStreaming(ctx, client)

	/* Client Streaming */
	//doClientStreaming(ctx, client)

	/* Bidirectional Streaming */
	doBidirectionalStreaming(ctx, client)

}

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {
	// Create a new request
	req := &proto.AddRequest{
		X: 10,
		Y: 20,
	}
	res, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}

	// Print the result
	log.Printf("10 + 20 = %d", res.GetResult())
}

func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	start := int32(5)
	end := int32(100)
	req := &proto.PrimeRequest{
		Start: start,
		End:   end,
	}
	stream, err := client.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalf("error while calling GeneratePrimes RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive a prime number: %v", err)
		}
		prime := res.GetPrimeNumber()
		log.Printf("Got a prime number: %v\n", prime)
	}
}

func doClientStreaming(ctx context.Context, client proto.AppServiceClient) {
	data := []int32{5, 2, 6, 1, 7, 4, 9, 8, 3}
	stream, err := client.CalculateAverage(ctx)
	if err != nil {
		log.Fatalf("failed to calculate average: %v", err)
	}
	for _, no := range data {
		fmt.Printf("Sending %d for calculating average\n", no)
		avgReq := &proto.AverageRequest{
			Number: no,
		}
		stream.Send(avgReq)
		time.Sleep(500 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to get average: %v", err)
	}
	log.Printf("Average is: %v\n", res.GetResult())
}

func doBidirectionalStreaming(ctx context.Context, client proto.AppServiceClient) {
	stream, err := client.GreetEveryone(ctx)
	if err != nil {
		log.Fatalf("failed to greet everyone: %v", err)
	}
	done := make(chan bool)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				done <- true
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Printf("Greeting: %v\n", res.GetGreeting())
		}
	}()

	users := []proto.UserName{
		proto.UserName{
			FirstName: "Magesh",
			LastName:  "Kuppan",
		},
		proto.UserName{
			FirstName: "Suresh",
			LastName:  "Kannan",
		},
		proto.UserName{
			FirstName: "Ramesh",
			LastName:  "Jayaraman",
		},
		proto.UserName{
			FirstName: "Rajesh",
			LastName:  "Pandit",
		},
		proto.UserName{
			FirstName: "Ganesh",
			LastName:  "Kumar",
		},
	}

	for _, user := range users {
		log.Printf("Sending user: %v\n", user)

		req := &proto.GreetRequest{
			User: &user,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalln(err)
		}
	}

	<-done
}
