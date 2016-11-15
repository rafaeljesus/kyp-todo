package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/rafaeljesus/kyp-todo/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

var KYP_TODO_DB = os.Getenv("KYP_TODO_DB")
var env *Env

func TestMain(m *testing.M) {
	db, err := models.NewDB(KYP_TODO_DB)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Todo{})
	env = &Env{db}
	code := m.Run()
	os.Exit(code)
}

func TestTodosCreate(t *testing.T) {
	response := `{"title":"buy a milk", "user_id":1}`
	e := echo.New()
	req, _ := http.NewRequest(echo.POST, "/v1/todos", strings.NewReader(response))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	if assert.NoError(t, env.TodosCreate(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestTodosIndex(t *testing.T) {
	response := `{"title":"buy a milk", "user_id":1}`
	e := echo.New()
	q := make(url.Values)
	q.Set("title", "buy a milk")
	q.Set("user_id", "1")

	req, _ := http.NewRequest(echo.GET, "/v1/todos?"+q.Encode(), strings.NewReader(response))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	if assert.NoError(t, env.TodosIndex(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
