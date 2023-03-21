package configs

import (
	"github.com/abe27/luckyapp/models"
	"gorm.io/gorm"
)

var (
	Store *gorm.DB
)

func SetDB() {
	if !Store.Migrator().HasTable(&models.Status{}) {
		Store.AutoMigrate(&models.Status{})
	}

	if !Store.Migrator().HasTable(&models.StepTitle{}) {
		Store.AutoMigrate(&models.StepTitle{})
	}

	if !Store.Migrator().HasTable(&models.Role{}) {
		Store.AutoMigrate(&models.Role{})
	}

	if !Store.Migrator().HasTable(&models.Permission{}) {
		Store.AutoMigrate(&models.Permission{})
	}

	if !Store.Migrator().HasTable(&models.RoleDetail{}) {
		Store.AutoMigrate(&models.RoleDetail{})
	}

	if !Store.Migrator().HasTable(&models.User{}) {
		Store.AutoMigrate(&models.User{})
	}

	if !Store.Migrator().HasTable(&models.VendorGroup{}) {
		Store.AutoMigrate(&models.VendorGroup{})
	}

	if !Store.Migrator().HasTable(&models.DocumentList{}) {
		Store.AutoMigrate(&models.DocumentList{})
	}

	if !Store.Migrator().HasTable(&models.Vendor{}) {
		Store.AutoMigrate(&models.Vendor{})
	}

	if !Store.Migrator().HasTable(&models.Billing{}) {
		Store.AutoMigrate(&models.Billing{})
	}

	if !Store.Migrator().HasTable(&models.BillingDocument{}) {
		Store.AutoMigrate(&models.BillingDocument{})
	}

	if !Store.Migrator().HasTable(&models.BillingStatus{}) {
		Store.AutoMigrate(&models.BillingStatus{})
	}
}
