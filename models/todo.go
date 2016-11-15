package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rafaeljesus/kyp-structs"
)

type Todo structs.Todo

func (repo *DB) Create(t *Todo) *gorm.DB {
	return repo.Create(t)
}

func (repo *DB) Search(q Query, todos *[]Todo) *gorm.DB {
	var r *gorm.DB

	if q.Title != "" {
		r = repo.Where("title = ?", q.Title)
	}

	if q.UserId != "" {
		r = repo.Where("user_id = ?", q.UserId)
	}

	return r.Find(todos)
}
