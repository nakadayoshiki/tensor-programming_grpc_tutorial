package main

import (
	"context"
	"net"

	"github.com/nakadayoshiki/tensor-programming_grpc_tutorial/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func main() {
	l, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(l); e != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, r *proto.Request) (*proto.Response, error) {
	a, b := r.GetA(), r.GetB()
	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, r *proto.Request) (*proto.Response, error) {
	a, b := r.GetA(), r.GetB()
	result := a * b

	return &proto.Response{Result: result}, nil
}
