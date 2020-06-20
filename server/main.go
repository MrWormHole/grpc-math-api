package main

import (
	"context"
	"net"

	"github.com/MrWormHole/grpc-math-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (server *server) Add(context context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &proto.Response{Result: result}, nil
}

func (server *server) Multiply(context context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &proto.Response{Result: result}, nil
}

func (server *server) Divide(context context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a / b
	return &proto.Response{Result: result}, nil
}

func (server *server) Substract(context context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a - b
	return &proto.Response{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterMathServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
