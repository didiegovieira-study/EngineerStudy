package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlDatabase struct {
	Db *sql.DB
}

func NewMysqlDb() *MysqlDatabase {
	//mecanismo de retentativa;;;;fazer
	var db *sql.DB
	var err error
	time.Sleep(10 * time.Second)

	for i := 10; i > 0; i-- {
		time.Sleep(4 * time.Second)
		dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")
		fmt.Println(dsn)

		db, err = sql.Open("mysql", dsn)

		if err != nil {
			fmt.Println("Tentando reconexão")
		} else {
			break
		}
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &MysqlDatabase{Db: db}

}
