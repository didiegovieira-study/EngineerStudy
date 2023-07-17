package main

import (
	"net/http"

	"github.com/didiegovieira/EngineerStudy/go-app/application/service"
	"github.com/didiegovieira/EngineerStudy/go-app/framework/controllers"
	"github.com/didiegovieira/EngineerStudy/go-app/framework/database"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Exemplo de configuração do MySQL:
	//userRepository := database.NewMySQLRepository()

	// Exemplo de configuração do MongoDB:
	userRepository := database.NewMongoDBRepository()

	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserController(*userService)

	r := httprouter.New()
	r.GET("/user/:id", userController.GetUser)
	r.POST("/user", userController.CreateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	http.ListenAndServe("0.0.0.0:9000", r)
}
