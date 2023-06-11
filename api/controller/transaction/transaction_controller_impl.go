package transactioncontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vanneeza/e-commerce-lite/internal/domain/web"
	transactionweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/transaction"
	transactionservice "github.com/vanneeza/e-commerce-lite/internal/service/transaction"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type TransactionControllerImpl struct {
	TxService transactionservice.TransactionService
}

// MakeOrder implements TransactionController.
func (tc *TransactionControllerImpl) MakeOrder(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	id := claims["id"].(string)

	var txCreateRequest transactionweb.CreateOrderRequest
	errBind := ctx.ShouldBindJSON(&txCreateRequest)
	if errBind != nil {
		webResponse := web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "sorry, server was down",
			Data:    "NULL",
		}
		ctx.JSON(http.StatusCreated, gin.H{"transaction": webResponse})
		return
	}

	txCreateRequest.CustomerId = id
	orderResponse, errService := tc.TxService.MakeOrder(ctx.Request.Context(), txCreateRequest)

	if errService != nil {
		webResponse := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Failed to create order",
			Data:    "NULL",
		}
		ctx.JSON(http.StatusCreated, gin.H{"transaction": webResponse})
		return
	}

	customerAddress := xendit.InvoiceCustomerAddress{
		Country:     "Indonesia",
		StreetLine1: orderResponse.CustomerId.Address,
	}

	customer := xendit.InvoiceCustomer{
		GivenNames:   orderResponse.CustomerId.Name,
		Email:        orderResponse.CustomerId.Email,
		MobileNumber: orderResponse.CustomerId.NoHp,
		Address:      []xendit.InvoiceCustomerAddress{customerAddress},
	}

	item := xendit.InvoiceItem{
		Name:     orderResponse.ProductId.Name,
		Quantity: int(txCreateRequest.Qty),
		Price:    float64(orderResponse.ProductId.Price),
	}

	items := []xendit.InvoiceItem{
		item,
	}

	data := invoice.CreateParams{
		ExternalID:  orderResponse.Detail.Id,
		Amount:      orderResponse.Detail.TotalPrice,
		Description: "Learning Create Invoice With Xendit",
		Customer:    customer,
		Currency:    "IDR",
		Items:       items,
	}

	resp, errInvoice := invoice.CreateWithContext(ctx.Request.Context(), &data)
	if errInvoice != nil {
		webResponse := web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Failed to create invoice",
			Data:    "NULL",
		}
		ctx.JSON(http.StatusCreated, gin.H{"transaction": webResponse})
		return
	}

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "Congratulations, you have successfully maked an order. Waiting to payment",
		Data:    resp,
	}
	ctx.JSON(http.StatusCreated, gin.H{"transaction": webResponse})
	// DSyLrivHlFcg4AgYWzWMT8F2gAW6WszcWragVZJWqtMcTwb9 ( Callback )
}

func (tc *TransactionControllerImpl) XenditCallback(ctx *gin.Context) {
	var callbackData transactionweb.CallbackResponse
	err := ctx.BindJSON(&callbackData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jsonData, err := json.MarshalIndent(callbackData, "", "  ")
	if err != nil {
		log.Println("Error marshaling callback data:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	log.Println("Callback Data:")
	log.Println(string(jsonData))

	errSaveCallback := tc.TxService.SaveCallBack(ctx.Request.Context(), callbackData)
	if errSaveCallback != nil {
		log.Println("error save callback:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	} else {
		log.Println("success save callback")
		log.Println("update status detail with ID: " + callbackData.DetailId)
	}
}

func NewTransactionController(txService transactionservice.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		TxService: txService,
	}
}
