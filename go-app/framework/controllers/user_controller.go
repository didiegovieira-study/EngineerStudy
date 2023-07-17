package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/didiegovieira/EngineerStudy/go-app/application/service"
	"github.com/didiegovieira/EngineerStudy/go-app/domain/entities"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(us service.UserService) *UserController {
	return &UserController{userService: &us}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	u, err := uc.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	ul, err := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ul)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := &entities.User{}
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = uc.userService.CreateUser(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(uj)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	err := uc.userService.DeleteUserByID(id)
	if err != nil {
		http.Error(w, id, err.ID)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted user " + id))
}
