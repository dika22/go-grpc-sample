package handler

import (
	"context"
	pb "go-grpc-sample/account-employee-service/proto/user"
	"go-grpc-sample/account-employee-service/user/service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserHandlerGrpc struct {
	userService service.UserService
	pb.UserServiceServer
}

func NewUserHandlerGrpc(userService service.UserService) *UserHandlerGrpc {
	return &UserHandlerGrpc{
		userService: userService,
	}
}

func (h UserHandlerGrpc) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error)  {
	err := h.userService.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Code: 201,
		Message: "success",
	}, nil
}

func (h UserHandlerGrpc) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error)  {
	err := h.userService.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Code: 200,
		Message: "success",
	}, nil
}

func (h UserHandlerGrpc) List(ctx context.Context, req *emptypb.Empty) (*pb.ListUserResponse, error)  {
	resp, err := h.userService.List(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h UserHandlerGrpc) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)  {
	err := h.userService.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{
		Id: req.Id,
		Code: 200,
		Message: "success",
	}, nil
}
