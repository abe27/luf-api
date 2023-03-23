package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Status struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `gorm:"not null;column:title;index;unique;size:50" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (Status) TableName() string {
	return "tbt_status"
}

func (obj *Status) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Role struct {
	ID          string        `gorm:"primaryKey;size:21;" json:"id"`
	Title       string        `gorm:"not null;column:title;index;unique;size:50" json:"title" form:"title"`
	Description string        `json:"description" form:"description"`
	IsActive    bool          `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time     `json:"created_at" default:"now"`
	UpdatedAt   time.Time     `json:"updated_at" default:"now"`
	RoleDetail  []*RoleDetail `json:"role_detail"`
}

func (obj *Role) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type VendorGroup struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `gorm:"not null;column:title;index;unique;size:50" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *VendorGroup) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type DocumentList struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `gorm:"not null;column:title;index;unique;size:50" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *DocumentList) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Vendor struct {
	ID            string        `gorm:"primaryKey;size:21;" json:"id"`
	VendorGroupID *string       `json:"vendor_group_id" form:"vendor_group_id"`
	DocumentID    *string       `json:"document_id" form:"document_id"`
	RoleID        *string       `json:"role_id" form:"role_id"`
	IsActive      bool          `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt     time.Time     `json:"created_at" default:"now"`
	UpdatedAt     time.Time     `json:"updated_at" default:"now"`
	VendorGroup   *VendorGroup  `gorm:"foreignKey:VendorGroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"vendor_group"`
	DocumentList  *DocumentList `gorm:"foreignKey:DocumentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"document"`
	Role          *Role         `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"role"`
}

func (obj *Vendor) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Permission struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `gorm:"not null;column:title;index;unique;size:50" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	Read        bool      `gorm:"null" json:"read" form:"read" default:"false"`
	Write       bool      `gorm:"null" json:"write" form:"write" default:"false"`
	Create      bool      `gorm:"null" json:"create" form:"create" default:"false"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type RoleDetail struct {
	ID string `gorm:"primaryKey;size:21;" json:"id"`
	// Title        string     `gorm:"not null;column:title;index;unique;size:50" json:"title" form:"title"`
	RoleID       string     `json:"role_id" form:"role_id"`
	PermissionID string     `json:"permission_id" form:"permission_id"`
	Description  string     `json:"description" form:"description"`
	IsActive     bool       `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt    time.Time  `json:"created_at" default:"now"`
	UpdatedAt    time.Time  `json:"updated_at" default:"now"`
	Role         Role       `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"role"`
	Permission   Permission `gorm:"foreignKey:PermissionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"permission"`
}

func (obj *RoleDetail) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type User struct {
	ID        string    `gorm:"primaryKey;size:21;" json:"id"`
	UserName  string    `gorm:"not null;column:username;index;unique;size:50" json:"username" form:"username" binding:"required,min:5"`
	FullName  string    `json:"full_name" form:"full_name" binding:"required"`
	Email     string    `gorm:"not null;unique;size:50;" json:"email" form:"email" binding:"required"`
	Password  string    `gorm:"not null;unique;size:60;" json:"-" form:"password" binding:"required,min:6"`
	RoleID    *string   `json:"role_id" form:"role_id"`
	AvatarURL string    `json:"avatar_url" form:"avatar_url"`
	IsActive  bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
	Role      *Role     `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"role"`
}

func (obj *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Billing struct {
	ID            string       `gorm:"primaryKey;size:21;" json:"id"`
	BillingNo     string       `gorm:"not null;index;unique;size:50" json:"billing_no" form:"billing_no"`
	BillingDate   time.Time    `gorm:"type:date;" json:"billing_date" form:"billing_date"`
	DueDate       time.Time    `gorm:"type:date;" json:"due_date" form:"due_date"`
	Amount        float64      `json:"amount" form:"amount" default:"0.00"`
	VendorCode    string       `json:"vendor_code" form:"vendor_code"`
	VendorName    string       `json:"vendor_name" form:"vendor_name"`
	PaymentDate   time.Time    `gorm:"type:date;" json:"payment_date" form:"payment_date"`
	Detail        string       `json:"detail" form:"detail"`
	StatusID      *string      `json:"status_id" form:"status_id"`
	VendorGroupID *string      `json:"vendor_group_id" form:"vendor_group_id"`
	IsActive      bool         `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt     time.Time    `json:"created_at" default:"now"`
	UpdatedAt     time.Time    `json:"updated_at" default:"now"`
	Status        *Status      `gorm:"foreignKey:StatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"status"`
	VendorGroup   *VendorGroup `gorm:"foreignKey:VendorGroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"vendor_group"`
}

func (obj *Billing) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type BillingDocument struct {
	ID           string        `gorm:"primaryKey;size:21;" json:"id"`
	FileName     string        `gorm:"not null" json:"file_name" form:"file_name"`
	FileSize     float64       `json:"file_size" form:"file_size" default:"0.00"`
	FileType     string        `json:"file_type" form:"file_type" default:"pdf"`
	FilePath     string        `json:"file_path" form:"file_path"`
	BillingID    *string       `json:"billing_id" form:"billing_id"`
	DocumentID   *string       `json:"document_id" form:"document_id"`
	IsActive     bool          `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt    time.Time     `json:"created_at" default:"now"`
	UpdatedAt    time.Time     `json:"updated_at" default:"now"`
	Billing      *Billing      `gorm:"foreignKey:BillingID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"billing"`
	DocumentList *DocumentList `gorm:"foreignKey:DocumentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"document"`
}

func (obj *BillingDocument) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type StepTitle struct {
	ID          string    `gorm:"primaryKey;size:21;" json:"id"`
	Title       string    `gorm:"not null;column:title;index;unique;size:50" json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *StepTitle) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type BillingStatus struct {
	ID          string     `gorm:"primaryKey;size:21;" json:"id"`
	StepTitleID *string    `json:"step_title_id" form:"step_title_id"`
	BillingID   *string    `json:"billing_id" form:"billing_id"`
	IsComplete  bool       `gorm:"null" json:"is_complete" form:"is_complete" default:"false"`
	CreatedAt   time.Time  `json:"created_at" default:"now"`
	UpdatedAt   time.Time  `json:"updated_at" default:"now"`
	Billing     *Billing   `gorm:"foreignKey:BillingID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"billing"`
	StepTitle   *StepTitle `gorm:"foreignKey:StepTitleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"step_title"`
}

func (obj *BillingStatus) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type AuthSession struct {
	ID       string `json:"id"`
	Header   string `json:"header"`
	JwtType  string `json:"type"`
	JwtToken string `json:"token"`
	User     *User  `json:"user"`
}

type UserLoginForm struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}