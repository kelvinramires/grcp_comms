package services

import (
	"github.com/kelvinramires/grcp_comms/pb"
	"golang.org/x/net/context"
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
