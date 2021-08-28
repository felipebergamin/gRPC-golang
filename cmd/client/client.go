package main

import (
	"context"
	"fmt"
	"github.com/felipebergamin/gRPC-golang/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddUserVerbose(client)
	AddUsers(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "21",
		Name:  "Felipe Bergamin",
		Email: "f@f.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	log.Println(res)
}

func AddUserVerbose (client pb.UserServiceClient) {
	req := &pb.User{
		Id: "0",
		Name: "Felipe",
		Email: "f@f.com",
	}
	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}
		fmt.Println("Status:", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "Felipe",
			Email: "f@f.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Felipe 2",
			Email: "f2@f.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Felipe 3",
			Email: "f3@f.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Felipe 4",
			Email: "f4@f.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}
