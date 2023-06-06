package customers

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) Save(customer *Customer) error {
	return r.db.Create(customer).Error
}

func (r Repository) GetAllCustomer() ([]Customer, error) {
	var customers []Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r Repository) Delete(customer *Customer) error {
	return r.db.Delete(&customer).Error
}

func (r Repository) FindById(customer *Customer, id string) error {
	return r.db.First(&customer, id).Error
}
