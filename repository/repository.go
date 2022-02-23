package repository

import (
	"Desktop/todo-backend/go-backend/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateTodo(todo model.TodoElements) error
	GetTodoElements() (todos []model.TodoElements, err error)
}

type repository struct {
	db                     *mongo.Database
	mongoClient            *mongo.Client
	TodoElementsCollection *mongo.Collection
}

var _ Repository = repository{}

func NewRepository(db *mongo.Database, mongoClient *mongo.Client, TodoElementsCollection *mongo.Collection) Repository {
	return repository{db: db, mongoClient: mongoClient, TodoElementsCollection: TodoElementsCollection}
}

func (r repository) GetTodoElements() (todos []model.TodoElements, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := r.TodoElementsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var elements []bson.M
	if err = cursor.All(ctx, &todos); err != nil {
		log.Fatal(err)
	}

	bsonBytes, _ := bson.Marshal(elements)
	bson.Unmarshal(bsonBytes, &todos)

	return todos, err
}

func (r repository) CreateTodo(todo model.TodoElements) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	r.TodoElementsCollection.InsertOne(ctx, todo)

	return nil
}
