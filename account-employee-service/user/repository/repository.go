package repository

import (
	"context"
	"go-grpc-sample/account-employee-service/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(ctx context.Context, param models.User) error
	Update(ctx context.Context, id string, param models.UpdateUser) error
	List(ctx context.Context) ([]*models.User, error)
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*models.User, error)
	FindByKodeCabang(ctx context.Context, kodeCabang string) (*models.User, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(mongodb *mongo.Database) *userRepository {
	return &userRepository{
		collection: mongodb.Collection("users"),
	}
}

func (r userRepository) Create(ctx context.Context, payload models.User) error {
	res, err := r.collection.InsertOne(ctx, payload)
	if err != nil {
		return err
	}
	res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (r userRepository) Update(ctx context.Context, id string, payload models.UpdateUser) error {
	Id, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"$set": payload}

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": Id}, update)
	if err != nil {
		log.Println("error update office", err)
		return err
	}
	return nil


}
	
func (r userRepository) List(ctx context.Context) ([]*models.User, error) {
	var results []*models.User

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}


	return results, nil
}

func (r userRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	Id, _ := primitive.ObjectIDFromHex(id)
	err := r.collection.FindOne(ctx, bson.M{"_id": Id}).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("office not found")
			return nil, err
		}
		log.Println("error find by id", err)
		return nil, err
	}

	return user, nil
}

func (r userRepository) FindByKodeCabang(ctx context.Context, kodeCabang string) (*models.User, error) {
	user := &models.User{}
	err := r.collection.FindOne(ctx, bson.M{"kodeCabang": kodeCabang}).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("kode cabang not found")
			return nil, err
		}
		log.Println("error find by kode cabang", err)
		return nil, err
	}

	return user, nil
}

func (r userRepository) Delete(ctx context.Context, id string) error {
	Id, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": Id})
	if err != nil {
		log.Println("error log delete office", err)
		return err
	}
	return nil

}