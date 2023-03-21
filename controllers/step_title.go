package controllers

import (
	"fmt"

	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/models"
	"github.com/gofiber/fiber/v2"
)

func GetStep(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var obj models.StepTitle
		if err := configs.Store.First(&obj, &models.StepTitle{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Message = fmt.Sprintf("Show by ID %s", c.Query("id"))
		r.Data = &obj
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var obj []models.StepTitle
	if err := configs.Store.Find(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = "Show All"
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func PostStep(c *fiber.Ctx) error {
	var r models.Response
	var frm models.StepTitle
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var obj models.StepTitle
	obj.Title = frm.Title
	obj.Description = frm.Description
	obj.IsActive = frm.IsActive
	if err := configs.Store.Create(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Data created successfully"
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func PutStep(c *fiber.Ctx) error {
	var r models.Response
	var frm models.StepTitle
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var obj models.StepTitle
	if err := configs.Store.First(&obj, &models.StepTitle{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	obj.Title = frm.Title
	obj.Description = frm.Description
	obj.IsActive = frm.IsActive

	if err := configs.Store.Save(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Update successfull"
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteStep(c *fiber.Ctx) error {
	var r models.Response
	var obj models.StepTitle
	if err := configs.Store.First(&obj, &models.StepTitle{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := configs.Store.Delete(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Delete data successfull"
	return c.Status(fiber.StatusOK).JSON(&r)
}
