package controllers

import (
	"github.com/abe27/luckyapp/models"
	"github.com/gofiber/fiber/v2"
)

func GetBillingRequireDocument(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func PostBillingRequireDocument(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func PutBillingRequireDocument(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteBillingRequireDocument(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
