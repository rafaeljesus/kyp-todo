package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var Repo *gorm.DB

func Connect() {
	conn, err := gorm.Open("postgres", os.Getenv("KYP_TODO_DB"))
	if err != nil {
		panic(err)
	}

	if err := conn.DB().Ping(); err != nil {
		panic(err)
	}

	conn.DB().SetMaxIdleConns(10)
	conn.DB().SetMaxOpenConns(100)

	Repo = conn
}
