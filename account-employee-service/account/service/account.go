package service

import (
	"go-grpc-sample/account-employee-service/account/repository"
)

type AccountService interface {
}

type accountRepository struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(acc repository.AccountRepository) AccountService {
	return &accountRepository{
		accountRepo: acc,
	}
}