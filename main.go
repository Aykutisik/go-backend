package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func main() {
	app := NewApplication()

	app.Listen(":8086")
}

type todoElements struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	Status string `json:"status"`
}

type Handler struct {
	mongoClient            *mongo.Client
	TodoDatabase           *mongo.Database
	TodoElementsCollection *mongo.Collection
	basketCollection       *mongo.Collection
}

func NewHandler(mongoClient *mongo.Client) Handler {
	todoDatabase := mongoClient.Database("todo_database")
	return Handler{
		mongoClient:            mongoClient,
		TodoDatabase:           todoDatabase,
		TodoElementsCollection: todoDatabase.Collection("todo_list_elements"),
	}
}

func (h Handler) GetTodoList(c *fiber.Ctx) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := h.TodoElementsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var products []bson.M
	if err = cursor.All(ctx, &products); err != nil {
		log.Fatal(err)
	}

	return c.JSON(products)
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

	handler := NewHandler(mongoClient)

	app.Get("/GetTodoList", handler.GetTodoList)

	return app
}
