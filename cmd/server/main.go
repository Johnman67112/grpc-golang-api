package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Johnman67112/grpc-golang-api/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + in.GetName()}, nil
}

func main() {
	fmt.Println("Runnign gRPC server")

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve grpc: %+v", err)
	}
}
