package services

import (
	"github.com/jeanSagaz/go-api/internal/customer/domain/entity"
	"github.com/jeanSagaz/go-api/internal/customer/domain/interfaces"
	"github.com/jeanSagaz/go-api/pkg/database"
)

type CustomerService struct {
	ICustomerRepository interfaces.ICustomerRepository
}

func NewCustomerService(
	ICustomerRepository interfaces.ICustomerRepository,
) *CustomerService {
	return &CustomerService{
		ICustomerRepository: ICustomerRepository,
	}
}

func (c *CustomerService) GetAllCustomers(pageSize, pageIndex int) (database.PagedResult, error) {
	pagedResult, err := c.ICustomerRepository.GetAll(pageSize, pageIndex)
	if err != nil {
		return database.PagedResult{}, err
	}

	return pagedResult, nil
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
