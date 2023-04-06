package controllers

import (
	"fmt"
	"log"

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
		log.Fatalln(err)
	}

	if status.Seq == 1 {
		var admin models.User
		if err := configs.Store.First(&admin, &models.User{IsAdmin: true}).Error; err != nil {
			log.Fatalln(err)
		}

		bodyTxt := fmt.Sprintf(`<h3>เรียน&nbsp;<strong style='color:blue;'>%s</strong></h3>
                                <p>
                                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ขณะนี้ทาง&nbsp;&nbsp;<span style='color:blue;'>%s</span>&nbsp;&nbsp;ได้ทำการอัพโหลดเอกสารประกอบการวางบิลเลขที่&nbsp;&nbsp;<span style='color:blue;'>%s</span>&nbsp;&nbsp;เรียบร้อยแล้ว
                                <br />
                                รบกวนตวรจสอบเอกสารที่ใช้ประกอบการวางบิลครับ/คะ
                                </p>
                                <p>ขอบคุณครับ/คะ</p>`, admin.FullName, "Vendor/Supplier", billing.BillingNo)
		// For test
		go services.SendMail("krumii.it@gmail.com", fmt.Sprintf("แจ้งผลการดำเนินงานการวางบิลเลขที่ %s", billing.BillingNo), bodyTxt)
	}

	for _, u := range user {
		var bodyTxt string
		switch status.Seq {
		case 0: //open
			bodyTxt = fmt.Sprintf("<h3>เรียน&nbsp;<strong style='color:blue;'>%s</strong></h3>"+
				"<p>"+
				"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;บริษัท ได้ทำการเปิดบิลหมายเลข&nbsp;<span style='color:blue;'>%s</span>"+
				"<br />"+
				"รบกวนแนบเอกสารประกอบการวางบิล เพื่อใช้ในการตรวจสอบ"+
				"</p>"+
				"<p>ขอบคุณครับ/คะ</p>", u.FullName, billing.BillingNo)

			go services.SendMail(u.Email, fmt.Sprintf("เปิดใบวางบิลใหม่เลขที่ %s", billing.BillingNo), bodyTxt)
		case 1: //on process

		case 2: //Verify
			bodyTxt = fmt.Sprintf("<h3>เรียน&nbsp;<strong style='color:blue;'>%s</strong></h3>"+
				"<p>"+
				"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;บริษัท ได้ทำการตรวจสอบเอกสารประกอบการวางบิลหมายเลข&nbsp;<span style='color:blue;'>%s</span>"+
				"&nbsp;เป็นที่เรียบร้อยแล้ว"+
				"<br />"+
				"<strong style='color:blue;'>เอกสารของท่านผ่านเกณฑ์ตามเงื่อนไขที่บริษัทกำหนด</strong>"+
				"</p>"+
				"<p>ขอบคุณครับ/คะ</p>", u.FullName, billing.BillingNo)

			go services.SendMail(u.Email, fmt.Sprintf("แจ้งผลการดำเนินงานการวางบิลเลขที่ %s", billing.BillingNo), bodyTxt)
		case 3: //Rejector
			bodyTxt = fmt.Sprintf("<h3>เรียน&nbsp;<strong style='color:blue;'>%s</strong></h3>"+
				"<p>"+
				"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;บริษัท ได้ทำการตรวจสอบเอกสารประกอบการวางบิลหมายเลข&nbsp;<span style='color:blue;'>%s</span>"+
				"&nbsp;<strong style='color:red;'>เอกสารของท่านไม่ผ่านเกณฑ์ตามเงื่อนไขที่บริษัทกำหนด</strong>"+
				"<br />"+
				"รบกวนทำการอัพโหลดเอกสารเอกสารประกอบการวางบิลอีกครั้ง"+
				"</p>"+
				"<p>ขอบคุณครับ/คะ</p>", u.FullName, billing.BillingNo)

			go services.SendMail(u.Email, fmt.Sprintf("แจ้งผลการดำเนินงานการวางบิลเลขที่ %s", billing.BillingNo), bodyTxt)
		case 4: //Approve
			bodyTxt = fmt.Sprintf("<h3>เรียน&nbsp;<strong style='color:blue;'>%s</strong></h3>"+
				"<p>"+
				"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;บริษัท ได้ทำการตรวจสอบข้อมูลและได้ทำการ<strong style='color:green;'>อนุมัติ</strong>บิลหมายเลข&nbsp;<span style='color:blue;'>%s</span>"+
				"&nbsp;เป็นที่เรียบร้อยแล้ว"+
				"</p>"+
				"<p>ขอบคุณครับ/คะ</p>", u.FullName, billing.BillingNo)

			go services.SendMail(u.Email, fmt.Sprintf("แจ้งผลการดำเนินงานการวางบิลเลขที่ %s", billing.BillingNo), bodyTxt)
		default: //Payment
			bodyTxt = fmt.Sprintf("<h3>เรียน&nbsp;<strong style='color:blue;'>%s</strong></h3>"+
				"<p>"+
				"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;บริษัท ได้ทำการตรวจสอบข้อมูลและได้ทำการ<strong style='color:green;'>อนุมัติ</strong>บิลหมายเลข&nbsp;<span style='color:blue;'>%s</span>"+
				"&nbsp;เป็นที่เรียบร้อยแล้ว"+
				"</p>"+
				"<p>ขอบคุณครับ/คะ</p>", u.FullName, billing.BillingNo)

			go services.SendMail(u.Email, fmt.Sprintf("แจ้งผลการดำเนินงานการวางบิลเลขที่ %s", billing.BillingNo), bodyTxt)
		}
	}
}
