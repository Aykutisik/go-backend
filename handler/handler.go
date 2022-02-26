package handler

import (
	"Desktop/todo-backend/go-backend/model"
	"Desktop/todo-backend/go-backend/service"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CreateTodo(ctx *fiber.Ctx) error
	GetTodoElements(ctx *fiber.Ctx) error
	DeleteTodo(ctx *fiber.Ctx) error
	UpdateTodo(ctx *fiber.Ctx) error
}

type handler struct {
	service service.Service
}

var _ Handler = handler{}

func NewHandler(service service.Service) Handler {
	return handler{service: service}
}

type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func (h handler) GetTodoElements(c *fiber.Ctx) error {
	model, err := h.service.GetTodoElements()
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}

	return c.Status(200).JSON(model)
}

func (h handler) CreateTodo(c *fiber.Ctx) error {
	todo := model.TodoElements{}

	err := c.BodyParser(&todo)

	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}

	err = h.service.CreateTodo(todo)

	return c.SendStatus(201)
}

func (h handler) DeleteTodo(c *fiber.Ctx) error {

	id := c.Params("id")

	err := h.service.DeleteTodo(id)

	return err
}

func (h handler) UpdateTodo(c *fiber.Ctx) error {

	todo := model.TodoElements{}

	err := c.BodyParser(&todo)

	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}

	err = h.service.UpdateTodo(todo)

	return err
}
