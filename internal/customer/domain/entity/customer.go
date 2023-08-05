package entity

import (
	"time"

	"github.com/go-playground/validator"
	pkgEntity "github.com/jeanSagaz/go-api/pkg/entity"
)

type Customer struct {
	// SqlServer
	Id        pkgEntity.ID `json:"id" validate:"-" gorm:"column:Id;type:uniqueidentifier;primary_key;not null"`
	Name      string       `json:"name" validate:"required" errormgs:"Name is required" gorm:"column:Name;type:varchar(100);not null"`
	Email     string       `json:"email" validate:"required" errormgs:"E-mail is required" gorm:"column:Email;type:varchar(100);not null"`
	CreatedAt time.Time    `json:"created_at" validate:"-" gorm:"column:Created_At;type:datetime2;not null"`
	UpdatedAt time.Time    `json:"updated_at" validate:"-" gorm:"column:Updated_At;type:datetime2;null"`

	// Mysql
	// Id        pkgEntity.ID `json:"id" valid:"-" gorm:"column:Id;type:varchar(36);primary_key;not null"`
	// Name      string       `json:"name" valid:"required~Name is required" gorm:"column:Name;type:varchar(100);not null"`
	// Email     string       `json:"email" valid:"email~E-mail is required" gorm:"column:Email;type:varchar(100);not null"`
	// CreatedAt time.Time    `json:"created_at" valid:"-" gorm:"column:Created_At;type:datetime(6);not null"`
	// UpdatedAt time.Time    `json:"updated_at" valid:"-" gorm:"column:Updated_At;type:datetime(6);null"`
}

func NewCustomer(name string, email string) (*Customer, []pkgEntity.DomainNotification) {
	customer := Customer{
		Name:  name,
		Email: email,
	}

	customer.prepare()

	err := customer.ValidateStruct()
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (customer *Customer) prepare() {
	customer.Id = pkgEntity.NewId()
	customer.CreatedAt = time.Now()
}

func (customer *Customer) Validator(validate *validator.Validate) error {
	return pkgEntity.ValidateFunc[Customer](*customer, validate)
}

func (customer *Customer) Validate() error {
	validate := validator.New()
	err := customer.Validator(validate)
	if err != nil {
		return err
	}

	return nil
}

func (customer *Customer) ValidateStruct() []pkgEntity.DomainNotification {
	validate := validator.New()
	err := pkgEntity.ValidateStruct[Customer](*customer, validate)

	if err != nil {
		return err
	}

	return nil
}
