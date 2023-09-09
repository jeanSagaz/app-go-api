package database

import (
	"fmt"
	"log"

	"github.com/jeanSagaz/go-api/internal/customer/domain/entity"

	"gorm.io/driver/sqlite"
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
	Test          bool
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Test = true
	dbInstance.AutoMigrateDb = true

	connection, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if !d.Test {
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			d.User,
			d.Password,
			d.Server,
			d.Port,
			d.Database)
		d.Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		// 	d.User,
		// 	d.Password,
		// 	d.Server,
		// 	d.Port,
		// 	d.Database)
		// d.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 	Logger: logger.Default.LogMode(logger.Info),
		// })
	} else {
		d.Db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	}

	if err != nil {
		panic(err)
	}

	if d.AutoMigrateDb {
		err = d.Db.AutoMigrate(&entity.Customer{})
		if err != nil {
			panic(err)
		}
	}

	return d.Db, nil
}
