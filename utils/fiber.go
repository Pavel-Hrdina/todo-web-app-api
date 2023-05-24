package utils

import "github.com/gofiber/fiber/v2"

// helper function for parsing the body
// if error occurs, panic
// its purpose is to avoid writing if else condition again and again
func ParseBody(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

// helper function that gets the todo id
func GetTodoId(c *fiber.Ctx) *uint {
	id, _ := c.Locals("id").(uint)

	return &id
}
