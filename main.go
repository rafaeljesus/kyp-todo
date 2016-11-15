package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"github.com/rafaeljesus/kyp-todo/handlers"
	"github.com/rafaeljesus/kyp-todo/models"
	"log"
	"os"
)

var KYP_TODO_PORT = os.Getenv("KYP_TODO_PORT")
var KYP_TODO_DB = os.Getenv("KYP_TODO_DB")

func main() {
	db, err := models.NewDB(KYP_TODO_DB)
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&models.Todo{})

	env := &handlers.Env{db}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	v1 := e.Group("/v1")
	v1.GET("/healthz", env.HealthzIndex)
	v1.GET("/todos", env.TodosIndex)
	v1.POST("/todos", env.TodosCreate)

	log.Print("Starting Kyp Todo Service at " + KYP_TODO_PORT)

	e.Run(fasthttp.New(":" + KYP_TODO_PORT))
}
