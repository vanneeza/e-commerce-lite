package productweb

type ProductCreateRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Store string `json:"store_id"`
}
