package routes

import (
	"github.com/abe27/luckyapp/controllers"
	controller "github.com/abe27/luckyapp/controllers"
	"github.com/abe27/luckyapp/services"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(c *fiber.App) {
	c.Get("/", controller.HandlerHello)

	// Prefix API routes
	r := c.Group("/api/v1")

	// LogIn API routes
	u := r.Group("/auth")
	u.Post("/register", controller.Register)
	u.Post("/login", controllers.Login)

	auth := r.Use(services.AuthorizationRequired)
	auth.Get("/profile", controllers.Profile)

	// Master Role
	role := auth.Group("/role")
	role.Get("", controller.GetRole)
	role.Post("", controller.PostRole)
	role.Put("/:id", controller.PutRole)
	role.Delete("/:id", controller.DeleteRole)

	// Master Status
	status := auth.Group("/status")
	status.Get("", controller.GetStatus)
	status.Post("", controller.PostStatus)
	status.Put("/:id", controller.PutStatus)
	status.Delete("/:id", controller.DeleteStatus)

	// Master Step Title
	step := auth.Group("/step")
	step.Get("", controller.GetStep)
	step.Post("", controller.PostStep)
	step.Put("/:id", controller.PutStep)
	step.Delete("/:id", controller.DeleteStep)

	// Master Vendor Group
	vendor := auth.Group("/vendor")
	vendorDetail := vendor.Group("/detail")
	vendorDetail.Get("", controller.GetVendorDetail)
	vendorDetail.Post("", controller.PostVendorDetail)
	vendorDetail.Put("/:id", controller.PutVendorDetail)
	vendorDetail.Delete("/:id", controller.DeleteVendorDetail)

	vendorGroup := vendor.Group("/group")
	vendorGroup.Get("", controller.GetVendorGroup)
	vendorGroup.Post("", controller.PostVendorGroup)
	vendorGroup.Put("/:id", controller.PutVendorGroup)
	vendorGroup.Delete("/:id", controller.DeleteVendorGroup)

	// Master Permission
	permission := auth.Group("/permission")
	permission.Get("", controller.GetPermission)
	permission.Post("", controller.PostPermission)
	permission.Put("/:id", controller.PutPermission)
	permission.Delete("/:id", controller.DeletePermission)

	// Master Document
	documentList := auth.Group("/document/list")
	documentList.Get("", controller.GetDocumentList)
	documentList.Post("", controller.PostDocumentList)
	documentList.Put("/:id", controller.PutDocumentList)
	documentList.Delete("/:id", controller.DeleteDocumentList)

	// Billing
	billingList := auth.Group("/billing/list")
	billingList.Get("", controller.GetBilling)
	billingList.Post("", controller.PostBilling)
	billingList.Put("/:id", controller.PutBilling)
	billingList.Delete("/:id", controller.DeleteBilling)

	// Billing Document
	billingDocument := auth.Group("/billing/document")
	billingDocument.Get("", controller.GetBillingDocument)
	billingDocument.Post("", controller.PostBillingDocument)
	billingDocument.Put("/:id", controller.PutBillingDocument)
	billingDocument.Delete("/:id", controller.DeleteBillingDocument)

	// Billing Document
	billinStep := auth.Group("/billing/step")
	billinStep.Get("", controller.GetBillingStep)
	billinStep.Post("", controller.PostBillingStep)
	billinStep.Put("/:id", controller.PutBillingStep)
	billinStep.Delete("/:id", controller.DeleteBillingStep)
}
