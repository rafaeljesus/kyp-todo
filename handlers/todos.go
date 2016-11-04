package handlers

import (
	"github.com/labstack/echo"
	"github.com/rafaeljesus/kyp-todo/models"
	"net/http"
)

func TodosIndex(c echo.Context) error {
	title := c.QueryParam("title")
	userId := c.QueryParam("user_id")
	query := models.Query{title, userId}

	todos := []models.Todo{}
	if err := models.Search(query, &todos).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, todos)
}

func TodosCreate(c echo.Context) error {
	todo := &models.Todo{}
	if err := c.Bind(todo); err != nil {
		return err
	}

	if err := todo.Create().Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &todo)
}
