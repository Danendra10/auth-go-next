package main

import (
	"github.com/danendra10/auth-go-next/database"
	"github.com/gofiber/fiber"
	"github.ocm/danendra10/auth-go-next/routes"
)

func main() {

	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(8000)
}
