package main

import (
	"log"
	"net/http"

	"os"

	"github.com/didiegovieira/EngineerStudy/go-app/application/repositories"
	"github.com/didiegovieira/EngineerStudy/go-app/application/repositories/mongodb"
	"github.com/didiegovieira/EngineerStudy/go-app/application/repositories/mysql"
	"github.com/didiegovieira/EngineerStudy/go-app/application/service"
	"github.com/didiegovieira/EngineerStudy/go-app/framework/controllers"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file loaded")
	}

}

func main() {

	dbType := os.Getenv("DB_TYPE")
	var userRepository repositories.UserRepository
	switch dbType {
	case "mysql":
		// Configuração do MySQL:
		userRepository = mysql.NewMySQLRepository()

	case "mongodb":
		// Configuração do MongoDB:
		userRepository = mongodb.NewMongoDBRepository()

	default:
		return
	}

	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserController(*userService)

	r := httprouter.New()
	r.GET("/user/:id", userController.GetUser)
	r.POST("/user", userController.CreateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	http.ListenAndServe("0.0.0.0:9000", r)
}
