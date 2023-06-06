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

func (r RepositoryAdmin) GetAdminByUsername(admin *Actors, username string) (Actors, error) {
	err := r.db.First(&admin, "username = ?", username).Error
	return *admin, err
}

func (r RepositoryAdmin) GetAdminById(approval *RegisterApproval, id string) (RegisterApproval, error) {
	err := r.db.First(&approval, "admin_id = ?", id).Error
	return *approval, err
}

func (r RepositoryAdmin) Update(approve *RegisterApproval) error {
	return r.db.Save(&approve).Error
}

func (r RepositoryAdmin) GetActorById(actors *Actors, id string) (Actors, error) {
	err := r.db.First(&actors, id).Error
	return *actors, err
}

func (r RepositoryAdmin) UpdateActive(actor *Actors) error {
	return r.db.Save(&actor).Error
}
