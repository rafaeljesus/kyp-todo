package handlers

import (
	"github.com/rafaeljesus/kyp-todo/config"
	"github.com/rafaeljesus/kyp-todo/models"
)

type Env struct {
	Repo     models.Repo
	EventBus *config.EventBus
}
