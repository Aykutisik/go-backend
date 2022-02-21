package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoElements struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	Status string `json:"status"`
}

type Handler interface {
	SaveTodo() interface{}
	GetTodoElements() interface{}
}

type DatabaseHandler struct {
	mongoClient            *mongo.Client
	TodoDatabase           *mongo.Database
	TodoElementsCollection *mongo.Collection
	basketCollection       *mongo.Collection
}

func NewDatabaseHandler(mc *mongo.Client, td *mongo.Database, tec *mongo.Collection, bc *mongo.Collection) DatabaseHandler {
	return DatabaseHandler{mongoClient: mc, TodoDatabase: td, TodoElementsCollection: tec, basketCollection: bc}
}

func (d DatabaseHandler) SaveTodo(c *fiber.Ctx) interface{} {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	todoElement := todoElements{}
	if err := c.BodyParser(&todoElement); err != nil {
		return err
	}

	_, err := d.TodoElementsCollection.InsertOne(ctx, todoElement)
	if err != nil {
		return err
	}

	return nil

}

func (d DatabaseHandler) GetTodoElements() interface{} {

	return nil
}
