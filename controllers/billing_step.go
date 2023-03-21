package controllers

import (
	"fmt"

	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/models"
	"github.com/gofiber/fiber/v2"
)

func GetBillingStep(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var billStatus models.BillingStatus
		if err := configs.Store.
			Preload("Billing").
			Preload("StepTitle").
			First(&billStatus, &models.BillingStatus{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Message = fmt.Sprintf("Show By ID: %v", c.Query("id"))
		r.Data = &billStatus
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var billStatus []models.BillingStatus
	if err := configs.Store.
		Preload("Billing").
		Preload("StepTitle").
		Find(&billStatus).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Message = "Show all"
	r.Data = &billStatus
	return c.Status(fiber.StatusOK).JSON(&r)
}

func PostBillingStep(c *fiber.Ctx) error {
	var r models.Response
	var frm models.BillingStatus
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

	// StepTitle
	var step models.StepTitle
	if err := configs.Store.First(&step, &models.StepTitle{Title: *frm.StepTitleID}).Error; err != nil {
		r.Message = fmt.Sprintf("Step Title %s", *frm.BillingID)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var billingStatus models.BillingStatus
	billingStatus.StepTitleID = &step.ID
	billingStatus.BillingID = &billing.ID
	billingStatus.IsComplete = frm.IsComplete
	if err := configs.Store.Create(&billingStatus).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Created Billing Status"
	r.Data = &billingStatus
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func PutBillingStep(c *fiber.Ctx) error {
	var r models.Response
	var frm models.BillingStatus
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var billingStatus models.BillingStatus
	if err := configs.Store.First(&billingStatus, &models.BillingStatus{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// Billing
	var billing models.Billing
	if err := configs.Store.First(&billing, &models.Billing{BillingNo: *frm.BillingID}).Error; err != nil {
		r.Message = fmt.Sprintf("Billing %s", *frm.BillingID)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// StepTitle
	var step models.StepTitle
	if err := configs.Store.First(&step, &models.StepTitle{Title: *frm.StepTitleID}).Error; err != nil {
		r.Message = fmt.Sprintf("Step Title %s", *frm.BillingID)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	billingStatus.StepTitleID = &step.ID
	billingStatus.BillingID = &billing.ID
	billingStatus.IsComplete = frm.IsComplete
	if err := configs.Store.Create(&billingStatus).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = fmt.Sprintf("Updated %s successfully", c.Params("id"))
	billingStatus.StepTitle = &step
	billingStatus.Billing = &billing
	r.Data = &billingStatus
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteBillingStep(c *fiber.Ctx) error {
	var r models.Response
	var billingStatus models.BillingStatus
	if err := configs.Store.First(&billingStatus, &models.BillingStatus{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	if err := configs.Store.Delete(&billingStatus).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = fmt.Sprintf("Delete %s successfully", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(&r)
}
