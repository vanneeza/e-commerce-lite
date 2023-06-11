package stockweb

import (
	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
	storeweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/store"
)

type StockResponse struct {
	Id      string                       `json:"id"`
	Stock   int                          `json:"stock"`
	Product entity.Product               `json:"product"`
	Store   storeweb.StoreResponseGlobal `json:"store"`
}
