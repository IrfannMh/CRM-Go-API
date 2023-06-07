package admin

import "gorm.io/gorm"

type Actors struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   Role   `json:"role_id"`
	Verified string `json:"verified"`
	Active   string `json:"active"`
}

//type Verified string
//
//const (
//	VerifiedTrue  Verified = "true"
//	VerifiedFalse Verified = "false"
//)
//
//type Active string
//
//const (
//	ActiveTrue  Active = "true"
//	ActiveFalse Active = "false"
//)

type Role struct {
	ID       uint   `json:"id"`
	RoleName string `json:"role_name"`
}

type RegisterApproval struct {
	ID           uint   `json:"id"`
	AdminID      uint   `json:"admin_id"`
	SuperAdminID uint   `json:"super_admin_id"`
	Status       string `json:"status"`
}
