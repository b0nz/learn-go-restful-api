package test

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"testing"
	"net/http/httptest"
	"strings"
	"github.com/stretchr/testify/assert"

	"learn-go-restful-api/app"
	"learn-go-restful-api/repository"
	"learn-go-restful-api/controller"
	"learn-go-restful-api/service"
	"learn-go-restful-api/middleware"
	"learn-go-restful-api/helper"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/personal")
	helper.PanicIfError(err)
	
	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func setupRouter() http.Handler {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategoryFailed(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"name": "Handphone"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3001/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
}
// func TestCreateCategorySuccess(t *testing.T) {}

// func TestUpdateCategoryFailed(t *testing.T) {}
// func TestUpdateCategorySuccess(t *testing.T) {}

// func TestGetCategoryFailure(t *testing.T) {}
// func TestGetCategorySuccess(t *testing.T) {}

// func TestDeleteCategoryFailure(t *testing.T) {}
// func TestDeleteCategorySuccess(t *testing.T) {}

// func TestListCategoriesSuccess(t *testing.T) {}

// func TestUnauthorized(t *testing.T) {}