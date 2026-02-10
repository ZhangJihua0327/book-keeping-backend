package repository

import (
	"book-keeping-backend/internal/model"
)

type CustomerRepository struct{}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (r *CustomerRepository) Create(customer *model.Customer) error {
	return DB.Create(customer).Error
}

func (r *CustomerRepository) GetAll() ([]model.Customer, error) {
	var customers []model.Customer
	err := DB.Find(&customers).Error
	return customers, err
}
