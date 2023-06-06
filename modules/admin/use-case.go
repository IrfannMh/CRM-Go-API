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
func (u UseCaseAdmin) FindByID(approval *RegisterApproval, id string) (RegisterApproval, error) {
	return u.repo.GetAdminById(approval, id)
}
func (u UseCaseAdmin) FindActorByID(actors *Actors, id string) (Actors, error) {
	return u.repo.GetActorById(actors, id)
}

func (u UseCaseAdmin) UpdateApprove(approval *RegisterApproval) error {
	return u.repo.Update(approval)
}

func (u UseCaseAdmin) UpdatedActive(actor *Actors) error {
	return u.repo.UpdateActive(actor)
}
