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
	Message string       `json:"message"`
	Data    ItemResponse `json:"data"`
}

type ItemResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
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
	admin := Actors{
		Username: req.Username,
		Password: hash,
	}
	err := c.useCaseAdmin.Create(&admin)
	if err != nil {
		return nil, err
	}

	approve := RegisterApproval{
		AdminID: admin.ID,
	}
	err = c.useCaseAdmin.CreateApproval(&approve)
	if err != nil {
		return nil, err
	}
	res := &CreateResponseAdmin{
		Message: "Success",
		Data: ItemResponse{
			ID:       admin.ID,
			Username: admin.Username,
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
			ID:      v.ID,
			AdminID: v.AdminID,
			Status:  v.Status,
		}
		result.Data = append(result.Data, item)
	}
	return result, nil
}

func (c ControllerAdmin) FindByUsername(username string) (*Actors, error) {
	var admin Actors
	_, err := c.useCaseAdmin.FindByUsername(&admin, username)
	if err != nil {
		return nil, err
	}
	return &admin, nil
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
	var admin Actors
	if err := c.useCaseAdmin.FindAdminByID(&admin, id); err != nil {
		return nil, err
	}
	return &admin, nil
}

func (c ControllerAdmin) UpdateApprove(req *RegisterApproval) error {
	var admin Actors
	id := req.AdminID
	if err := c.useCaseAdmin.UpdateRole(&admin); err != nil {
		return err
	}
	if err := c.useCaseAdmin.FindAdminByID(&admin, id); err != nil {
		return err
	}
	admin.RoleID = 1
	return c.useCaseAdmin.UpdateApprove(req)
}
func (c ControllerAdmin) UpdateActiveAdmin(actor *Actors) error {
	return c.useCaseAdmin.UpdatedActive(actor)
}

func (c ControllerAdmin) DeleteAdminById(data *Actors) error {
	return c.useCaseAdmin.DeleteAdminById(data)
}
