package stockcontroller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanneeza/e-commerce-lite/internal/domain/web"
	stockweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/stock"
	stockservice "github.com/vanneeza/e-commerce-lite/internal/service/stock"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type StockControllerImpl struct {
	StockService stockservice.StockService
}

func NewStockController(stockService stockservice.StockService) StockController {
	return &StockControllerImpl{
		StockService: stockService,
	}
}

// Register implements StockController.
func (s *StockControllerImpl) Register(ctx *gin.Context) {
	var stock stockweb.StockCreateRequest
	err := ctx.ShouldBindJSON(&stock)
	helper.PanicError(err)

	log.Println(stock)
	stockResponse, err := s.StockService.Register(ctx.Request.Context(), stock)
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "Congratulations, you have successfully created a stock.",
		Data:    stockResponse,
	}

	ctx.JSON(http.StatusCreated, gin.H{"stock": webResponse})
}

// GetAll implements StockController.
func (s *StockControllerImpl) GetAll(ctx *gin.Context) {
	stockResponse, _ := s.StockService.GetAll(ctx.Request.Context())
	if len(stockResponse) == 0 {
		webResponse := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: "stock data not found",
			Data:    "null",
		}
		ctx.JSON(http.StatusOK, gin.H{"stock": webResponse})
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all stock data",
		Data:    stockResponse,
	}
	ctx.JSON(http.StatusOK, gin.H{"stock": webResponse})
}

// GetById implements StockController.
func (s *StockControllerImpl) GetById(ctx *gin.Context) {
	stockId := ctx.Param("id")
	stockResponse, _ := s.StockService.GetById(ctx.Request.Context(), stockId)
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "data stock from id " + stockId,
		Data:    stockResponse,
	}
	ctx.JSON(http.StatusOK, gin.H{"stock": webResponse})
}

// Edit implements StockController.
func (s *StockControllerImpl) Edit(ctx *gin.Context) {
	var stock stockweb.StockUpdateRequest
	err := ctx.ShouldBindJSON(&stock)
	helper.PanicError(err)

	stockId := ctx.Param("id")
	stock.Id = stockId

	stockResponse, err := s.StockService.Edit(ctx.Request.Context(), stock)
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "Congratulations, you have successfully update a stock.",
		Data:    stockResponse,
	}

	ctx.JSON(http.StatusCreated, gin.H{"stock": webResponse})
}

// Unreg implements StockController.
func (s *StockControllerImpl) Unreg(ctx *gin.Context) {

	stockId := ctx.Param("id")

	stockResponse, err := s.StockService.Unreg(ctx.Request.Context(), stockId)
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "CREATED",
		Message: "Congratulations, you have successfully delete a stock.",
		Data:    stockResponse,
	}

	ctx.JSON(http.StatusOK, gin.H{"stock": webResponse})
}
