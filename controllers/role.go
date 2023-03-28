package controllers

import (
	"fmt"

	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/models"
	"github.com/gofiber/fiber/v2"
)

func GetRole(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var role models.Role
		if err := configs.Store.Preload("RoleDetail.Role").First(&role, &models.Role{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Message = fmt.Sprintf("Show role %s", c.Query("id"))
		r.Data = &role
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var role []models.Role
	if err := configs.Store.Preload("RoleDetail.Role").Preload("RoleDetail.Permission").Find(&role).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = "Show All role"
	r.Data = &role
	return c.Status(fiber.StatusOK).JSON(&r)
}

func PostRole(c *fiber.Ctx) error {
	var r models.Response
	var frm models.FrmRole
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var role models.Role
	role.Title = frm.Title
	role.Description = frm.Description
	role.IsActive = frm.IsActive
	if err := configs.Store.Create(&role).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	for _, v := range frm.FrmRoleDetail {
		var RoleDetail models.RoleDetail
		RoleDetail.RoleID = role.ID
		RoleDetail.PermissionID = v.ID
		switch v.Type {
		case "create":
			RoleDetail.Create = v.Status
		case "read":
			RoleDetail.Write = v.Status
		case "write":
			RoleDetail.Read = v.Status
		}

		if err := configs.Store.FirstOrCreate(&RoleDetail, &models.RoleDetail{RoleID: role.ID, PermissionID: v.ID}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		// After Create Save Data
		if err := configs.Store.Save(&RoleDetail).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
	}

	r.Message = "Role created successfully"
	r.Data = &frm
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func PutRole(c *fiber.Ctx) error {
	var r models.Response
	var frm models.FrmRole
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var role models.Role
	if err := configs.Store.First(&role, &models.Role{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	role.Title = frm.Title
	role.Description = frm.Description
	role.IsActive = frm.IsActive

	if err := configs.Store.Save(&role).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	for _, v := range frm.FrmRoleDetail {
		var RoleDetail models.RoleDetail
		RoleDetail.RoleID = role.ID
		RoleDetail.PermissionID = v.ID
		switch v.Type {
		case "create":
			RoleDetail.Create = v.Status
		case "read":
			RoleDetail.Write = v.Status
		case "write":
			RoleDetail.Read = v.Status
		}

		if err := configs.Store.FirstOrCreate(&RoleDetail, &models.RoleDetail{RoleID: role.ID, PermissionID: v.ID}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		// After Create Save Data
		if err := configs.Store.Save(&RoleDetail).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
	}

	r.Message = "Update successfull"
	r.Data = &role
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteRole(c *fiber.Ctx) error {
	var r models.Response
	var role models.Role
	if err := configs.Store.First(&role, &models.Role{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := configs.Store.Delete(&models.RoleDetail{}, &models.RoleDetail{RoleID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if err := configs.Store.Delete(&role).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Delete role successfull"
	return c.Status(fiber.StatusOK).JSON(&r)
}
