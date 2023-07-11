package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"goapp/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}

	u := models.User{}

	collection := uc.session.Database("mongo-golang").Collection("users")
	result := collection.FindOne(context.Background(), bson.M{"_id": oid})
	if result.Err() != nil {
		log.Println(result.Err().Error())
		w.WriteHeader(404)
		return
	}

	err = result.Decode(&u)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}

	ul, err := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", string(ul))
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = primitive.NewObjectID()

	collection := uc.session.Database("mongo-golang").Collection("users")
	collection.InsertOne(context.TODO(), u)

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}

	collection := uc.session.Database("mongo-golang").Collection("users")
	result := collection.FindOneAndDelete(context.Background(), bson.M{"_id": oid})
	if result.Err() != nil {
		log.Println(result.Err().Error())
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
