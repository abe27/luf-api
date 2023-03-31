package routes

import (
	"github.com/abe27/luckyapp/controllers"
	"github.com/abe27/luckyapp/services"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(c *fiber.App) {
	c.Get("/", controllers.HandlerHello)

	// Prefix API routes
	r := c.Group("/api/v1")

	// Test sendMail
	r.Get("/test", controllers.TestSendMail)
	r.Get("/token", controllers.TestGetUserID)

	// LogIn API routes
	u := r.Group("/auth")
	u.Post("/register", controllers.Register)
	u.Post("/login", controllers.Login)

	auth := r.Use(services.AuthorizationRequired)
	auth.Get("/profile", controllers.Profile)

	// Master Member
	member := auth.Group("/member")
	member.Get("", controllers.GetMember)
	member.Put("/:id", controllers.UpdateMember)
	member.Delete("/:id", controllers.DeleteMember)

	// Master Role
	role := auth.Group("/role")
	role.Get("", controllers.GetRole)
	role.Post("", controllers.PostRole)
	role.Put("/:id", controllers.PutRole)
	role.Delete("/:id", controllers.DeleteRole)

	// Master Status
	status := auth.Group("/status")
	status.Get("", controllers.GetStatus)
	status.Post("", controllers.PostStatus)
	status.Put("/:id", controllers.PutStatus)
	status.Delete("/:id", controllers.DeleteStatus)

	// Master Step Title
	step := auth.Group("/step")
	step.Get("", controllers.GetStep)
	step.Post("", controllers.PostStep)
	step.Put("/:id", controllers.PutStep)
	step.Delete("/:id", controllers.DeleteStep)

	// Master Vendor Group
	vendor := auth.Group("/vendor")
	vendorDetail := vendor.Group("/detail")
	vendorDetail.Get("", controllers.GetVendorDetail)
	vendorDetail.Post("", controllers.PostVendorDetail)
	vendorDetail.Put("/:id", controllers.PutVendorDetail)
	vendorDetail.Delete("/:id", controllers.DeleteVendorDetail)

	vendorGroup := vendor.Group("/group")
	vendorGroup.Get("", controllers.GetVendorGroup)
	vendorGroup.Post("", controllers.PostVendorGroup)
	vendorGroup.Put("/:id", controllers.PutVendorGroup)
	vendorGroup.Delete("/:id", controllers.DeleteVendorGroup)

	// Master Permission
	permission := auth.Group("/permission")
	permission.Get("", controllers.GetPermission)
	permission.Post("", controllers.PostPermission)
	permission.Put("/:id", controllers.PutPermission)
	permission.Delete("/:id", controllers.DeletePermission)

	// Master Document
	documentList := auth.Group("/document/list")
	documentList.Get("", controllers.GetDocumentList)
	documentList.Post("", controllers.PostDocumentList)
	documentList.Put("/:id", controllers.PutDocumentList)
	documentList.Delete("/:id", controllers.DeleteDocumentList)

	// Billing
	billingListImport := auth.Group("/billing/import")
	billingListImport.Post("", controllers.ImportBilling)
	billingList := auth.Group("/billing/list")
	billingList.Get("", controllers.GetBilling)
	billingList.Post("", controllers.PostBilling)
	billingList.Put("/:id", controllers.PutBilling)
	billingList.Delete("/:id", controllers.DeleteBilling)

	// Billing Document
	billingDocumentUpload := auth.Group("/billing/upload")
	billingDocumentUpload.Post("", controllers.UploadBillingDocument)
	billingDocumentUpload.Delete("/:id", controllers.DeleteUploadBillingDocument)
	billingDocument := auth.Group("/billing/document")
	billingDocument.Get("", controllers.GetBillingDocument)
	billingDocument.Post("", controllers.PostBillingDocument)
	billingDocument.Put("/:id", controllers.PutBillingDocument)
	billingDocument.Delete("/:id", controllers.DeleteBillingDocument)

	// Billing Document
	billinStep := auth.Group("/billing/step")
	billinStep.Get("", controllers.GetBillingStep)
	billinStep.Post("", controllers.PostBillingStep)
	billinStep.Put("/:id", controllers.PutBillingStep)
	billinStep.Delete("/:id", controllers.DeleteBillingStep)

	// Billing Document
	billinRequire := auth.Group("/billing/require")
	billinRequire.Get("", controllers.GetBillingRequireDocument)
	billinRequire.Post("/:id", controllers.PostBillingRequireDocument)
	billinRequire.Put("/:id", controllers.PutBillingRequireDocument)
	billinRequire.Delete("/:id", controllers.DeleteBillingRequireDocument)
	// Billing Approve
	auth.Put("/billing/approve/:id", controllers.BillingApprove)
}
