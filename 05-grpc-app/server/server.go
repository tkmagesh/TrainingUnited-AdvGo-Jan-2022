package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAppServiceServer
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	fmt.Printf("Serving Add request for %d and %d with result %d\n", x, y, result)
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func (s *server) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	for no := start; no <= end; no++ {
		if isPrime(no) {
			fmt.Printf("Responding with prime %d\n", no)
			res := &proto.PrimeResponse{
				PrimeNumber: no,
			}
			err := serverStream.Send(res)
			if err != nil {
				log.Fatalln(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func (s *server) CalculateAverage(stream proto.AppService_CalculateAverageServer) error {
	var sum int32
	var count int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Received %d for calculating average\n", req.GetNumber())
		sum += req.GetNumber()
		count++
	}
	avg := sum / count
	res := &proto.AverageResponse{
		Result: avg,
	}
	fmt.Printf("Sending average : %d\n", avg)
	return stream.SendAndClose(res)

}

func isPrime(no int32) bool {
	if no <= 0 {
		return false
	}
	if no <= 3 {
		return true
	}
	var i int32
	for i = 2; i < no-1; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true

}

func (s *server) GreetEveryone(stream proto.AppService_GreetEveryoneServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		user := req.GetUser()
		firstName := user.GetFirstName()
		lastName := user.GetLastName()
		fmt.Printf("Received req for greeting for %s and %s\n", firstName, lastName)
		res := &proto.GreetResponse{
			Greeting: fmt.Sprintf("Hello %s %s, Have a nice day!", firstName, lastName),
		}
		err = stream.Send(res)
		if err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func main() {
	s := &server{}
	listner, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	proto.RegisterAppServiceServer(server, s)
	server.Serve(listner)
}
