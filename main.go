package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	getCPUPercentage()

	app.Use(cors.New())
	app.Get("/memstats", memstats)
	app.Get("/cpustats", cpustats)

	log.Fatal(app.Listen(":3000"))
}
