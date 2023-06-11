package productweb

type ProductUpdateRequest struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	StoreId string `json:"store_id"`
}
