package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/models"
	"github.com/gofiber/fiber/v2"
	"github.com/shakinm/xlsReader/xls"
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

	if c.Query("billing_no") != "" && c.Query("billing_no") != "null" {
		if err := configs.Store.
			Preload("Status").
			Preload("VendorGroup").
			Find(&billing, &models.Billing{BillingNo: c.Query("billing_no")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = "Show All"
		r.Data = &billing
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("billing_date") != "" && c.Query("billing_date") != "null" {
		ftime, _ := time.Parse("2006-01-02", c.Query("billing_date"))
		if err := configs.Store.
			Preload("Status").
			Preload("VendorGroup").
			Find(&billing, &models.Billing{BillingDate: ftime}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = "Show All"
		r.Data = &billing
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("vendor_group") != "" && c.Query("vendor_group") != "null" {
		if err := configs.Store.
			Preload("Status").
			Preload("VendorGroup").
			Find(&billing, &models.Billing{VendorGroupID: c.Query("vendor_group")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = "Show All"
		r.Data = &billing
		return c.Status(fiber.StatusOK).JSON(&r)
	}

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
	if err := configs.Store.First(&vendorGroup, &models.VendorGroup{ID: frm.VendorGroupID}).Error; err != nil {
		r.Message = fmt.Sprintf("Vendor Group: %v", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var billing models.Billing
	billing.BillingNo = strings.ToUpper(frm.BillingNo)
	billing.BillingDate = frm.BillingDate
	billing.DueDate = frm.DueDate
	billing.Amount = frm.Amount
	billing.VendorCode = strings.ToUpper(frm.VendorCode)
	billing.VendorName = strings.ToUpper(frm.VendorName)
	// if frm.PaymentDate != "" {
	// 	billing.PaymentDate = frm.PaymentDate
	// }
	billing.Detail = frm.Detail
	billing.StatusID = &status.ID
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
	if err := configs.Store.First(&vendorGroup, &models.VendorGroup{ID: frm.VendorGroupID}).Error; err != nil {
		r.Message = fmt.Sprintf("Vendor Group: %v", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var billing models.Billing
	if err := configs.Store.First(&billing, &models.Billing{ID: c.Params("id")}).Error; err != nil {
		r.Message = fmt.Sprintf("Notfound ID: %s", c.Params("id"))
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	billing.BillingNo = strings.ToUpper(frm.BillingNo)
	billing.BillingDate = frm.BillingDate
	billing.DueDate = frm.DueDate
	billing.Amount = frm.Amount
	billing.VendorCode = strings.ToUpper(frm.VendorCode)
	billing.VendorName = strings.ToUpper(frm.VendorName)
	// if frm.PaymentDate != "" {
	// 	billing.PaymentDate = frm.PaymentDate
	// }
	billing.Detail = frm.Detail
	billing.StatusID = &status.ID
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

func ImportBilling(c *fiber.Ctx) error {
	var res models.Response
	file, err := c.FormFile("filename")
	if err != nil {
		res.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}

	fName := fmt.Sprintf("./public/import/%s", file.Filename)
	if err := c.SaveFile(file, fName); err != nil {
		res.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}

	workbook, err := xls.OpenFile(fName)
	if err != nil {
		res.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&res)
	}

	for _, sheet := range workbook.GetSheets() {
		i := 0
		for _, r := range sheet.GetRows() {
			if i > 0 {
				if billing_no, err := r.GetCol(0); err == nil {
					billing_date, _ := r.GetCol(1)
					due_date, _ := r.GetCol(2)
					amount, _ := r.GetCol(3)
					vendor_code, _ := r.GetCol(4)
					vendor_name, _ := r.GetCol(5)
					vendor_group_id, _ := r.GetCol(6)

					billing_no.GetString()
					billing_date.GetString()
					due_date.GetString()
					amount.GetFloat64()
					vendor_code.GetString()
					vendor_name.GetString()
					vendor_group_id.GetString()

					var status models.Status
					if err := configs.Store.First(&status, &models.Status{Title: "Open"}).Error; err != nil {
						res.Message = fmt.Sprintf("Status: %v", err.Error())
						return c.Status(fiber.StatusNotFound).JSON(&res)
					}

					var vendorGroup models.VendorGroup
					vendorGroup.Title = strings.ToUpper(vendor_group_id.GetString())
					if err := configs.Store.FirstOrCreate(&vendorGroup, &models.VendorGroup{Title: vendor_group_id.GetString()}).Error; err != nil {
						res.Message = fmt.Sprintf("Vendor Group: %v", err.Error())
						return c.Status(fiber.StatusNotFound).JSON(&res)
					}

					bDate, _ := time.Parse("02/01/2006", billing_date.GetString())
					dDate, _ := time.Parse("02/01/2006", due_date.GetString())

					var billing models.Billing
					billing.BillingNo = strings.ToUpper(billing_no.GetString())
					billing.BillingDate = bDate
					billing.DueDate = dDate
					billing.Amount = amount.GetFloat64()
					billing.VendorCode = strings.ToUpper(vendor_code.GetString())
					billing.VendorName = strings.ToUpper(vendor_name.GetString())
					billing.StatusID = &status.ID
					billing.VendorGroupID = vendorGroup.ID
					billing.IsActive = true
					if errDB := configs.Store.FirstOrCreate(&billing, &models.Billing{BillingNo: strings.ToUpper(billing_no.GetString())}).Error; errDB != nil {
						res.Message = errDB.Error()
						return c.Status(fiber.StatusInternalServerError).JSON(&res)
					}
				}
			}
			i++
		}
	}

	res.Message = "Upload Excel Completed"
	return c.Status(fiber.StatusOK).JSON(&res)
}
