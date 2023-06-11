package productcontroller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanneeza/e-commerce-lite/internal/domain/web"
	productweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/product"
	productservice "github.com/vanneeza/e-commerce-lite/internal/service/product"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type ProductControllerImpl struct {
	ProductService productservice.ProductService
}

func NewProductController(productService productservice.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

// Register implements ProductController.
func (s *ProductControllerImpl) Register(ctx *gin.Context) {
	var product productweb.ProductCreateRequest
	err := ctx.ShouldBindJSON(&product)
	helper.PanicError(err)

	log.Println(product)
	productResponse, err := s.ProductService.Register(ctx.Request.Context(), product)
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "Congratulations, you have successfully created a product.",
		Data:    productResponse,
	}

	ctx.JSON(http.StatusCreated, gin.H{"product": webResponse})
}

// GetAll implements ProductController.
func (s *ProductControllerImpl) GetAll(ctx *gin.Context) {
	productResponse, _ := s.ProductService.GetAll(ctx.Request.Context())
	if len(productResponse) == 0 {
		webResponse := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: "product data not found",
			Data:    "null",
		}
		ctx.JSON(http.StatusOK, gin.H{"product": webResponse})
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all product data",
		Data:    productResponse,
	}
	ctx.JSON(http.StatusOK, gin.H{"product": webResponse})
}

// GetById implements ProductController.
func (s *ProductControllerImpl) GetById(ctx *gin.Context) {
	productId := ctx.Param("id")
	productResponse, _ := s.ProductService.GetById(ctx.Request.Context(), productId)
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "data product from id " + productId,
		Data:    productResponse,
	}
	ctx.JSON(http.StatusOK, gin.H{"product": webResponse})
}

// Edit implements ProductController.
func (s *ProductControllerImpl) Edit(ctx *gin.Context) {
	var product productweb.ProductUpdateRequest
	err := ctx.ShouldBindJSON(&product)
	helper.PanicError(err)

	productId := ctx.Param("id")
	product.Id = productId

	productResponse, err := s.ProductService.Edit(ctx.Request.Context(), product)
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "Congratulations, you have successfully update a product.",
		Data:    productResponse,
	}

	ctx.JSON(http.StatusCreated, gin.H{"product": webResponse})
}

// Unreg implements ProductController.
func (s *ProductControllerImpl) Unreg(ctx *gin.Context) {

	productId := ctx.Param("id")

	productResponse, err := s.ProductService.Unreg(ctx.Request.Context(), productId)
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "CREATED",
		Message: "Congratulations, you have successfully delete a product.",
		Data:    productResponse,
	}

	ctx.JSON(http.StatusOK, gin.H{"product": webResponse})
}
