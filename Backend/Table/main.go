package main

import (
	"log"
	database "service/table/Database"
	routes "service/table/Routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connection()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://gofiber.io",
	}))

	routes.Routing(app)

	err := app.Listen(":8005")
	if err != nil {
		log.Fatalf("Failde to listen: %v", err)
	}
}
