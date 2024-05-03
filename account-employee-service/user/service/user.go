package service

import (
	"context"
	"errors"
	"go-grpc-sample/account-employee-service/models"
	pb "go-grpc-sample/account-employee-service/proto/user"
	"go-grpc-sample/account-employee-service/user/repository"
	"time"
)

type UserService interface {
	Create(context.Context, *pb.CreateUserRequest) error
	Update(context.Context, *pb.UpdateUserRequest) error
	List(context.Context) (*pb.ListUserResponse, error)
	Delete(context.Context, *pb.DeleteUserRequest) error
}

type UserRepository struct {
	userRepo repository.UserRepository
}

func NewUsertService(userRepo repository.UserRepository) UserService {
	return &UserRepository{
		userRepo: userRepo,
	}
}

func (r UserRepository) Create(ctx context.Context, req *pb.CreateUserRequest) error {
	
	mappingPayload := models.User{
		Nama: req.Nama,
		Alamat: req.Alamat,
		KodePos: req.KodePos,
		Provinsi: req.Provinsi,
		Kantor: req.Kantor,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := r.userRepo.Create(ctx, mappingPayload)
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepository) Update(ctx context.Context, req *pb.UpdateUserRequest)  error {
	resp, err := r.userRepo.FindByID(ctx, req.User.Id)
	if err != nil {
		return err
	}
	if resp == nil {
		return errors.New("data not found")
	}

	mappingPayload := models.UpdateUser{
		Nama: req.User.Nama,
		Alamat: req.User.Alamat,
		KodePos: req.User.KodePos,
		Provinsi: req.User.Provinsi,
		Kantor: req.User.Kantor,
		UpdatedAt: time.Now(),
	}
	err = r.userRepo.Update(ctx, req.User.Id, mappingPayload)
	if err != nil {
		return err
	}
	return nil
}
	
func (r UserRepository) List(ctx context.Context) (*pb.ListUserResponse, error) {
	resp, err := r.userRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	var users []*pb.User
	var user  pb.User
	for _, v := range resp {
		user.Id = v.ID
		user.Nama = v.Nama
		user.Alamat= v.Alamat
		user.KodePos= v.KodePos
		user.Provinsi= v.Provinsi
		user.Kantor= v.Kantor
		user.CreatedAt= v.CreatedAt.String()
		users = append(users, &user)
	}

	return &pb.ListUserResponse{
		Users: users,
	}, err
}

func (r UserRepository) Delete(ctx context.Context, req *pb.DeleteUserRequest)  error {
	resp, err := r.userRepo.FindByID(ctx, req.Id)
	if err != nil {
		return err
	}
	if resp == nil {
		return errors.New("data not found")
	}
	// Delete
	err = r.userRepo.Delete(ctx, req.Id)
	if err != nil {
		return err
	}
	return nil
}