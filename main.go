package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/memstats", memstats)
	app.Get("/cpustats", cpustats)

	log.Fatal(app.Listen(":3000"))
}
