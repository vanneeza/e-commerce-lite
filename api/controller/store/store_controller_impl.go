package storecontroller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanneeza/e-commerce-lite/internal/domain/web"
	storeweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/store"
	storeservice "github.com/vanneeza/e-commerce-lite/internal/service/store"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type StoreControllerImpl struct {
	StoreService storeservice.StoreService
}

func NewStoreController(storeService storeservice.StoreService) StoreController {
	return &StoreControllerImpl{
		StoreService: storeService,
	}
}

// Register implements StoreController.
func (s *StoreControllerImpl) Register(ctx *gin.Context) {
	var store storeweb.StoreCreateRequest
	err := ctx.ShouldBindJSON(&store)
	helper.PanicError(err)

	log.Println(store, "di controller")

	storeResponse, err := s.StoreService.Register(ctx.Request.Context(), store)
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "Congratulations, you have successfully created a store.",
		Data:    storeResponse,
	}

	ctx.JSON(http.StatusCreated, gin.H{"store": webResponse})
}

// GetAll implements StoreController.
func (s *StoreControllerImpl) GetAll(ctx *gin.Context) {
	storeResponse, _ := s.StoreService.GetAll(ctx.Request.Context())
	if len(storeResponse) == 0 {
		webResponse := web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: "store data not found",
			Data:    "null",
		}
		ctx.JSON(http.StatusOK, gin.H{"store": webResponse})
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all store data",
		Data:    storeResponse,
	}
	ctx.JSON(http.StatusOK, gin.H{"store": webResponse})
}

// GetById implements StoreController.
func (s *StoreControllerImpl) GetById(ctx *gin.Context) {
	storeId := ctx.Param("id")
	storeResponse, _ := s.StoreService.GetById(ctx.Request.Context(), storeId)
	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "list of all store data",
		Data:    storeResponse,
	}
	ctx.JSON(http.StatusOK, gin.H{"store": webResponse})
}

// Edit implements StoreController.
func (s *StoreControllerImpl) Edit(ctx *gin.Context) {
	var store storeweb.StoreUpdateRequest
	err := ctx.ShouldBindJSON(&store)
	helper.PanicError(err)

	storeId := ctx.Param("id")

	store.Id = storeId

	storeResponse, err := s.StoreService.Edit(ctx.Request.Context(), store)
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "Congratulations, you have successfully update a store.",
		Data:    storeResponse,
	}

	ctx.JSON(http.StatusCreated, gin.H{"store": webResponse})
}

// Unreg implements StoreController.
func (s *StoreControllerImpl) Unreg(ctx *gin.Context) {

	storeId := ctx.Param("id")

	storeResponse, err := s.StoreService.Unreg(ctx.Request.Context(), storeId)
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Status:  "CREATED",
		Message: "Congratulations, you have successfully delete a store.",
		Data:    storeResponse,
	}

	ctx.JSON(http.StatusOK, gin.H{"store": webResponse})
}
