package cmd

import (
	"fmt"
	"log"
	"net"
	"os"

	app "go-grpc-sample/account-employee-service/delivery"
	"go-grpc-sample/account-employee-service/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"google.golang.org/grpc"
)

func Execute()  {

	service := app.SetupFeature()

	// Start grpc server
	grpcServer := grpc.NewServer()
	server.SetupHandlerGrpc(grpcServer, service)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_SERVER_PORT")))
	if err != nil {
		log.Fatalf("Failed to listen on port :%s, err: %v", os.Getenv("GRPC_SERVER_PORT"), err)
	}

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server over port :%s, err: %v", os.Getenv("GRPC_SERVER_PORT"), err)
		}
	}()

	// Start http server
	f := fiber.New()

	f.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	f.Use(recover.New())

	f.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	f.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}
