package controllers

import (
	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/models"
	"github.com/gofiber/fiber/v2"
)

func GetBillingRequireDocument(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func PostBillingRequireDocument(c *fiber.Ctx) error {
	var r models.Response
	var frm models.FrmReject
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var status models.Status
	if err := configs.Store.First(&status, &models.Status{Title: "Rejected"}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	for _, v := range frm.Reason {
		var bDoc models.BillingRequiredDocument
		bDoc.BillingID = c.Params("id")
		bDoc.DocumentID = v.ID
		bDoc.IsActive = v.Checked
		if err := configs.Store.Create(&bDoc).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
	}

	var billing models.Billing
	if err := configs.Store.First(&billing, &models.Billing{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	billing.StatusID = status.ID
	billing.Detail = frm.Remark
	if err := configs.Store.Save(&billing).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// loging
	CreateVendorLogger(&billing.ID, &status.ID, &billing.VendorGroupID, c)
	r.Message = "Update data successfully!"
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func PutBillingRequireDocument(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Status
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var status models.Status
	if err := configs.Store.First(&status, &models.Status{Title: frm.Title}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var billing models.Billing
	if err := configs.Store.First(&billing, &models.Billing{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	billing.StatusID = status.ID
	billing.Detail = ""

	if err := configs.Store.Save(&billing).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// loging
	CreateVendorLogger(&billing.ID, &status.ID, &billing.VendorGroupID, c)
	configs.Store.Delete(&models.BillingRequiredDocument{BillingID: c.Params("id")})
	r.Message = "Update data successfully!"
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteBillingRequireDocument(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
