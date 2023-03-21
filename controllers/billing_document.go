package controllers

import (
	"fmt"

	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/models"
	"github.com/gofiber/fiber/v2"
)

func GetBillingDocument(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var billDocument models.BillingDocument
		if err := configs.Store.
			Preload("Billing").
			Preload("DocumentList").
			First(&billDocument, &models.BillingDocument{}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Message = fmt.Sprintf("Show By %s", c.Query("id"))
		r.Data = &billDocument
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var billDocument []models.BillingDocument
	if err := configs.Store.
		Preload("Billing").
		Preload("DocumentList").
		Find(&billDocument).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Message = "Show All"
	r.Data = &billDocument
	return c.Status(fiber.StatusOK).JSON(&r)
}

func PostBillingDocument(c *fiber.Ctx) error {
	var r models.Response
	var frm models.BillingDocument
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	// Billing
	var billing models.Billing
	if err := configs.Store.First(&billing, &models.Billing{BillingNo: *frm.BillingID}).Error; err != nil {
		r.Message = fmt.Sprintf("Billing %s", *frm.BillingID)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// DocumentList
	var docList models.DocumentList
	if err := configs.Store.First(&docList, &models.DocumentList{Title: *frm.DocumentID}).Error; err != nil {
		r.Message = fmt.Sprintf("Document %s", *frm.DocumentID)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var billDocument models.BillingDocument
	billDocument.FileName = frm.FileName
	billDocument.FileSize = frm.FileSize
	billDocument.FileType = frm.FileType
	billDocument.FilePath = frm.FilePath
	billDocument.BillingID = &billing.ID
	billDocument.DocumentID = &docList.ID
	billDocument.IsActive = frm.IsActive
	if err := configs.Store.Create(&billDocument).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	billDocument.Billing = &billing
	billDocument.DocumentList = &docList
	r.Message = "Created Billing Document"
	r.Data = &billDocument
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func PutBillingDocument(c *fiber.Ctx) error {
	var r models.Response
	var frm models.BillingDocument
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var billDocument models.BillingDocument
	if err := configs.Store.First(&billDocument, &models.BillingDocument{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	// Billing
	var billing models.Billing
	if err := configs.Store.First(&billing, &models.Billing{BillingNo: *frm.BillingID}).Error; err != nil {
		r.Message = fmt.Sprintf("Billing %s", *frm.BillingID)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// DocumentList
	var docList models.DocumentList
	if err := configs.Store.First(&docList, &models.DocumentList{Title: *frm.DocumentID}).Error; err != nil {
		r.Message = fmt.Sprintf("Document %s", *frm.DocumentID)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	billDocument.FileName = frm.FileName
	billDocument.FileSize = frm.FileSize
	billDocument.FileType = frm.FileType
	billDocument.FilePath = frm.FilePath
	billDocument.BillingID = &billing.ID
	billDocument.DocumentID = &docList.ID
	billDocument.IsActive = frm.IsActive
	if err := configs.Store.Save(&billDocument).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	billDocument.Billing = &billing
	billDocument.DocumentList = &docList
	r.Message = "Updated Billing Document"
	r.Data = &billDocument
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func DeleteBillingDocument(c *fiber.Ctx) error {
	var r models.Response
	var billDocument models.BillingDocument
	if err := configs.Store.First(&billDocument, &models.BillingDocument{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := configs.Store.Delete(&billDocument).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Message = fmt.Sprintf("Deleted %s", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(&r)
}
