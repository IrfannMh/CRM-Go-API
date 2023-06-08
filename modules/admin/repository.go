package admin

import "gorm.io/gorm"

type RepositoryAdmin struct {
	db *gorm.DB
}

func NewRepositoryAdmin(db *gorm.DB) *RepositoryAdmin {
	return &RepositoryAdmin{db: db}
}

func (r RepositoryAdmin) Save(admin *Actors) error {
	return r.db.Create(admin).Error
}

func (r RepositoryAdmin) GetAllApprove() ([]RegisterApproval, error) {
	var approvals []RegisterApproval
	err := r.db.Find(&approvals).Error
	return approvals, err
}

func (r RepositoryAdmin) GetAdminByUsername(username string) (Actors, error) {
	var admin Actors
	err := r.db.Find(&admin, "username = ?", username).Error
	return admin, err
}

func (r RepositoryAdmin) GetApprovalById(approval *RegisterApproval, id string) error {
	return r.db.First(&approval, "admin_id = ?", id).Error
}

func (r RepositoryAdmin) Update(approve *RegisterApproval) error {
	return r.db.Save(&approve).Error
}

func (r RepositoryAdmin) GetAdminById(id any) (Actors, error) {
	var actor Actors
	err := r.db.First(&actor, id).Error
	return actor, err
}

func (r RepositoryAdmin) UpdateActive(actor *Actors) error {
	return r.db.Save(&actor).Error
}

func (r RepositoryAdmin) CreateApproval(approve *RegisterApproval) error {
	return r.db.Create(approve).Error
}

func (r RepositoryAdmin) DeleteAdminById(data *Actors) error {
	return r.db.Delete(data).Error
}

func (r RepositoryAdmin) UpdateRole(a *Actors) error {
	return r.db.Save(a).Error
}
