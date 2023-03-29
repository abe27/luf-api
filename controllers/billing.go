package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/models"
	"github.com/gofiber/fiber/v2"
	g "github.com/matoous/go-nanoid/v2"
	"github.com/shakinm/xlsReader/xls"
)

func GetBilling(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var billing models.Billing
		if err := configs.Store.
			Preload("Status").
			Preload("VendorGroup").
			Preload("DocumentList.DocumentList").
			Preload("BillingStep").
			First(&billing, &models.Billing{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}

		r.Message = fmt.Sprintf("Show %s", c.Query("id"))
		r.Data = &billing
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var billing []models.Billing

	if c.Query("billing_no") != "" && c.Query("billing_date") != "" && c.Query("status_id") != "" {
		ftime, _ := time.Parse("2006-01-02", c.Query("billing_date"))
		if err := configs.Store.
			Preload("Status").
			Preload("VendorGroup").
			Preload("DocumentList.DocumentList").
			Preload("BillingStep").
			Where("billing_no like ?", "%"+c.Query("billing_no")+"%").
			Where("billing_date", ftime).
			Where("status_id", c.Query("status_id")).
			Find(&billing).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = "Show All"
		r.Data = &billing
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("billing_no") != "" && c.Query("status_id") != "" {
		if err := configs.Store.
			Preload("Status").
			Preload("VendorGroup").
			Preload("DocumentList.DocumentList").
			Preload("BillingStep").
			Where("billing_no like ?", "%"+c.Query("billing_no")+"%").
			Where("status_id", c.Query("status_id")).
			Find(&billing).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = "Show All"
		r.Data = &billing
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("billing_date") != "" && c.Query("status_id") != "" {
		ftime, _ := time.Parse("2006-01-02", c.Query("billing_date"))
		if err := configs.Store.
			Preload("Status").
			Preload("VendorGroup").
			Preload("DocumentList.DocumentList").
			Preload("BillingStep").
			Where("billing_date", ftime).
			Where("status_id", c.Query("status_id")).
			Find(&billing).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = "Show All"
		r.Data = &billing
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("billing_no") != "" && c.Query("billing_date") != "" {
		ftime, _ := time.Parse("2006-01-02", c.Query("billing_date"))
		if err := configs.Store.
			Preload("Status").
			Preload("VendorGroup").
			Preload("DocumentList.DocumentList").
			Preload("BillingStep").
			Where("billing_no like ?", "%"+c.Query("billing_no")+"%").
			Where("billing_date", ftime).
			Find(&billing).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = "Show All"
		r.Data = &billing
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("billing_no") != "" && c.Query("billing_no") != "null" {
		if err := configs.Store.
			Preload("Status").
			Preload("VendorGroup").
			Preload("DocumentList.DocumentList").
			Preload("BillingStep").
			Where("billing_no like ?", "%"+c.Query("billing_no")+"%").
			Find(&billing).Error; err != nil {
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
			Preload("DocumentList.DocumentList").
			Preload("BillingStep").
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
			Preload("DocumentList.DocumentList").
			Preload("BillingStep").
			Find(&billing, &models.Billing{VendorGroupID: c.Query("vendor_group")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Message = "Show All"
		r.Data = &billing
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("status_id") != "" {
		if err := configs.Store.
			Preload("Status").
			Preload("VendorGroup").
			Preload("DocumentList.DocumentList").
			Preload("BillingStep").
			Find(&billing, &models.Billing{StatusID: c.Query("status_id")}).Error; err != nil {
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
		Preload("DocumentList.DocumentList").
		Preload("BillingStep").
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
	if err := configs.Store.First(&status, &models.Status{Title: frm.StatusID}).Error; err != nil {
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
	billing.StatusID = status.ID
	billing.VendorGroupID = frm.VendorGroupID
	billing.IsActive = frm.IsActive

	if err := configs.Store.Create(&billing).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var step []models.StepTitle
	if err := configs.Store.Find(&step).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	for _, step := range step {
		var billinStep models.BillingStatus
		billinStep.StepTitleID = &step.ID
		billinStep.BillingID = &billing.ID
		if err := configs.Store.FirstOrCreate(&billinStep, &models.BillingStatus{
			StepTitleID: &step.ID,
			BillingID:   &billing.ID,
		}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
	}

	billing.Status = &status
	billing.VendorGroup = &vendorGroup
	r.Message = "Created successfully"
	r.Data = &billing
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func PutBilling(c *fiber.Ctx) error {
	var r models.Response
	var frm models.FrmUpdateBilling
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	var status models.Status
	if err := configs.Store.First(&status, &models.Status{Title: frm.Status}).Error; err != nil {
		r.Message = fmt.Sprintf("Status: %v", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// var vendorGroup models.VendorGroup
	// if err := configs.Store.First(&vendorGroup, &models.VendorGroup{ID: frm.VendorGroupID}).Error; err != nil {
	// 	r.Message = fmt.Sprintf("Vendor Group: %v", err.Error())
	// 	return c.Status(fiber.StatusNotFound).JSON(&r)
	// }

	var billing models.Billing
	if err := configs.Store.First(&billing, &models.Billing{ID: c.Params("id")}).Error; err != nil {
		r.Message = fmt.Sprintf("Notfound ID: %s", c.Params("id"))
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// billing.BillingNo = strings.ToUpper(frm.BillingNo)
	// billing.BillingDate = frm.BillingDate
	// billing.DueDate = frm.DueDate
	// billing.Amount = frm.Amount
	// billing.VendorCode = strings.ToUpper(frm.VendorCode)
	// billing.VendorName = strings.ToUpper(frm.VendorName)
	// if frm.PaymentDate != "" {
	// 	billing.PaymentDate = frm.PaymentDate
	// }
	// billing.Detail = frm.Detail
	billing.StatusID = status.ID
	// billing.VendorGroupID = frm.VendorGroupID
	// billing.IsActive = frm.IsActive

	if err := configs.Store.Save(&billing).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	/// Update billing step
	var step models.StepTitle
	if err := configs.Store.First(&step, &models.StepTitle{Title: frm.Step}).Error; err != nil {
		r.Message = fmt.Sprintf("Step Title %s", frm.Step)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var billingStatus models.BillingStatus
	billingStatus.StepTitleID = &step.ID
	billingStatus.BillingID = &billing.ID
	if err := configs.Store.FirstOrCreate(&billingStatus, &models.BillingStatus{
		StepTitleID: &step.ID,
		BillingID:   &billing.ID,
	}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	billingStatus.IsComplete = true
	if err := configs.Store.Save(&billingStatus).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	billing.Status = &status
	// billing.VendorGroup = &vendorGroup
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

	fileId, _ := g.New()
	fName := fmt.Sprintf("./public/import/%s-%s", fileId, file.Filename)
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

					// billing_no.GetString()
					// billing_date.GetString()
					// due_date.GetString()
					// amount.GetFloat64()
					// vendor_code.GetString()
					// vendor_name.GetString()
					// vendor_group_id.GetString()

					var status models.Status
					status.Title = "Open"
					if err := configs.Store.FirstOrCreate(&status, &models.Status{Title: "Open"}).Error; err != nil {
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
					billing.StatusID = status.ID
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
