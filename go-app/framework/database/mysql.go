package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlDatabase struct {
	Db *sql.DB
}

func NewMysqlDb() *MysqlDatabase {

	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")

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
