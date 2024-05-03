package server

import (
	"go-grpc-sample/account-employee-service/account/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	app "go-grpc-sample/account-employee-service/delivery"
	userGrpcService "go-grpc-sample/account-employee-service/proto/user"
	userHandler "go-grpc-sample/account-employee-service/user/handler"
)



type HandlerGrpc struct {
	accountHandler handler.AccountHandlerGrpc
	userHandler    userHandler.UserHandlerGrpc
}

func SetupHandlerGrpc(grpcServer *grpc.Server, app app.AccountService) {

	handler := HandlerGrpc{
		accountHandler: *handler.NewAccountHandlerGrpc(app.AccountService),
		userHandler: *userHandler.NewUserHandlerGrpc(app.UserService),
	}

	userGrpcService.RegisterUserServiceServer(grpcServer, &handler.userHandler)
	reflection.Register(grpcServer)
}
