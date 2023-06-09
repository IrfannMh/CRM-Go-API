package customers

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}
type RepositoryInterface interface {
	Save(customer *Customer) error
	GetAllCustomer(page int, email string) ([]Customer, error)
	Delete(customer *Customer) error
	FindById(customer *Customer, id string) error
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &repository{db: db}
}

func (r repository) Save(customer *Customer) error {
	return r.db.Create(customer).Error
}

func (r repository) GetAllCustomer(page int, email string) ([]Customer, error) {
	var customers []Customer
	err := r.db.Limit(6).Offset(page).Find(&customers, "email = ?", email).Error
	return customers, err
}

func (r repository) Delete(customer *Customer) error {
	return r.db.Delete(&customer).Error
}

func (r repository) FindById(customer *Customer, id string) error {
	return r.db.First(&customer, id).Error
}
