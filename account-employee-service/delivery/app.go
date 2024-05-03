package delivery

import (
	"context"
	accountRepository "go-grpc-sample/account-employee-service/account/repository"
	accountService "go-grpc-sample/account-employee-service/account/service"
	"go-grpc-sample/account-employee-service/infrastructure/database"
	userRepository "go-grpc-sample/account-employee-service/user/repository"
	userService "go-grpc-sample/account-employee-service/user/service"
	"log"

	"github.com/joho/godotenv"
)
type AccountService struct {
	AccountService accountService.AccountService
	UserService userService.UserService
}


func SetupFeature() AccountService {
	ctx := context.Background()

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	mongoDb := database.InitMongoDB(ctx)

	// init repository
	accountRepository := accountRepository.NewAccountRepository(mongoDb)
	userRepository := userRepository.NewUserRepository(mongoDb)
	// init Service
	accountService := accountService.NewAccountService(accountRepository)
	userService   := userService.NewUsertService(userRepository)
	return AccountService{
		AccountService:  accountService,
		UserService: userService,
	}
}
