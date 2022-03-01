package handler

import (
	"Desktop/todo-backend/go-backend/model"
	"Desktop/todo-backend/go-backend/repository"
	"Desktop/todo-backend/go-backend/service"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetTodoElements(t *testing.T) {

	t.Run("Database connection should be established.",
		func(t *testing.T) {
			mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017/?authSource=admin"))
			assert.Nil(t, err)

			err = mongoClient.Connect(context.Background())
			assert.Nil(t, err)
		})

	t.Run("GetTodoElements should response with status code 200",
		func(t *testing.T) {
			mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017/?authSource=admin"))
			assert.Nil(t, err)

			err = mongoClient.Connect(context.Background())
			assert.Nil(t, err)

			database := mongoClient.Database("todo_database")
			collection := database.Collection("todo_list_elements")

			repo := repository.NewRepository(database, mongoClient, collection)
			service := service.NewService(repo)
			handler := NewHandler(service)

			app := fiber.New()

			app.Get("/GetTodoElements", handler.GetTodoElements)

			req := httptest.NewRequest("GET", fmt.Sprintf("/GetTodoElements"), nil)

			res, err := app.Test(req)
			assert.Nil(t, err)
			assert.Equal(t, 200, res.StatusCode)

		})

}

func TestCreateTodo(t *testing.T) {

	t.Run("CreateTodo should response with status code 200",
		func(t *testing.T) {
			mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017/?authSource=admin"))
			assert.Nil(t, err)

			err = mongoClient.Connect(context.Background())
			assert.Nil(t, err)

			database := mongoClient.Database("todo_database")
			collection := database.Collection("todo_list_elements")

			repo := repository.NewRepository(database, mongoClient, collection)
			service := service.NewService(repo)
			handler := NewHandler(service)

			app := fiber.New()

			app.Post("/CreateTodo", handler.CreateTodo)
			testBody := model.TodoElements{Text: "testText", Status: 0}

			requestByte, _ := json.Marshal(testBody)
			requestReader := bytes.NewReader(requestByte)

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/CreateTodo"), requestReader)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			res, err := app.Test(req)
			assert.Nil(t, err)
			fmt.Println(err)
			assert.Equal(t, 201, res.StatusCode)

		})

}

func TestDeleteTodo(t *testing.T) {

	t.Run("DeleteTodo should response with status code 200",
		func(t *testing.T) {
			mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017/?authSource=admin"))
			assert.Nil(t, err)

			err = mongoClient.Connect(context.Background())
			assert.Nil(t, err)

			database := mongoClient.Database("todo_database")
			collection := database.Collection("todo_list_elements")

			repo := repository.NewRepository(database, mongoClient, collection)
			service := service.NewService(repo)
			handler := NewHandler(service)

			app := fiber.New()

			app.Post("/DeleteTodo", handler.DeleteTodo)
			// testBody := model.TodoElements{Text: "testText", Status: 0}

			// requestByte, _ := json.Marshal(testBody)
			// requestReader := bytes.NewReader(requestByte)
			id := "621a1b360e2eeedb9c57f99b"
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/DeleteTodo/%s", id), nil)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			res, err := app.Test(req)
			assert.Nil(t, err)
			fmt.Println(err)
			assert.Equal(t, 200, res.StatusCode)

		})

}
