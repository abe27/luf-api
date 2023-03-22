package controllers

import (
	"github.com/abe27/luckyapp/models"
	"github.com/abe27/luckyapp/services"
	"github.com/gofiber/fiber/v2"
)

func TestSendMail(c *fiber.Ctx) error {
	var r models.Response
	services.SendMail("krumii.it@gmail.com", "test")
	return c.Status(fiber.StatusOK).JSON(&r)
}
