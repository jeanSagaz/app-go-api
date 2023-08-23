package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jeanSagaz/go-api/internal/customer/application/dto"
	"github.com/jeanSagaz/go-api/internal/customer/application/services"
	"github.com/jeanSagaz/go-api/internal/customer/domain/entity"
	"github.com/jeanSagaz/go-api/internal/customer/domain/interfaces"
	pkgEntity "github.com/jeanSagaz/go-api/pkg/entity"
)

type GinHandler struct {
	ICustomerRepository interfaces.ICustomerRepository
}

func NewGinHandler(
	ICustomerRepository interfaces.ICustomerRepository,
) *GinHandler {
	return &GinHandler{
		ICustomerRepository: ICustomerRepository,
	}
}

func (h *GinHandler) GetCustomers(c *gin.Context) {
	pageSize := c.DefaultQuery("pageSize", "10")
	pageIndex := c.DefaultQuery("pageIndex", "1")
	// pageIndex := c.Query("pageIndex")

	ps, _ := strconv.Atoi(pageSize)
	pi, _ := strconv.Atoi(pageIndex)
	service := services.NewCustomerServices(h.ICustomerRepository)
	pagedResult, err := service.GetAllCustomers(ps, pi)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, pagedResult)
}

func (h *GinHandler) GetCustomerById(c *gin.Context) {
	id := c.Param("id")

	service := services.NewCustomerServices(h.ICustomerRepository)
	customer, err := service.FindCustomer(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, customer)
}

func (h *GinHandler) PostCustomer(c *gin.Context) {
	var request dto.CustomerRequest

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&request); err != nil {
		return
	}

	newCustomer, errors := entity.NewCustomer(request.Name, request.Email)
	if errors != nil {
		// c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": errors})
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	service := services.NewCustomerServices(h.ICustomerRepository)
	// customer, err := service.CustomerRepository.Insert(newCustomer)
	customer, err := service.AddCustomer(newCustomer)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": errors})
		return
	}

	c.IndentedJSON(http.StatusCreated, customer)
}

func (h *GinHandler) PutCustomer(c *gin.Context) {
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

		// c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": errors})
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	service := services.NewCustomerServices(h.ICustomerRepository)
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

	// customerChanged, err := service.CustomerRepository.Update(customer)
	customerChanged, err := service.ChangeCustomer(customer)
	if err != nil {
		// c.Status(http.StatusInternalServerError)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, customerChanged)
}

func (h *GinHandler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	service := services.NewCustomerServices(h.ICustomerRepository)
	_, err := service.FindCustomer(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	// _, err = service.CustomerRepository.Delete(id)
	_, err = service.RemoveCustomer(id)
	if err != nil {
		// c.Status(http.StatusInternalServerError)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
