package routes

import (
	"github.com/Pavel-Hrdina/todo/backend/services"
	"github.com/gofiber/fiber/v2"
)

// Contains all routes related to TODO
// Takes in a router object and return nothing
func TodoRoutes(app fiber.Router) {
	r := app.Group("/todo")

	r.Post("/", services.CreateTodo)
	r.Get("/", services.GetTodos)
	r.Get("/:todoID", services.GetTodo)
	r.Patch("/:todoID", services.UpdateTodoTitle)
	r.Patch("/:todoID/check", services.CheckTodo)
	r.Delete("/:todoID", services.DeleteTodo)
}
