package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	pkgEntity "github.com/jeanSagaz/go-sample/pkg/entity"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Customer struct {
	Id        pkgEntity.ID `json:"id" valid:"-" gorm:"column:Id;type:uniqueidentifier;primary_key;not null"`
	Name      string       `json:"name" valid:"required~Name is required" gorm:"column:Name;type:varchar(100);not null"`
	Email     string       `json:"email" valid:"email~E-mail is required" gorm:"column:Email;type:varchar(100);not null"`
	CreatedAt time.Time    `json:"created_at" valid:"-" gorm:"column:Created_At;type:datetime2;not null"`
	UpdatedAt time.Time    `json:"updated_at" valid:"-" gorm:"column:Updated_At;type:datetime2;null"`
}

func NewCustomer(name string, email string) (*Customer, error) {
	customer := Customer{
		Name:  name,
		Email: email,
	}

	customer.prepare()

	err := customer.Validate()
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (customer *Customer) prepare() {
	customer.Id = pkgEntity.NewId()
	customer.CreatedAt = time.Now()
}

func (customer *Customer) Validate() error {
	_, err := govalidator.ValidateStruct(customer)

	if err != nil {
		return err
	}

	return nil
}
