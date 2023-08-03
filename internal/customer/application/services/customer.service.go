package services

import (
	"github.com/jeanSagaz/go-api/internal/customer/domain/entity"
	"github.com/jeanSagaz/go-api/internal/customer/domain/interfaces"
)

type CustomerService struct {
	Customer            *entity.Customer
	ICustomerRepository interfaces.ICustomerRepository
}

func (c *CustomerService) GetAllCustomers() (*[]entity.Customer, error) {
	customers, err := c.ICustomerRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (c *CustomerService) FindCustomer(id string) (*entity.Customer, error) {
	customer, err := c.ICustomerRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *CustomerService) AddCustomer(customer *entity.Customer) (*entity.Customer, error) {
	customer, err := c.ICustomerRepository.Insert(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *CustomerService) ChangeCustomer(customer *entity.Customer) (*entity.Customer, error) {
	customer, err := c.ICustomerRepository.Update(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *CustomerService) RemoveCustomer(id string) (*entity.Customer, error) {
	customer, err := c.ICustomerRepository.Delete(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
