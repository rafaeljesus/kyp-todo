package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"github.com/rafaeljesus/kyp-todo/db"
	"github.com/rafaeljesus/kyp-todo/handlers"
	"github.com/rafaeljesus/kyp-todo/models"
	"log"
	"os"
)

func main() {
	db.Connect()
	db.Repo.AutoMigrate(&models.Todo{})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	v1 := e.Group("/v1")
	v1.GET("/healthz", handlers.HealthzIndex)

	v1.Use(middleware.JWT([]byte(os.Getenv("KYP_SECRET_KEY"))))
	v1.GET("/todos", handlers.TodosIndex)
	v1.POST("/todos", handlers.TodosCreate)

	log.Print("Starting Kyp Todo Service...")

	e.Run(fasthttp.New(":" + os.Getenv("KYP_TODO_PORT")))
}
