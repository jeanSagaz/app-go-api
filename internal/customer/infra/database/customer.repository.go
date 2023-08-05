package database

import (
	"fmt"

	"github.com/jeanSagaz/go-api/internal/customer/domain/entity"

	"gorm.io/gorm"
)

type CustomerRepositoryDb struct {
	Db *gorm.DB
}

func NewCustomerRepositoryDb(db *gorm.DB) *CustomerRepositoryDb {
	return &CustomerRepositoryDb{Db: db}
}

func (repo CustomerRepositoryDb) GetAll() (*[]entity.Customer, error) {
	var customers []entity.Customer
	// result := repo.Db.Find(&customers)
	result := repo.Db.Raw("SELECT UPPER([Id]) AS Id, [Name] AS Name, [Email] AS Email, [Created_At] AS Created_At, [Updated_At] AS Updated_At FROM [poc].[dbo].[Customers]").Scan(&customers)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Customer does not exist")
	}

	fmt.Println(&customers)
	return &customers, nil
}

func (repo CustomerRepositoryDb) FindById(id string) (*entity.Customer, error) {
	var customer entity.Customer
	// result := repo.Db.First(&customer, "id = ?", id)
	// result := repo.Db.Find(&customer, "id = ?", id)
	// result := repo.Db.Where("id = ?", id).First(&customer)
	result := repo.Db.Raw("SELECT UPPER([Id]) AS Id, [Name] AS Name, [Email] AS Email, [Created_At] AS Created_At, [Updated_At] AS Updated_At FROM [poc].[dbo].[Customers] WHERE [Id] = ?", id).Scan(&customer)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Customer does not exist")
	}

	return &customer, nil
}

func (repo CustomerRepositoryDb) Insert(customer *entity.Customer) (*entity.Customer, error) {
	err := repo.Db.Create(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo CustomerRepositoryDb) Update(customer *entity.Customer) (*entity.Customer, error) {
	err := repo.Db.Save(&customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo CustomerRepositoryDb) Delete(id string) (*entity.Customer, error) {
	var customer entity.Customer

	// if err := repo.Db.Delete(&customer).Error; err != nil {
	if err := repo.Db.Where("id = ?", id).Delete(&entity.Customer{}).Error; err != nil {
		return nil, fmt.Errorf("error deleting customer with id: %s", id)
	}

	return &customer, nil
}
