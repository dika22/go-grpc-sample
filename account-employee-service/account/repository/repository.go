package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountRepository interface {

}

type accountRepository struct {
	collection *mongo.Collection
}

func NewAccountRepository(mongodb *mongo.Database) *accountRepository {
	return &accountRepository{
		collection: mongodb.Collection("employees"),
	}
}
