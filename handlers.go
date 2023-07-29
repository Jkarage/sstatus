package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func memstats(c *fiber.Ctx) error {
	mem, err := getMemoryStatus()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON("Couldn't read memory details")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Total Memory": mem.Total,
		"Free Memory":  mem.Free,
		"Available":    mem.Available,
	})
}

func cpustats(c *fiber.Ctx) error {
	t, i, err := getCPUStats()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON("Couldn't read cpu details")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"CPU time": t,
		"CPU idle": i,
	})
}
