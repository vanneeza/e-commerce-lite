package stockweb

type StockUpdateRequest struct {
	Id        string `json:"id"`
	Stock     int    `json:"stock"`
	ProductId string `json:"product_id"`
}
