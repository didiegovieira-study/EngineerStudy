package database

import (
	"context"
	"time"

	"github.com/didiegovieira/EngineerStudy/go-app/application/repositories"
	"github.com/didiegovieira/EngineerStudy/go-app/application/repositories/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDBRepository() repositories.UserRepository {
	return &mongodb.UserMongoDBRepository{}
}

type MongodbDatabase struct {
	Client *mongo.Client
}

func NewMongodb() *MongodbDatabase {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	return &MongodbDatabase{Client: client}
}

func (m MongodbDatabase) GetConnection() any {
	return m.Client
}
