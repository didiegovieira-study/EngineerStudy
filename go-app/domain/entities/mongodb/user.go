package mongodb

import (
	"github.com/didiegovieira/EngineerStudy/go-app/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//type User entities.User

type User struct {
	entities.User
	Id     primitive.ObjectID `json:"id" valid:"notnull" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Gender string             `json:"gender" bson:"gender"`
	Age    int                `json:"age" bson:"age"`
}
