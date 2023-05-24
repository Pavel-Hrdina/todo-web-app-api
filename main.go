package main

import (
	"fmt"

	"github.com/Pavel-Hrdina/todo/backend/api"
	"github.com/Pavel-Hrdina/todo/backend/config"
	"github.com/Pavel-Hrdina/todo/backend/config/database"
	"github.com/Pavel-Hrdina/todo/backend/routes"
	"github.com/Pavel-Hrdina/todo/backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()
	database.Migrate(&api.Todo{})

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())
	routes.TodoRoutes(app)

	app.Listen(fmt.Sprintf(":%v", config.PORT))
}
