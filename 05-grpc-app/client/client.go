package main

import (
	"context"
	"grpc-app/proto"
	"io"
	"log"

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
	doServerStreaming(ctx, client)

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
