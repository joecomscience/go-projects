package main

import (
	"context"
	"fmt"
	pb "github.com/joecomscience/go-projects/grpc"
	"google.golang.org/grpc"
	"net"
)

func main() {
	sv := Server{}
	listen, err := net.Listen("tcp", "localhost:3001")

	grpcSv := grpc.NewServer()
	pb.RegisterHelloServer(grpcSv, &sv)

	fmt.Println("Server start")
	if err = grpcSv.Serve(listen); err != nil {
		panic(err)
	}
}

type Server struct{}

func (s *Server) Start(ctx context.Context, p *pb.Msg) (*pb.Msg, error) {
	fmt.Printf("Ping received id: %d, msg: %s\n", p.Id, p.Msg)
	res := pb.Msg{
		Id:  p.Id,
		Msg: p.Msg,
	}
	return &res, nil
}
