package main

import (
	"context"
	"log"
	"net"

	pb "hello-kubernetes/proto"

	"google.golang.org/grpc"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recovered form panic")
		}
	}()
	lis, err := net.Listen("tcp", ":50076")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	helloServer := NewHelloServer()
	pb.RegisterHelloServiceServer(grpcServer, helloServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

type helloServer struct {
	pb.UnimplementedHelloServiceServer
}

func NewHelloServer() *helloServer {
	return &helloServer{}
}
func (h *helloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "HI " + req.Name}, nil

}
