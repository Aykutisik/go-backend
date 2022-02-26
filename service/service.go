package service

import (
	"Desktop/todo-backend/go-backend/model"
	"Desktop/todo-backend/go-backend/repository"
)

type Service interface {
	CreateTodo(todo model.TodoElements) error
	GetTodoElements() (todos []model.TodoElements, err error)
	DeleteTodo(id string) error
}

type service struct {
	repo repository.Repository
}

var _ Service = service{}

func NewService(repo repository.Repository) Service {
	return service{repo: repo}
}

func (s service) GetTodoElements() (todos []model.TodoElements, err error) {
	return s.repo.GetTodoElements()
}

func (s service) CreateTodo(todo model.TodoElements) error {
	return s.repo.CreateTodo(todo)
}

func (s service) DeleteTodo(id string) error {

	return s.repo.DeleteTodo(id)
}
