package transactionweb

import (
	"time"

	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
)

type ResponseOrder struct {
	Id         string             `json:"id"`
	Qty        int16              `json:"qty"`
	ProductId  entity.Product     `json:"product"`
	CustomerId entity.Customer    `json:"customer"`
	StoreId    entity.Store       `json:"store"`
	Detail     entity.OrderDetail `json:"order_detail"`
}

type CallbackResponse struct {
	Id                 string    `json:"id"`
	DetailId           string    `json:"external_id"`
	UserId             string    `json:"user_id"`
	IsHigh             bool      `json:"is_high"`
	PaymentMethod      string    `json:"payment_method"`
	Status             string    `json:"status"`
	MerchantName       string    `json:"merchant_name"`
	Amount             float64   `json:"amount"`
	PaidAmount         float64   `json:"paid_amount"`
	BankCode           string    `json:"bank_code"`
	PaidDate           time.Time `json:"paid_at"`
	PayerEmail         string    `json:"payer_email"`
	Description        string    `json:"description"`
	Created            time.Time `json:"created"`
	Updated            time.Time `json:"updated"`
	Currency           string    `json:"currency"`
	PaymentChannel     string    `json:"payment_channel"`
	PaymentDestination string    `json:"payment_destination"`
}
