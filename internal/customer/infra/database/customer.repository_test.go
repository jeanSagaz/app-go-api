package database_test

import (
	"log"
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

	// customerSaved, _ := repo.FindById(newCustomer.Id)
	var customerSaved entity.Customer
	result := db.Find(&customerSaved, "id = ?", customer.Id)

	require.NotEmpty(t, customerSaved.Id)
	require.Nil(t, err)
	require.NotNil(t, newCustomer)
	// require.Equal(t, customerSaved.Id, newCustomer.Id, customer.Id)
	require.Equal(t, result.RowsAffected, int64(1))
}
