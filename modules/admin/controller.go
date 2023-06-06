package admin

import "CRM/helpers"

type ControllerAdmin struct {
	UseCaseAdmin *UseCaseAdmin
}

func NewControllerAdmin(useCase *UseCaseAdmin) *ControllerAdmin {
	return &ControllerAdmin{
		UseCaseAdmin: useCase,
	}
}

type CreateResponseAdmin struct {
	Message string            `json:"message"`
	Data    AdminItemResponse `json:"data"`
}

type AdminItemResponse struct {
	ID       uint     `json:"id"`
	Username string   `json:"username"`
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
	admin := Actors{
		Username: req.Username,
		Password: hash,
		Verified: req.Verified,
		Active:   req.Active,
	}
	err := c.UseCaseAdmin.Create(&admin)
	if err != nil {
		return nil, err
	}

	res := &CreateResponseAdmin{
		Message: "Success",
		Data: AdminItemResponse{
			ID:       admin.ID,
			Username: admin.Username,
			Verified: admin.Verified,
			Active:   admin.Active,
		},
	}

	return res, nil
}

func (c ControllerAdmin) GetAllApprove() (*ReadAdminApprove, error) {
	customers, err := c.UseCaseAdmin.GetAll()
	if err != nil {
		return nil, err
	}
	result := &ReadAdminApprove{}

	for _, v := range customers {
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
	_, err := c.UseCaseAdmin.FindByUsername(&admin, username)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
func (c ControllerAdmin) FindByID(id string) (*RegisterApproval, error) {
	var approval RegisterApproval
	_, err := c.UseCaseAdmin.FindByID(&approval, id)
	if err != nil {
		return nil, err
	}
	return &approval, nil
}
func (c ControllerAdmin) FindActorByID(id string) (*Actors, error) {
	var actor Actors
	_, err := c.UseCaseAdmin.FindActorByID(&actor, id)
	if err != nil {
		return nil, err
	}
	return &actor, nil
}

func (c ControllerAdmin) UpdateApprove(req *RegisterApproval) error {
	return c.UseCaseAdmin.UpdateApprove(req)
}
func (c ControllerAdmin) UpdateActiveAdmin(actor *Actors) error {
	return c.UseCaseAdmin.UpdatedActive(actor)
}
