package main

import (
	"context"
	"github.com/gharsallahmoez/Basic-API-with-gRPC-and-Protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

// we have to implement all function of proto.AddServiceServer interface to can call RegisterAddServiceServer function.
type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	// create new grpc server
	srv := grpc.NewServer()
	// register the server
	proto.RegisterAddServiceServer(srv, &server{})
	// to serialize and deserialize data
	reflection.Register(srv)
	if err := srv.Serve(listener); err != nil {
		panic(err)
	}

}
func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &proto.Response{Result: result,}, nil

}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &proto.Response{Result: result,}, nil

}
