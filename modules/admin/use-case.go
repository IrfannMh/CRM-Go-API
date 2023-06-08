package admin

type UseCaseAdmin struct {
	repo *RepositoryAdmin
}

func NewUseCaseAdmin(repo *RepositoryAdmin) *UseCaseAdmin {
	return &UseCaseAdmin{
		repo: repo,
	}
}

func (u UseCaseAdmin) Create(admin *Actors) error {
	return u.repo.Save(admin)
}

func (u UseCaseAdmin) GetAll() ([]RegisterApproval, error) {
	return u.repo.GetAllApprove()
}
func (u UseCaseAdmin) FindByUsername(admin *Actors, username string) (Actors, error) {
	return u.repo.GetAdminByUsername(admin, username)
}
func (u UseCaseAdmin) FindByID(approval *RegisterApproval, id string) error {
	return u.repo.GetApprovalById(approval, id)
}
func (u UseCaseAdmin) FindAdminByID(actors *Actors, id any) error {
	return u.repo.GetAdminById(actors, id)
}

func (u UseCaseAdmin) UpdateApprove(approval *RegisterApproval) error {
	return u.repo.Update(approval)
}

func (u UseCaseAdmin) UpdatedActive(actor *Actors) error {
	return u.repo.UpdateActive(actor)
}

func (u UseCaseAdmin) CreateApproval(approve *RegisterApproval) error {
	return u.repo.CreateApproval(approve)
}

func (u UseCaseAdmin) DeleteAdminById(data *Actors) error {
	return u.repo.DeleteAdminById(data)
}

func (u UseCaseAdmin) UpdateRole(a *Actors) error {
	return u.repo.UpdateRole(a)
}
