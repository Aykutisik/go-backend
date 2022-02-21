package service

import "Desktop/todo-backend/go-backend/handler"

type Service interface {
	SaveTodo() interface{}
	GetTodoElements() interface{}
}

type ActionService struct {
	handler handler.Handler
}

func newActionService(handler handler.Handler) Service {
	return ActionService{handler}
}

func (a ActionService) SaveTodo() interface{} {
	return a.handler.SaveTodo()
}

func (a ActionService) GetTodoElements() interface{} {
	return a.handler.GetTodoElements()
}
