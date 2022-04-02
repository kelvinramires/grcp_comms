package main

import (
	"fmt"
	"github.com/kelvinramires/grcp_comms/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddUserStream(client)
	//AddUsers(client)
	AddUserBiStream(client)
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

func AddUserStream(client pb.UserServiceClient) {
	req := &pb.User{
		Name:  "Valter Teste",
		Email: "valter@teste.com",
	}

	responseStream, err := client.AddUserStream(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not add user: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not add receive stream: %v", err)
		}
		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "01",
			Name:  "Kelvin",
			Email: "kelvinramires@teste.com",
		},
		&pb.User{
			Id:    "02",
			Name:  "Alissar",
			Email: "alissar@teste.com",
		}, &pb.User{
			Id:    "03",
			Name:  "Sopa",
			Email: "sopa@teste.com",
		}, &pb.User{
			Id:    "04",
			Name:  "Papi",
			Email: "papi@teste.com",
		}, &pb.User{
			Id:    "05",
			Name:  "Vanilla",
			Email: "vanilla@teste.com",
		},
	}
	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error Creating Request %v", err)
	}

	for _, req := range reqs {
		err := stream.Send(req)
		if err != nil {
			return
		}
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error Receiving Response %v", err)
	}

	fmt.Println(res)
}

func AddUserBiStream(client pb.UserServiceClient) {
	stream, err := client.AddUserBiStream(context.Background())
	if err != nil {
		log.Fatalf("Error creating request %v", err)
	}
	reqs := []*pb.User{
		&pb.User{
			Id:    "01",
			Name:  "Kelvin",
			Email: "kelvinramires@teste.com",
		},
		&pb.User{
			Id:    "02",
			Name:  "Alissar",
			Email: "alissar@teste.com",
		}, &pb.User{
			Id:    "03",
			Name:  "Sopa",
			Email: "sopa@teste.com",
		}, &pb.User{
			Id:    "04",
			Name:  "Papi",
			Email: "papi@teste.com",
		}, &pb.User{
			Id:    "05",
			Name:  "Vanilla",
			Email: "vanilla@teste.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Error sending stream user: %v", err)
				return
			}
			time.Sleep(time.Second * 2)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("Error closing the stream: %v", err)
			return
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receiving data %v", err)
				break
			}
			fmt.Println("Receiving User:", res.GetUser().GetName(), "with status:", res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
