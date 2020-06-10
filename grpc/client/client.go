package main


import (
	"context"
	"fmt"
	pb "github.com/joecomscience/go-projects/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func main() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("127.0.0.1:3001", opts...)
	if err != nil {
		panic(err)
	}

	client := pb.NewHelloClient(conn)
	if _, err := sendMsg(client); err != nil {
		panic(err)
	}
	fmt.Println("Finish Pinging")
}

func sendMsg(c pb.HelloClient) (*pb.Msg, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	m := pb.Msg{
		Id:  2,
		Msg: "hello1",
	}

	r, err := c.Start(ctx, &m)
	statusCode := status.Code(err)
	if statusCode != codes.OK {
		return nil, err
	}

	fmt.Printf("Pong: %d, statusCode: %s\n", r.Id, statusCode.String())
	return r, nil
}
