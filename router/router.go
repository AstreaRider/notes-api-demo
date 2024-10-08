package router

import (
	noteRoutes "github.com/AstreaRider/notes-api-demo/internal/routes/note"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api", logger.New())

    noteRoutes.SetupNoteRoutes(api)
}