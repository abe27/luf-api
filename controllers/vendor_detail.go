package controllers

import (
	"fmt"

	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/models"
	"github.com/gofiber/fiber/v2"
)

func GetVendorDetail(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var obj models.Vendor
		if err := configs.Store.
			Preload("VendorGroup").
			Preload("DocumentList").
			Preload("Role").
			First(&obj, &models.Vendor{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Message = fmt.Sprintf("Show by ID %s", c.Query("id"))
		r.Data = &obj
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var obj []models.Vendor
	if err := configs.Store.
		Preload("VendorGroup").
		Preload("DocumentList").
		Preload("Role").
		Find(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = "Show All"
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func PostVendorDetail(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Vendor
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var vGroup models.VendorGroup
	if err := configs.Store.First(&vGroup, &models.VendorGroup{Title: *frm.VendorGroupID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var document models.DocumentList
	if err := configs.Store.First(&document, &models.DocumentList{Title: *frm.DocumentID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var role models.Role
	if err := configs.Store.First(&role, &models.Role{Title: *frm.RoleID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var obj models.Vendor
	obj.VendorGroupID = &vGroup.ID
	obj.DocumentID = &document.ID
	obj.RoleID = &role.ID
	obj.IsActive = frm.IsActive
	if err := configs.Store.Create(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Data created successfully"
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func PutVendorDetail(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Vendor
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var obj models.Vendor
	if err := configs.Store.First(&obj, &models.Vendor{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var vGroup models.VendorGroup
	if err := configs.Store.First(&vGroup, &models.VendorGroup{Title: *frm.VendorGroupID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var document models.DocumentList
	if err := configs.Store.First(&document, &models.DocumentList{Title: *frm.DocumentID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var role models.Role
	if err := configs.Store.First(&role, &models.Role{Title: *frm.RoleID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	obj.VendorGroupID = &vGroup.ID
	obj.DocumentID = &document.ID
	obj.RoleID = &role.ID
	obj.IsActive = frm.IsActive
	if err := configs.Store.Save(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Update successfull"
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteVendorDetail(c *fiber.Ctx) error {
	var r models.Response
	var obj models.Vendor
	if err := configs.Store.First(&obj, &models.Vendor{ID: c.Params("id")}).Error; err != nil {
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
