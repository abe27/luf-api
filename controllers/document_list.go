package controllers

import (
	"fmt"

	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/models"
	"github.com/gofiber/fiber/v2"
)

func GetDocumentList(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var obj models.DocumentList
		if err := configs.Store.First(&obj, &models.DocumentList{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Message = fmt.Sprintf("Show by ID %s", c.Query("id"))
		r.Data = &obj
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var obj []models.DocumentList
	if err := configs.Store.Find(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = "Show All"
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func PostDocumentList(c *fiber.Ctx) error {
	var r models.Response
	var frm models.DocumentList
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var obj models.DocumentList
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

func PutDocumentList(c *fiber.Ctx) error {
	var r models.Response
	var obj models.DocumentList
	if err := configs.Store.First(&obj, &models.DocumentList{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var frm models.DocumentList
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
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

func DeleteDocumentList(c *fiber.Ctx) error {
	var r models.Response
	var obj models.DocumentList
	if err := configs.Store.First(&obj, &models.DocumentList{ID: c.Params("id")}).Error; err != nil {
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
