package customers

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u UseCase) Create(customer *Customer) error {
	return u.repo.Save(customer)
}

func (u UseCase) GetAll() ([]Customer, error) {
	return u.repo.GetAllCustomer()
}

func (u UseCase) Delete(customer *Customer) error {
	return u.repo.Delete(customer)
}

func (u UseCase) FindById(customer *Customer, id string) error {
	return u.repo.FindById(customer, id)
}
