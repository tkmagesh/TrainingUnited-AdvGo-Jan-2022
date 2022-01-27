package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
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
