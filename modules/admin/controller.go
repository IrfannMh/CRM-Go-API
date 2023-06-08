package admin

import (
	"CRM/helpers"
)

type ControllerAdmin struct {
	useCaseAdmin *UseCaseAdmin
}

func NewControllerAdmin(useCase *UseCaseAdmin) *ControllerAdmin {
	return &ControllerAdmin{
		useCaseAdmin: useCase,
	}
}

type CreateResponseAdmin struct {
	Message string             `json:"message"`
	Data    ItemResponseCreate `json:"data"`
}

type ItemResponse struct {
	Id       uint     `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	RoleId   uint     `json:"role_id"`
	Verified Verified `json:"verified"`
	Active   Active   `json:"active"`
}
type ItemResponseCreate struct {
	Id       uint     `json:"id"`
	Username string   `json:"username"`
	RoleId   uint     `json:"role_id"`
	Verified Verified `json:"verified"`
	Active   Active   `json:"active"`
}

type ReadAdminApprove struct {
	Data []ItemApprove `json:"data"`
}

type ItemApprove struct {
	ID      uint   `json:"id"`
	AdminID uint   `json:"admin_id"`
	Status  string `json:"status"`
}

func (c ControllerAdmin) Create(req *CreateAdminRequest) (*CreateResponseAdmin, error) {
	hash := helpers.HashPass(req.Password)
	_ = VerifiedTrue
	admin := Actors{
		Username: req.Username,
		Password: hash,
		RoleID:   3,
		Verified: VerifiedFalse,
		Active:   ActiveFalse,
	}
	err := c.useCaseAdmin.Create(&admin)
	if err != nil {
		return nil, err
	}

	approve := RegisterApproval{
		AdminID: admin.Id,
	}
	err = c.useCaseAdmin.CreateApproval(&approve)
	if err != nil {
		return nil, err
	}
	res := &CreateResponseAdmin{
		Message: "Create Admin Success",
		Data: ItemResponseCreate{
			Id:       admin.Id,
			Username: admin.Username,
			RoleId:   admin.RoleID,
			Verified: admin.Verified,
			Active:   admin.Active,
		},
	}

	return res, nil
}

func (c ControllerAdmin) GetAllApprove() (*ReadAdminApprove, error) {
	listApprove, err := c.useCaseAdmin.GetAll()
	if err != nil {
		return nil, err
	}
	result := &ReadAdminApprove{}

	for _, v := range listApprove {
		item := ItemApprove{
			ID:      v.Id,
			AdminID: v.AdminID,
			Status:  v.Status,
		}
		result.Data = append(result.Data, item)
	}
	return result, nil
}

func (c ControllerAdmin) FindByUsername(username string) (*ItemResponse, error) {

	data, err := c.useCaseAdmin.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	result := &ItemResponse{
		Id:       data.Id,
		Username: data.Username,
		Password: data.Password,
		RoleId:   data.RoleID,
		Verified: data.Verified,
		Active:   data.Active,
	}
	return result, nil
}
func (c ControllerAdmin) FindApprovalID(id string) (*RegisterApproval, error) {
	var approval RegisterApproval
	err := c.useCaseAdmin.FindByID(&approval, id)
	if err != nil {
		return nil, err
	}
	return &approval, nil
}
func (c ControllerAdmin) FindAdminByID(id string) (*Actors, error) {
	//return c.useCaseAdmin.FindAdminByID(admin, id)
	data, err := c.useCaseAdmin.FindAdminByID(id)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c ControllerAdmin) UpdateApprove(req *RegisterApproval) error {

	id := req.AdminID
	data, err := c.useCaseAdmin.FindAdminByID(id)
	if err != nil {
		return err
	}
	data.RoleID = 1
	if err := c.useCaseAdmin.UpdateRole(&data); err != nil {
		return err
	}
	return c.useCaseAdmin.UpdateApprove(req)
}
func (c ControllerAdmin) UpdateActiveAdmin(actor *Actors) error {
	return c.useCaseAdmin.UpdatedActive(actor)
}

func (c ControllerAdmin) DeleteAdminById(data *Actors) error {
	return c.useCaseAdmin.DeleteAdminById(data)
}
