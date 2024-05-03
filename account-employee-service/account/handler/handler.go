package handler

import (
	"go-grpc-sample/account-employee-service/account/service"
)

type AccountHandlerGrpc struct {
	accountService service.AccountService
}

func NewAccountHandlerGrpc(accountService service.AccountService) *AccountHandlerGrpc {
	return &AccountHandlerGrpc{
		accountService: accountService,
	}
}

func Create() error  {
	return nil
}

func Update() error  {
	return nil
}

func List() error  {
	return nil
}

func Delete() error  {
	return nil
}