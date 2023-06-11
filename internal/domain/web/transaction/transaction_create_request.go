package transactionweb

type CreateOrderRequest struct {
	Qty        int16  `json:"qty" form:"qty"`
	ProductId  string `json:"product_id" form:"product_id"`
	CustomerId string `json:"customer_id"`
}
