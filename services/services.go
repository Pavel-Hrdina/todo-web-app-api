package services

import (
	"errors"

	"github.com/Pavel-Hrdina/todo/backend/api"
	"github.com/Pavel-Hrdina/todo/backend/api/types"
	"github.com/Pavel-Hrdina/todo/backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// creates a new todo
func CreateTodo(c *fiber.Ctx) error {
	b := new(types.CreateDTO)

	if err := utils.ParseBody(c, b); err != nil {
		return err
	}

	a := &api.Todo{
		Task: b.Task,
	}

	if err := api.CreateTodo(a).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.TodoCreateRespone{
		Todo: &types.TodoResponse{
			ID:        a.ID,
			Task:      a.Task,
			Completed: a.Completed,
		},
	})
}

// returns a list of todos
// TODO: find why this method works, and why
// GetTodo() doesn't
func GetTodos(c *fiber.Ctx) error {
	a := &[]types.TodoResponse{}

	// This is the breaking point
	err := api.FindTodos(a).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.TodosResponse{
		Todos: a,
	})
}

// returns a single todo
// TODO: fix, this method, no idea what is wrong, help me
func GetTodo(c *fiber.Ctx) error {
	// Pass context as parameter
	todoID := c.Params("todoID")

	if todoID != "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	// save the answer in a variable so it can
	// be passed as a destination.
	a := &types.TodoResponse{}

	// TODO: fix wrong id parsing, probably something wrong
	// in GetTodoId().
	// NOTE: This is the breaking point
	err := api.FindTodoByID(a, todoID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(&types.TodoCreateRespone{})
	}

	// returns todo if no problems occur
	return c.JSON(&types.TodoCreateRespone{
		Todo: a,
	})
}

// deletes a single todo
// TODO: fix, todo id not recognized
func DeleteTodo(c *fiber.Ctx) error {
	todoID := c.Params("todoID")

	if todoID != "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todo id")
	}

	// NOTE: possible breaking point
	res := api.DeleteTodo(todoID)
	if res.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusConflict, "Unable to delete todo")
	}

	err := res.Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Message: "Todo deleted",
	})
}

// check todo
// TODO: fix this method, return a bad request
func CheckTodo(c *fiber.Ctx) error {
	b := new(types.CheckTodoDTO)
	todoID := c.Params("todoID")
	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoid")
	}

	if err := utils.ParseBody(c, b); err != nil {
		return err
	}

	// NOTE: possible breaking point
	err := api.UpdateTodo(todoID, map[string]interface{}{"completed": b.Completed}).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Message: "Todo updated",
	})
}

// update todo tile
// Does not work, the todo title will not update
// Possible couse is a bad request to the api
func UpdateTodoTitle(c *fiber.Ctx) error {
	b := new(types.CreateDTO)
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	err := api.UpdateTodo(todoID, &api.Todo{Task: b.Task}).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Message: "Todo title updated",
	})
}
