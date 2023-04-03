package controllers

import (
	"fmt"
	"strings"

	"github.com/abe27/luckyapp/configs"
	"github.com/abe27/luckyapp/models"
	"github.com/abe27/luckyapp/services"
	"github.com/gofiber/fiber/v2"
)

func TestGetUserID(c *fiber.Ctx) error {
	var r models.Response
	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	obj, err := services.ValidateToken(token)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	r.Data = &obj
	r.Message = token
	return c.Status(fiber.StatusOK).JSON(&r)
}

func Register(c *fiber.Ctx) error {
	var r models.Response
	var frm models.User
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	if len(strings.TrimSpace(frm.UserName)) <= 0 {
		r.Message = "Please enter UserName!"
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	if len(strings.TrimSpace(frm.Password)) <= 0 {
		r.Message = "Please enter Password!"
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	if len(strings.TrimSpace(strings.ToLower(frm.Email))) <= 0 {
		r.Message = "Please enter Email!"
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	if len(frm.FullName) <= 0 {
		r.Message = "Please enter Full Name!"
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	password := services.HashingPassword(frm.Password)
	isMatch := services.CheckPasswordHashing(frm.Password, password)

	if !isMatch {
		r.Message = "เกิดข้อผิดพลาดหร่างการเข้ารหัส"
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var role models.Role
	if err := configs.Store.Find(&role, &models.Role{Title: *frm.RoleID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	fmt.Println(*frm.VendorGroupID)
	var vendor models.VendorGroup
	if err := configs.Store.Find(&vendor, &models.VendorGroup{Title: *frm.VendorGroupID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var user models.User
	user.UserName = strings.ToLower(frm.UserName)
	user.FullName = frm.FullName
	user.Email = strings.ToLower(frm.Email)
	user.Company = frm.Company
	user.Password = password
	user.RoleID = &role.ID
	user.VendorGroupID = &vendor.ID

	// Upload GEDI File To Directory
	file, err := c.FormFile("avatar")
	if err == nil {
		fName := fmt.Sprintf("./public/user/%s", file.Filename)
		if err := c.SaveFile(file, fName); err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		user.AvatarURL = fmt.Sprintf("/user/%s", file.Filename)
	}
	if err := configs.Store.Create(&user).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Data = &user
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func Verify(c *fiber.Ctx) error {
	var r models.Response
	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	if _, err := services.ValidateToken(token); err != nil {
		r.Message = err.Error()
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&r)
}

func Login(c *fiber.Ctx) error {
	var r models.Response
	var user models.UserLoginForm
	if err := c.BodyParser(&user); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	// Check AuthorizationRequired
	db := configs.Store
	var userData models.User
	if err := db.Preload("Role").Preload("VendorGroup").Where("username=?", user.UserName).First(&userData).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	isMatched := services.CheckPasswordHashing(c.FormValue("password"), userData.Password)
	if !isMatched {
		r.Message = "Password not match!"
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// Create Token
	auth := services.CreateToken(userData)
	r.Message = "Auth success!"
	r.Data = &auth
	return c.Status(fiber.StatusOK).JSON(&r)
}

// func Verify(c *fiber.Ctx) error {
// 	var r models.Response
// 	s := c.Get("Authorization")
// 	token := strings.TrimPrefix(s, "Bearer ")
// 	return c.Status(fiber.StatusOK).JSON(&r)
// }

func Profile(c *fiber.Ctx) error {
	var r models.Response

	return c.Status(fiber.StatusOK).JSON(&r)
}

func GetMember(c *fiber.Ctx) error {
	var r models.Response
	if c.Query("id") != "" {
		var member models.User
		if err := configs.Store.Preload("Role").Preload("VendorGroup").Find(&member, &models.User{ID: c.Query("id")}).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Data = &member
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("username") != "" {
		var member []models.User
		if err := configs.Store.Preload("Role").Preload("VendorGroup").Where("username like ?", "%"+c.Query("username")+"%").Find(&member).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Data = &member
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("name") != "" {
		var member []models.User
		if err := configs.Store.Preload("Role").Preload("VendorGroup").Where("full_name like ?", "%"+c.Query("name")+"%").Find(&member).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Data = &member
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	if c.Query("role") != "" {
		var member []models.User
		if err := configs.Store.Preload("Role").Preload("VendorGroup").Where("full_name like ?", "%"+c.Query("name")+"%").Find(&member).Error; err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusNotFound).JSON(&r)
		}
		r.Data = &member
		return c.Status(fiber.StatusOK).JSON(&r)
	}

	var member []models.User
	if err := configs.Store.Preload("Role").Preload("VendorGroup").Find(&member).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = "Show all data"
	r.Data = &member
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateMember(c *fiber.Ctx) error {
	var r models.Response
	var frm models.User
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	if len(frm.UserName) <= 0 {
		r.Message = "Please enter UserName!"
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	if len(frm.Password) <= 0 {
		r.Message = "Please enter Password!"
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	if len(frm.Email) <= 0 {
		r.Message = "Please enter Email!"
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	if len(frm.FullName) <= 0 {
		r.Message = "Please enter Full Name!"
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	password := services.HashingPassword(frm.Password)
	isMatch := services.CheckPasswordHashing(frm.Password, password)

	if !isMatch {
		r.Message = "เกิดข้อผิดพลาดหร่างการเข้ารหัส"
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var role models.Role
	if err := configs.Store.Find(&role, &models.Role{Title: *frm.RoleID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var user models.User
	if err := configs.Store.First(&user, &models.User{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// Upload GEDI File To Directory
	if file, err := c.FormFile("avatar"); err == nil {
		fName := fmt.Sprintf("./public/user/%s", file.Filename)
		if err := c.SaveFile(file, fName); err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		user.AvatarURL = fmt.Sprintf("/user/%s", file.Filename)
	}

	user.UserName = frm.UserName
	user.FullName = frm.FullName
	user.Email = frm.Email
	user.Company = frm.Company
	user.Password = password
	user.RoleID = &role.ID

	if err := configs.Store.Save(&user).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = "Update data successfully."
	r.Data = &user
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func DeleteMember(c *fiber.Ctx) error {
	var r models.Response
	var user models.User
	if err := configs.Store.First(&user, &models.User{ID: c.Params("id")}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := configs.Store.Delete(&user).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = "Delete success!"
	return c.Status(fiber.StatusOK).JSON(&r)
}
