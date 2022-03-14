package handler

import (
	"Desktop/todo-backend/go-backend/model"
	"Desktop/todo-backend/go-backend/repository"
	"Desktop/todo-backend/go-backend/service"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			serviceMockController := gomock.NewController(t)
			service := service.NewMockService(serviceMockController)
			handler := NewHandler(service)
			service.EXPECT().GetTodoElements().Return(getAllTodoElements(), nil)
			app := fiber.New()
			app.Get("/GetTodoElements", handler.GetTodoElements)
			req := httptest.NewRequest("GET", fmt.Sprintf("/GetTodoElements"), nil)
			res, err := app.Test(req)
			assert.Nil(t, err)
			assert.Equal(t, 200, res.StatusCode)

		})

	t.Run("When GetTodoElements failed",
		func(t *testing.T) {
			serviceMockController := gomock.NewController(t)
			service := service.NewMockService(serviceMockController)
			handler := NewHandler(service)
			service.EXPECT().GetTodoElements().Return(getAllTodoElements(), errors.New("An error occured"))
			app := fiber.New()
			app.Get("/GetTodoElements", handler.GetTodoElements)
			req := httptest.NewRequest("GET", fmt.Sprintf("/GetTodoElements"), nil)
			res, err := app.Test(req)
			assert.Nil(t, err)
			assert.Equal(t, 400, res.StatusCode)
		})
}

func TestCreateTodo(t *testing.T) {

	t.Run("Database connection should be established.",
		func(t *testing.T) {
			mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017/?authSource=admin"))
			assert.Nil(t, err)

			err = mongoClient.Connect(context.Background())
			assert.Nil(t, err)
		})

	t.Run("CreateTodo should response with status code 201",
		func(t *testing.T) {
			serviceMockController := gomock.NewController(t)
			service := service.NewMockService(serviceMockController)
			handler := NewHandler(service)
			testBody := model.SendTodoElements{Text: "asfwkfwkqfwqkfqw", Status: 0}

			requestByte, _ := json.Marshal(testBody)
			requestReader := bytes.NewReader(requestByte)
			service.EXPECT().CreateTodo(testBody).Return(nil)

			app := fiber.New()
			app.Post("/CreateTodo", handler.CreateTodo)

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/CreateTodo"), requestReader)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			res, err := app.Test(req)
			fmt.Println(err)
			assert.Equal(t, 201, res.StatusCode)
		})

	t.Run("Cannot request with empty text",
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
			testBody := model.SendTodoElements{Text: "", Status: 0}

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

			app.Post("/DeleteTodo/:id", handler.DeleteTodo)

			id := "622c785eed9caa5977fbad5c"
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/DeleteTodo/%s", id), nil)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			res, err := app.Test(req)
			assert.Nil(t, err)
			fmt.Println(err)
			assert.Equal(t, 200, res.StatusCode)
		})

}

func TestUpdateTodo(t *testing.T) {
	t.Run("UpdateTodo should response with status code 200 for status 0 to 1",
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

			app.Post("/UpdateTodo", handler.CreateTodo)
			testBody := model.TodoElements{Text: "testText", Status: 0}

			requestByte, _ := json.Marshal(testBody)
			requestReader := bytes.NewReader(requestByte)

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/UpdateTodo"), requestReader)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			res, err := app.Test(req)
			assert.Nil(t, err)
			fmt.Println(err)
			assert.Equal(t, 201, res.StatusCode)

		})

	t.Run("UpdateTodo should response with status code 200 for status 1 to 0",
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

			app.Post("/UpdateTodo", handler.CreateTodo)
			testBody := model.TodoElements{Text: "testText", Status: 1}

			requestByte, _ := json.Marshal(testBody)
			requestReader := bytes.NewReader(requestByte)

			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/UpdateTodo"), requestReader)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			res, err := app.Test(req)
			assert.Nil(t, err)
			fmt.Println(err)
			assert.Equal(t, 201, res.StatusCode)

		})

}

func getAllTodoElements() []model.TodoElements {
	var list []model.TodoElements
	//	[{ "_id": "62234346c2a65768f2c03ca5", "status": 0, "text": "drink water" }, { "_id": "6225bfaec2a65768f2c03ca6", "status": 0, "text": "bla bla" }]

	var item model.TodoElements

	item.Id, _ = primitive.ObjectIDFromHex("62234346c2a65768f2c03ca5")
	item.Text = "drink water"
	item.Status = 0

	list = append(list, item)

	item.Id, _ = primitive.ObjectIDFromHex("6225bfaec2a65768f2c03ca6")
	item.Text = "bla bla"
	item.Status = 0

	list = append(list, item)

	return list
}
