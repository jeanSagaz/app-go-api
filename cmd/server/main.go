package main

import (
	"log"

	"github.com/jeanSagaz/go-sample/internal/routers"
	"github.com/jeanSagaz/go-sample/pkg/database"
)

var db database.Database

func init() {
	db.Server = "localhost"
	db.Port = 1434
	db.User = "sa"
	db.Password = "SqlServer2019!"
	db.Database = "poc"
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

	routers.GinHandleRequests(dbConnection)
	//routers.MuxHandleRequests()
	//routers.ChiHandleRequests()
}
