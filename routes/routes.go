package routes

import (
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})
}
