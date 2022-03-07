package main

import (
	"Desktop/todo-backend/go-backend/handler"
	"Desktop/todo-backend/go-backend/repository"
	"Desktop/todo-backend/go-backend/service"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	app := NewApplication()

	app.Listen(":8086")
}

func NewApplication() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017/?authSource=admin"))
	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("could not ping to mongo db service: %v\n", err)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	database := mongoClient.Database("todo_database")
	collection := database.Collection("todo_list_elements")
	repo := repository.NewRepository(database, mongoClient, collection)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	app.Get("/GetTodoElements", handler.GetTodoElements)
	app.Post("/CreateTodo", handler.CreateTodo)
	app.Put("/DeleteTodo/:id", handler.DeleteTodo)
	app.Put("/UpdateTodo", handler.UpdateTodo)

	return app
}
