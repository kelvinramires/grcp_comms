package services

import (
	"fmt"
	"github.com/kelvinramires/grcp_comms/pb"
	"golang.org/x/net/context"
	"io"
	"log"
	"time"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*UserService) AddUserStream(req *pb.User, stream pb.UserService_AddUserStreamServer) error {
	err := stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	if err != nil {
		return err
	}
	time.Sleep(time.Second * 3)

	err = stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})

	if err != nil {
		return err
	}
	time.Sleep(time.Second * 3)

	err = stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User:   &pb.User{Name: req.GetName(), Email: req.GetEmail()},
	})

	if err != nil {
		return err
	}

	time.Sleep(time.Second * 3)

	err = stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User:   &pb.User{Name: req.GetName(), Email: req.GetEmail()},
	})

	if err != nil {
		return err
	}
	time.Sleep(time.Second * 3)

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	var users []*pb.User

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}
		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})
		fmt.Println("Adding", req.GetName())
	}
}

func (*UserService) AddUserBiStream(stream pb.UserService_AddUserBiStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Ok",
			User:   req,
		})

		if err != nil {
			log.Fatalf("Error sending stream to client: %v", err)
		}

	}
}
