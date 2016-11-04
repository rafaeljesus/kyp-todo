package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rafaeljesus/kyp-todo/db"
	"time"
)

type Todo struct {
	Id        uint       `json:"id", sql:"primary_key"`
	Title     string     `json:"title", sql:"not null"`
	Done      bool       `json:"done"`
	UserId    uint       `json:"user_id, sql:"not null`
	CreatedAt time.Time  `json:"created_at", sql:"not null"`
	UpdatedAt time.Time  `json:"updated_at", sql:"not null`
	DeletedAt *time.Time `json:"-" "created_at"`
}

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
