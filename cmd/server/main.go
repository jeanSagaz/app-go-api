package main

import (
	"log"

	"github.com/jeanSagaz/go-api/internal/handlers"
	"github.com/jeanSagaz/go-api/pkg/database"
)

var db database.Database

func init() {
	// SqlServer
	db.Server = "localhost"
	db.Port = 1434
	db.User = "sa"
	db.Password = "SqlServer2019!"
	db.Database = "poc"

	// Mysql
	// db.Server = "localhost"
	// db.Port = 3306
	// db.User = "root"
	// db.Password = ""
	// db.Database = "poc"

	db.AutoMigrateDb = true
}

func main() {
	dbConnection, err := db.Connect()
	if err != nil {
		log.Fatalf("error connecting to DB")
	}

	sqlDB, err := dbConnection.DB()
	if err != nil {
		log.Fatalf("error sql to DB")
	}
	defer sqlDB.Close()

	handlers.GinHandleRequests(dbConnection)
	//handlers.MuxHandleRequests()
	//handlers.ChiHandleRequests()
}
