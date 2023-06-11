package entity

import (
	"time"
)

type Order struct {
	Id         string
	Qty        int16
	ProductId  Product
	CustomerId Customer
	StoreId    Store
	TxDetailId OrderDetail
}

type OrderDetail struct {
	Id         string
	TotalPrice float64
	Status     string
	BuyDate    time.Time
}

type Callback struct {
	Id                 string
	DetailId           string
	UserId             string
	IsHigh             bool
	PaymentMethod      string
	Status             string
	MerchantName       string
	Amount             float64
	PaidAmount         float64
	BankCode           string
	PaidDate           time.Time
	PayerEmail         string
	Description        string
	Created            time.Time
	Updated            time.Time
	Currency           string
	PaymentChannel     string
	PaymentDestination string
}
