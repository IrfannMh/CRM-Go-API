package admin

type Actors struct {
	Id       uint     `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	RoleID   uint     `json:"role_id"`
	Verified Verified `json:"verified"`
	Active   Active   `json:"active"`
}

type Verified string

const (
	VerifiedTrue  Verified = "true"
	VerifiedFalse Verified = "false"
)

type Active string

const (
	ActiveTrue  Active = "true"
	ActiveFalse Active = "false"
)

type Role struct {
	Id       uint   `json:"id"`
	RoleName string `json:"role_name"`
}

type RegisterApproval struct {
	Id           uint   `json:"id"`
	AdminID      uint   `json:"admin_id"`
	SuperAdminID uint   `json:"super_admin_id"`
	Status       string `json:"status"`
}
