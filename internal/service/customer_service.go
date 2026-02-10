package service

import (
	"book-keeping-backend/internal/model"
	"book-keeping-backend/internal/repository"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) AddCustomer(name string) error {
	customer := &model.Customer{CustomerName: name}
	return s.repo.Create(customer)
}

func (s *CustomerService) GetAllCustomers() ([]model.Customer, error) {
	return s.repo.GetAll()
}
