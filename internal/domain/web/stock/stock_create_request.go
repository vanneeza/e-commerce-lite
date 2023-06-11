package stockweb

type StockCreateRequest struct {
	Stock     int    `json:"stock"`
	ProductId string `json:"product_id"`
}
