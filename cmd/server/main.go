package main

import (
	"log"

	"github.com/gin-gonic/gin"
	infraDatabase "github.com/jeanSagaz/go-api/internal/customer/infra/database"
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
	go handlers.MuxHandleRequests()
	go handlers.ChiHandleRequests()

	dbConnection, err := db.Connect()
	if err != nil {
		log.Fatalf("error connecting to DB")
	}

	sqlDB, err := dbConnection.DB()
	if err != nil {
		log.Fatalf("error sql to DB")
	}
	defer sqlDB.Close()

	customerRepository := infraDatabase.NewCustomerRepositoryDb(dbConnection)
	ginHandler := handlers.NewGinHandler(customerRepository)

	router := gin.Default()
	router.GET("/customer/:id", ginHandler.GetCustomerById)
	router.GET("/customer", ginHandler.GetCustomers)
	router.POST("/customer", ginHandler.PostCustomer)
	router.PUT("/customer/:id", ginHandler.PutCustomer)
	router.DELETE("/customer/:id", ginHandler.DeleteCustomer)

	log.Fatal(router.Run(":8080"), router)
}
