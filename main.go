package main

import "github.com/gofiber/fiber"

func setupRoutes() {
	app.Get()
	app.Get()
	app.Save()
	app.Post()
	app.Delete()
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}
