package controllers

import (
	"fmt"

	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/models"
	"github.com/abe27/luckyapp/services"
	"github.com/gofiber/fiber/v2"
)

func TestSendMail(c *fiber.Ctx) error {
	var r models.Response
	go services.SendMail("krumii.it@gmail.com", "test billing", "บริษัท  ได้ทำการ อนุมัติ ฺBill  หมายเลข  00000000\nตามเอกสารท่านได้แนบมา\nท่านสามารถตรวจสอบสถานะ ของเอกสารได้ ที่  www.fffff.com")
	return c.Status(fiber.StatusOK).JSON(&r)
}

func SendMailBillingStatus(vendorID string, billing *models.Billing, status *models.Status) {
	var user []models.User
	if err := configs.Store.Find(&user, &models.User{VendorGroupID: &vendorID}).Error; err != nil {
		panic(err)
	}

	subject := "เปิดใบวางบิลใหม่"
	for _, u := range user {
		var bodyTxt string
		switch status.Seq {
		case 0: //open
			bodyTxt = fmt.Sprintf("เรียน <b>%s</b> \nบริษัท  ได้ทำการเปิดบิลหมายเลข %s\nรบกวนแนบเอกสารประกอบการวางบิลเพื่อใช้ในการตรวจสอบ\nขอบคุณครับ/คะ", u.FullName, billing.BillingNo)
		case 1: //on process
			bodyTxt = fmt.Sprintf("เรียน %s \nบริษัท  ได้ทำการเปิดบิลหมายเลข %s\nรบกวนแนบเอกสารประกอบการวางบิลเพื่อใช้ในการตรวจสอบ\nขอบคุณครับ/คะ", u.FullName, billing.BillingNo)
		case 2: //Verify
			bodyTxt = fmt.Sprintf("เรียน %s \nบริษัท  ได้ทำการเปิดบิลหมายเลข %s\nรบกวนแนบเอกสารประกอบการวางบิลเพื่อใช้ในการตรวจสอบ\nขอบคุณครับ/คะ", u.FullName, billing.BillingNo)
		case 3: //Rejectr
			bodyTxt = fmt.Sprintf("เรียน %s \nบริษัท  ได้ทำการเปิดบิลหมายเลข %s\nรบกวนแนบเอกสารประกอบการวางบิลเพื่อใช้ในการตรวจสอบ\nขอบคุณครับ/คะ", u.FullName, billing.BillingNo)
		case 4: //Approve
			bodyTxt = fmt.Sprintf("เรียน %s \nบริษัท  ได้ทำการเปิดบิลหมายเลข %s\nรบกวนแนบเอกสารประกอบการวางบิลเพื่อใช้ในการตรวจสอบ\nขอบคุณครับ/คะ", u.FullName, billing.BillingNo)

		default: //Payment
			bodyTxt = ""
		}

		go services.SendMail(u.Email, subject, bodyTxt)
	}
}
