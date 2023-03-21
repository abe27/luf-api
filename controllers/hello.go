package controllers

import (
	"github.com/abe27/luckyapp/models"
	"github.com/gofiber/fiber/v2"
)

func HandlerHello(c *fiber.Ctx) error {
	var r models.Response

	return c.Status(fiber.StatusOK).JSON(&r)
}
