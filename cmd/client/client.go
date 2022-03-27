package main

import (
	"fmt"
	"github.com/kelvinramires/grcp_comms/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddUser(client)

}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Name:  "Valter Teste",
		Email: "valter@teste.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not add user: %v", err)
	}

	fmt.Println(res)
}
