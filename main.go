package main

import (
	"github.com/AstreaRider/notes-api-demo/database"
	"github.com/AstreaRider/notes-api-demo/router"
	"github.com/gofiber/fiber/v2"
)


func main()  {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)
	
	app.Get("/healthCheck", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})
	
	app.Listen(":3000")
}
