package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rafaeljesus/kyp-structs"
	"github.com/rafaeljesus/kyp-todo/db"
)

type Todo structs.Todo

func (t *Todo) Create() *gorm.DB {
	return db.Repo.Create(t)
}

func Search(q Query, todos *[]Todo) *gorm.DB {
	repo := db.Repo

	if q.Title != "" {
		repo = repo.Where("title = ?", q.Title)
	}

	if q.UserId != "" {
		repo = repo.Where("user_id = ?", q.UserId)
	}

	return repo.Find(todos)
}
