package handler

import (
	"book-keeping-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service *service.CustomerService
}

func NewCustomerHandler(s *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: s}
}

type CreateCustomerRequest struct {
	CustomerName string `json:"customer_name" binding:"required"`
}

// AddCustomer godoc
// @Summary Add a new customer
// @Description Add a new customer to the database
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body CreateCustomerRequest true "Customer Name"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/customers [post]
func (h *CustomerHandler) AddCustomer(c *gin.Context) {
	var req CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.AddCustomer(req.CustomerName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer added"})
}

// GetAllCustomers godoc
// @Summary Get all customers
// @Description Get a list of all customers
// @Tags customers
// @Produce json
// @Success 200 {array} model.Customer
// @Failure 500 {object} map[string]string
// @Router /api/customers [get]
func (h *CustomerHandler) GetAllCustomers(c *gin.Context) {
	customers, err := h.service.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}
