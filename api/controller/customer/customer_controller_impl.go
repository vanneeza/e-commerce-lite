package customercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanneeza/e-commerce-lite/internal/domain/web"
	customerweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/customer"
	customerservice "github.com/vanneeza/e-commerce-lite/internal/service/customer"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type CustomerControllerImpl struct {
	CustomerService customerservice.CustomerService
}

func NewCustomerController(customerService customerservice.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

// Register implements CustomerController.
func (s *CustomerControllerImpl) Register(ctx *gin.Context) {
	var customer customerweb.CustomerCreateRequest
	err := ctx.ShouldBindJSON(&customer)
	helper.PanicError(err)

	customerResponse, err := s.CustomerService.Register(ctx.Request.Context(), customer)
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "Congratulations, you have successfully created a customer.",
		Data:    customerResponse,
	}

	ctx.JSON(http.StatusCreated, gin.H{"customer": webResponse})
}

// GetAll implements CustomerController.
func (s *CustomerControllerImpl) GetAll(ctx *gin.Context) {
	customerResponse, _ := s.CustomerService.GetAll(ctx.Request.Context())
	if len(customerResponse) == 0 {
		webResponse := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: "customer data not found",
			Data:    "null",
		}
		ctx.JSON(http.StatusOK, gin.H{"customer": webResponse})
		return
	}
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all customer data",
		Data:    customerResponse,
	}
	ctx.JSON(http.StatusOK, gin.H{"customer": webResponse})
}

// GetById implements CustomerController.
func (s *CustomerControllerImpl) GetById(ctx *gin.Context) {
	customerId := ctx.Param("id")
	customerResponse, _ := s.CustomerService.GetByParams(ctx.Request.Context(), customerId, "")
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all customer data",
		Data:    customerResponse,
	}
	ctx.JSON(http.StatusOK, gin.H{"customer": webResponse})
}

// Edit implements CustomerController.
func (s *CustomerControllerImpl) Edit(ctx *gin.Context) {
	var customer customerweb.CustomerUpdateRequest
	err := ctx.ShouldBindJSON(&customer)
	helper.PanicError(err)

	customerId := ctx.Param("id")

	customer.Id = customerId

	customerResponse, err := s.CustomerService.Edit(ctx.Request.Context(), customer)
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "Congratulations, you have successfully update a customer.",
		Data:    customerResponse,
	}

	ctx.JSON(http.StatusCreated, gin.H{"customer": webResponse})
}

// Unreg implements CustomerController.
func (s *CustomerControllerImpl) Unreg(ctx *gin.Context) {

	customerId := ctx.Param("id")

	customerResponse, err := s.CustomerService.Unreg(ctx.Request.Context(), customerId, "")
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "CREATED",
		Message: "Congratulations, you have successfully delete a customer.",
		Data:    customerResponse,
	}

	ctx.JSON(http.StatusOK, gin.H{"customer": webResponse})
}
