package grpc

import (
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcConn struct {
	*grpc.ClientConn
}

func NewGrpcClient() (client *GrpcConn, err error) {
	addr := fmt.Sprintf("%s:%s", os.Getenv("GRPC_SERVER_HOST"), os.Getenv("GRPC_SERVER_PORT"))

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		err = fmt.Errorf("error when connect to grpc client service '%s', err: %v", os.Getenv("GRPC_SERVER_NAME"), err)
		return
	}

	client = &GrpcConn{
		ClientConn: conn,
	}

	return
}
