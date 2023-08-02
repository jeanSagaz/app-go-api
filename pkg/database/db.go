package database

import (
	"fmt"
	"log"

	"github.com/jeanSagaz/go-sample/internal/customer/domain/entity"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db            *gorm.DB
	Server        string
	Port          int
	User          string
	Password      string
	Database      string
	AutoMigrateDb bool
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Server = "localhost"
	dbInstance.Port = 1434
	dbInstance.User = "sa"
	dbInstance.Password = "SqlServer2019!"
	dbInstance.Database = "poc"

	connection, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		d.User,
		d.Password,
		d.Server,
		d.Port,
		d.Database)
	d.Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&entity.Customer{})
	}

	return d.Db, nil
}
