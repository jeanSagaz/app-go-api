package interfaces

import "github.com/jeanSagaz/go-api/internal/customer/domain/entity"

type ICustomerRepository interface {
	GetAll() (*[]entity.Customer, error)
	FindById(id string) (*entity.Customer, error)
	Insert(customer *entity.Customer) (*entity.Customer, error)
	Update(customer *entity.Customer) (*entity.Customer, error)
	Delete(id string) (*entity.Customer, error)
}
