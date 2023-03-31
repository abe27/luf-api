package configs

import (
	"encoding/json"

	"github.com/abe27/luckyapp/models"
	"github.com/abe27/luckyapp/services"
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

	if !Store.Migrator().HasTable(&models.BillingRequiredDocument{}) {
		Store.AutoMigrate(&models.BillingRequiredDocument{})
	}

	if !Store.Migrator().HasTable(&models.VendorGroupHistory{}) {
		Store.AutoMigrate(&models.VendorGroupHistory{})
	}
}

func DBSeed() {
	// Seed Permision
	// Open our jsonFile
	data, _ := services.ReadJson("public/mock/permission.json")
	var perms []models.Permission
	json.Unmarshal(data, &perms)

	for _, p := range perms {
		if err := Store.FirstOrCreate(&p, &models.Permission{Title: p.Title}).Error; err != nil {
			panic(err)
		}
	}

	statusData, _ := services.ReadJson("public/mock/status.json")
	var sData []models.Status
	json.Unmarshal(statusData, &sData)
	for _, s := range sData {
		if err := Store.FirstOrCreate(&s, &models.Status{Title: s.Title}).Error; err != nil {
			panic(err)
		}
	}

	docData, _ := services.ReadJson("public/mock/document_list.json")
	var docs []models.DocumentList
	json.Unmarshal(docData, &docs)
	for _, s := range docs {
		if err := Store.FirstOrCreate(&s, &models.DocumentList{Title: s.Title}).Error; err != nil {
			panic(err)
		}
	}

	roleData, _ := services.ReadJson("public/mock/role.json")
	var role []models.Role
	json.Unmarshal(roleData, &role)
	for _, s := range role {
		if err := Store.FirstOrCreate(&s, &models.Role{Title: s.Title}).Error; err != nil {
			panic(err)
		}
	}

	vgData, _ := services.ReadJson("public/mock/vendor_group.json")
	var vendGroup []models.VendorGroup
	json.Unmarshal(vgData, &vendGroup)
	for _, s := range vendGroup {
		if err := Store.FirstOrCreate(&s, &models.VendorGroup{Title: s.Title}).Error; err != nil {
			panic(err)
		}
	}

	objRoleDetail, _ := services.ReadJson("public/mock/role_detail.json")
	var roleDetail []models.RoleDetail
	json.Unmarshal(objRoleDetail, &roleDetail)
	for _, s := range roleDetail {
		var r models.Role
		if err := Store.First(&r, &models.Role{Title: s.RoleID}).Error; err != nil {
			panic(err)
		}

		var p models.Permission
		if err := Store.First(&p, &models.Permission{Title: s.PermissionID}).Error; err != nil {
			panic(err)
		}

		s.RoleID = s.ID
		s.PermissionID = p.ID
		if err := Store.FirstOrCreate(&s, &models.RoleDetail{RoleID: r.ID, PermissionID: p.ID}).Error; err != nil {
			panic(err)
		}
	}

	stData, _ := services.ReadJson("public/mock/step_title.json")
	var stepTitle []models.StepTitle
	json.Unmarshal(stData, &stepTitle)
	for _, s := range stepTitle {
		if err := Store.FirstOrCreate(&s, &models.StepTitle{Title: s.Title}).Error; err != nil {
			panic(err)
		}
	}

	vendorData, _ := services.ReadJson("public/mock/vendor.json")
	var vendor []models.Vendor
	json.Unmarshal(vendorData, &vendor)
	for _, s := range vendor {
		var g models.VendorGroup
		if err := Store.First(&g, &models.VendorGroup{Title: *s.VendorGroupID}).Error; err != nil {
			panic(err)
		}

		var doc models.DocumentList
		if err := Store.First(&doc, &models.DocumentList{Title: *s.DocumentID}).Error; err != nil {
			panic(err)
		}

		var v models.Role
		if err := Store.First(&v, &models.Role{Title: *s.RoleID}).Error; err != nil {
			panic(err)
		}

		if err := Store.FirstOrCreate(&s, &models.Vendor{VendorGroupID: &g.ID, DocumentID: &doc.ID, RoleID: &v.ID}).Error; err != nil {
			panic(err)
		}
	}

	userData, _ := services.ReadJson("public/mock/user.json")
	var user []models.FrmUserSeed
	json.Unmarshal(userData, &user)
	for _, u := range user {
		password := services.HashingPassword(u.Password)
		isMatch := services.CheckPasswordHashing(u.Password, password)
		if isMatch {
			var r models.Role
			if err := Store.First(&r, &models.Role{Title: *u.RoleID}).Error; err != nil {
				panic(err)
			}
			u.RoleID = &r.ID

			var vendorGroup models.VendorGroup
			if err := Store.First(&vendorGroup, &models.VendorGroup{Title: *u.VendorGroupID}).Error; err != nil {
				panic(err)
			}
			u.VendorGroupID = &vendorGroup.ID
			u.Password = password

			var user models.User
			user.UserName = u.UserName
			user.FullName = u.FullName
			user.Email = u.Email
			user.Company = u.Company
			user.Password = u.Password
			user.RoleID = u.RoleID
			user.VendorGroupID = u.VendorGroupID
			// user.AvatarURL = nil
			user.IsActive = true
			if err := Store.FirstOrCreate(&user, &models.User{UserName: u.UserName}).Error; err != nil {
				panic(err)
			}
		}
	}
}
