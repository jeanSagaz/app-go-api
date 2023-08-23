package interfaces

import (
	"github.com/jeanSagaz/go-api/internal/customer/domain/entity"
	"github.com/jeanSagaz/go-api/pkg/database"
)

type ICustomerRepository interface {
	GetAll(pageSize, pageIndex int) (database.PagedResult, error)
	FindById(id string) (*entity.Customer, error)
	Insert(customer *entity.Customer) (*entity.Customer, error)
	Update(customer *entity.Customer) (*entity.Customer, error)
	Delete(id string) (*entity.Customer, error)
}
