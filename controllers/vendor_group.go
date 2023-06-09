package controllers

import (
	"fmt"

	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/models"
	"github.com/gofiber/fiber/v2"
)

func GetVendorGroup(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var obj models.VendorGroup
		if err := configs.Store.First(&obj, &models.VendorGroup{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Message = fmt.Sprintf("Show by ID %s", c.Query("id"))
		r.Data = &obj
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var obj []models.VendorGroup
	if err := configs.Store.
		Preload("Documents.VendorGroup").
		Preload("Documents.DocumentList").
		Preload("Documents.Role").
		Where("title <> ?", "-").
		Find(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = "Show All"
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func PostVendorGroup(c *fiber.Ctx) error {
	var r models.Response
	var frm models.FrmVendorGroup
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var obj models.VendorGroup
	obj.Title = frm.Title
	obj.Description = frm.Description
	obj.IsActive = frm.IsActive
	if err := configs.Store.Create(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	for _, v := range frm.Documents {
		var objVendor models.Vendor
		objVendor.VendorGroupID = &obj.ID
		objVendor.DocumentID = &v
		objVendor.IsActive = frm.IsActive
		if err := configs.Store.Create(&objVendor).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
	}

	r.Message = "Data created successfully"
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func PutVendorGroup(c *fiber.Ctx) error {
	var r models.Response
	var frm models.FrmVendorGroup
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var obj models.VendorGroup
	if err := configs.Store.First(&obj, &models.VendorGroup{ID: c.Params("id")}).Error; err != nil {
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

	for _, v := range frm.Documents {
		var objVendor models.Vendor
		objVendor.VendorGroupID = &obj.ID
		objVendor.DocumentID = &v
		objVendor.IsActive = frm.IsActive
		if err := configs.Store.First(&objVendor, &models.Vendor{VendorGroupID: &obj.ID, DocumentID: &v}).Error; err != nil {
			if err := configs.Store.Create(&objVendor).Error; err != nil {
				r.Message = err.Error()
				return c.Status(fiber.StatusInternalServerError).JSON(&r)
			}
		} else {
			if err := configs.Store.Save(&objVendor).Error; err != nil {
				r.Message = err.Error()
				return c.Status(fiber.StatusInternalServerError).JSON(&r)
			}
		}
	}

	r.Message = "Update successfull"
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteVendorGroup(c *fiber.Ctx) error {
	var r models.Response
	var obj models.VendorGroup
	if err := configs.Store.First(&obj, &models.VendorGroup{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	configs.Store.Delete(&models.Vendor{}, &models.Vendor{VendorGroupID: &obj.ID})
	if err := configs.Store.Delete(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Delete data successfull"
	return c.Status(fiber.StatusOK).JSON(&r)
}
