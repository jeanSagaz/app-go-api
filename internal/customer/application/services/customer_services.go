package services

import (
	"github.com/jeanSagaz/go-api/internal/customer/domain/entity"
	"github.com/jeanSagaz/go-api/internal/customer/domain/interfaces"
	"github.com/jeanSagaz/go-api/pkg/database"
)

type CustomerServices struct {
	ICustomerRepository interfaces.ICustomerRepository
}

func NewCustomerServices(
	ICustomerRepository interfaces.ICustomerRepository,
) *CustomerServices {
	return &CustomerServices{
		ICustomerRepository: ICustomerRepository,
	}
}

func (c *CustomerServices) GetAllCustomers(pageSize, pageIndex int) (database.PagedResult, error) {
	pagedResult, err := c.ICustomerRepository.GetAll(pageSize, pageIndex)
	if err != nil {
		return database.PagedResult{}, err
	}

	return pagedResult, nil
}

func (c *CustomerServices) FindCustomer(id string) (*entity.Customer, error) {
	customer, err := c.ICustomerRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *CustomerServices) AddCustomer(customer *entity.Customer) (*entity.Customer, error) {
	customer, err := c.ICustomerRepository.Insert(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *CustomerServices) ChangeCustomer(customer *entity.Customer) (*entity.Customer, error) {
	customer, err := c.ICustomerRepository.Update(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *CustomerServices) RemoveCustomer(id string) (*entity.Customer, error) {
	customer, err := c.ICustomerRepository.Delete(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
