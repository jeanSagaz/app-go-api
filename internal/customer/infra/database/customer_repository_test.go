package database_test

import (
	"log"
	"strings"
	"testing"

	"github.com/jeanSagaz/go-api/internal/customer/domain/entity"
	infraDatabase "github.com/jeanSagaz/go-api/internal/customer/infra/database"
	pkgDatabase "github.com/jeanSagaz/go-api/pkg/database"
	"github.com/stretchr/testify/require"
)

func TestCustomerRepositoryInsert(t *testing.T) {
	db := pkgDatabase.NewDbTest()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error connecting to DB")
	}

	defer sqlDB.Close()

	customer, _ := entity.NewCustomer("fulano", "fulano@gmail.com")

	// repo := infraDatabase.CustomerRepositoryDb {Db: db}
	repo := infraDatabase.NewCustomerRepositoryDb(db)
	newCustomer, err := repo.Insert(customer)

	var customerSaved entity.Customer
	result := db.Find(&customerSaved, "id = ?", customer.Id)
	// result := repo.Db.Raw("SELECT UPPER([Id]) AS Id, [Name] AS Name, [Email] AS Email, [Created_At] AS Created_At, [Updated_At] AS Updated_At FROM [poc].[dbo].[Customers] WHERE [Id] = ?", customer.Id).Scan(&customerSaved)

	require.NotEmpty(t, customerSaved.Id)
	require.Nil(t, err)
	require.NotNil(t, newCustomer)
	require.Equal(t, strings.ToUpper(customerSaved.Id.String()),
		strings.ToUpper(newCustomer.Id.String()),
		strings.ToUpper(customer.Id.String()))
	require.Equal(t, result.RowsAffected, int64(1))
}

func TestCustomerRepositoryDelete(t *testing.T) {
	db := pkgDatabase.NewDbTest()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error connecting to DB")
	}

	defer sqlDB.Close()

	customer, _ := entity.NewCustomer("fulano", "fulano@gmail.com")

	// repo := infraDatabase.CustomerRepositoryDb {Db: db}
	repo := infraDatabase.NewCustomerRepositoryDb(db)
	db.Create(&customer)
	_, err = repo.Delete(customer.Id.String())

	var customerDeleted entity.Customer
	result := db.Find(&customerDeleted, "id = ?", customer.Id)
	// result := repo.Db.Raw("SELECT UPPER([Id]) AS Id, [Name] AS Name, [Email] AS Email, [Created_At] AS Created_At, [Updated_At] AS Updated_At FROM [poc].[dbo].[Customers] WHERE [Id] = ?", customer.Id).Scan(&customerDeleted)

	require.Empty(t, customerDeleted.Id)
	require.Nil(t, err)
	require.NotEqual(t, customer.Id, customerDeleted.Id)
	require.Equal(t, result.RowsAffected, int64(0))
}
