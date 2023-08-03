package routers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jeanSagaz/go-api/internal/customer/application/dto"
	"github.com/jeanSagaz/go-api/internal/customer/application/services"
	"github.com/jeanSagaz/go-api/internal/customer/domain/entity"
	"github.com/jeanSagaz/go-api/internal/customer/infra/database"
	pkgEntity "github.com/jeanSagaz/go-api/pkg/entity"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func GinHandleRequests(Db *gorm.DB) {
	fmt.Println("Rest API v2.0 - gin Routers")

	db = Db
	router := gin.Default()
	router.GET("/customer/:id", getCustomerById)
	router.GET("/customer", getCustomers)
	router.POST("/customer", postCustomer)
	router.PUT("/customer/:id", putCustomer)
	router.DELETE("/customer/:id", deleteCustomer)

	log.Fatal(router.Run(":8080"), router)
}

func getCustomers(c *gin.Context) {
	service := getService()
	customers, err := service.GetAllCustomers()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, customers)
}

func getCustomerById(c *gin.Context) {
	id := c.Param("id")

	service := getService()
	customer, err := service.FindCustomer(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, customer)
}

func postCustomer(c *gin.Context) {
	var request dto.CustomerRequest

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&request); err != nil {
		return
	}

	newCustomer, errors := entity.NewCustomer(request.Name, request.Email)
	if errors != nil {
		//c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": errors})
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	service := getService()
	//customer, err := service.CustomerRepository.Insert(newCustomer)
	customer, err := service.AddCustomer(newCustomer)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": errors})
		return
	}

	c.IndentedJSON(http.StatusCreated, customer)
}

func putCustomer(c *gin.Context) {
	var request dto.CustomerRequest

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&request); err != nil {
		return
	}

	idParam := c.Param("id")
	id, err := pkgEntity.ParseID(idParam)
	if err != nil {
		var errors []pkgEntity.DomainNotification
		errors = append(errors, pkgEntity.DomainNotification{
			Key:   "",
			Value: err.Error(),
		})

		//c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": errors})
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	service := getService()
	customer, err := service.FindCustomer(idParam)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	customer.Id = id
	customer.Name = request.Name
	customer.Email = request.Email
	customer.UpdatedAt = time.Now()

	errors := customer.ValidateStruct()
	if errors != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	//customerChanged, err := service.CustomerRepository.Update(customer)
	customerChanged, err := service.ChangeCustomer(customer)
	if err != nil {
		//c.Status(http.StatusInternalServerError)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, customerChanged)
}

func deleteCustomer(c *gin.Context) {
	id := c.Param("id")

	service := getService()
	_, err := service.FindCustomer(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	//_, err = service.CustomerRepository.Delete(id)
	_, err = service.RemoveCustomer(id)
	if err != nil {
		//c.Status(http.StatusInternalServerError)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func getService() services.CustomerService {
	return services.CustomerService{
		ICustomerRepository: database.NewCustomerRepositoryDb(db),
	}

	// return services.CustomerService{
	// 	ICustomerRepository: database.CustomerRepositoryDb{Db: db},
	// }
}
