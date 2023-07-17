package database

import (
	"database/sql"
	"fmt"

	"github.com/didiegovieira/EngineerStudy/go-app/application/repositories"
	"github.com/didiegovieira/EngineerStudy/go-app/application/repositories/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlDatabase struct {
	Db *sql.DB
}

func NewMySQLRepository() repositories.UserRepository {
	return &mysql.UserMySQLRepository{}
}

func NewMysqlDb() *MysqlDatabase {

	username := "ROOT"
	password := "password"
	host := "localhost"
	port := "33061"
	dbName := "teste"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &MysqlDatabase{Db: db}

}
