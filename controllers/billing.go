package controllers

import (
	"fmt"

	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/models"
	"github.com/gofiber/fiber/v2"
)

func GetBilling(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var billing models.Billing
		if err := configs.Store.
			Preload("Status").
			Preload("VendorGroup").
			First(&billing, &models.Billing{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Message = fmt.Sprintf("Show %s", c.Query("id"))
		r.Data = &billing
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var billing []models.Billing
	if err := configs.Store.
		Preload("Status").
		Preload("VendorGroup").
		Find(&billing).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = "Show All"
	r.Data = &billing
	return c.Status(fiber.StatusOK).JSON(&r)
}

func PostBilling(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Billing
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var status models.Status
	if err := configs.Store.First(&status, &models.Status{Title: *frm.StatusID}).Error; err != nil {
		r.Message = fmt.Sprintf("Status: %v", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var vendorGroup models.VendorGroup
	if err := configs.Store.First(&vendorGroup, &models.VendorGroup{Title: *frm.VendorGroupID}).Error; err != nil {
		r.Message = fmt.Sprintf("Vendor Group: %v", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var billing models.Billing
	billing.BillingNo = frm.BillingNo
	billing.BillingDate = frm.BillingDate
	billing.DueDate = frm.DueDate
	billing.Amount = frm.Amount
	billing.VendorCode = frm.VendorCode
	billing.VendorName = frm.VendorName
	billing.PaymentDate = frm.PaymentDate
	billing.Detail = frm.Detail
	billing.StatusID = frm.StatusID
	billing.VendorGroupID = frm.VendorGroupID
	billing.IsActive = frm.IsActive

	if err := configs.Store.Create(&billing).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	billing.Status = &status
	billing.VendorGroup = &vendorGroup
	r.Message = "Created successfully"
	r.Data = &billing
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func PutBilling(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Billing
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var status models.Status
	if err := configs.Store.First(&status, &models.Status{Title: *frm.StatusID}).Error; err != nil {
		r.Message = fmt.Sprintf("Status: %v", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var vendorGroup models.VendorGroup
	if err := configs.Store.First(&vendorGroup, &models.VendorGroup{Title: *frm.VendorGroupID}).Error; err != nil {
		r.Message = fmt.Sprintf("Vendor Group: %v", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var billing models.Billing
	if err := configs.Store.First(&billing, &models.Billing{ID: c.Params("id")}).Error; err != nil {
		r.Message = fmt.Sprintf("Notfound ID: %s", c.Params("id"))
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	billing.BillingNo = frm.BillingNo
	billing.BillingDate = frm.BillingDate
	billing.DueDate = frm.DueDate
	billing.Amount = frm.Amount
	billing.VendorCode = frm.VendorCode
	billing.VendorName = frm.VendorName
	billing.PaymentDate = frm.PaymentDate
	billing.Detail = frm.Detail
	billing.StatusID = frm.StatusID
	billing.VendorGroupID = frm.VendorGroupID
	billing.IsActive = frm.IsActive

	if err := configs.Store.Save(&billing).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	billing.Status = &status
	billing.VendorGroup = &vendorGroup
	r.Message = "Updated successfully"
	r.Data = &billing
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteBilling(c *fiber.Ctx) error {
	var r models.Response
	var billing models.Billing
	if err := configs.Store.First(&billing, &models.Billing{ID: c.Params("id")}).Error; err != nil {
		r.Message = fmt.Sprintf("Notfound ID: %s", c.Params("id"))
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := configs.Store.Delete(&billing).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	return c.Status(fiber.StatusOK).JSON(&r)
}
