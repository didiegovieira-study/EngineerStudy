package mongodb

import (
	"context"
	"fmt"

	"github.com/didiegovieira/EngineerStudy/go-app/domain/entities/mongodb"
	"github.com/didiegovieira/EngineerStudy/go-app/framework/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserMongoDBRepository struct {
	Db   *database.MongodbDatabase
	user mongodb.User
}

func (repository UserMongoDBRepository) getCollection() *mongo.Collection {
	return repository.Db.Client.Database("mongo-golang").Collection("users")
}

func (repository UserMongoDBRepository) CreateUser(user mongodb.User) error {
	u := user
	_, err := repository.getCollection().InsertOne(context.Background(), u)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}

func (repository UserMongoDBRepository) DeleteUserByID(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid user ID")
	}

	filter := bson.M{"_id": oid}
	result, err := repository.getCollection().DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (repository UserMongoDBRepository) GetByID(id string) (mongodb.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return mongodb.User{}, fmt.Errorf("invalid user ID")
	}

	filter := bson.M{"_id": oid}
	var user mongodb.User

	err = repository.getCollection().FindOne(context.Background(), filter).Decode(&user)

	return user, err
}
