package productweb

import (
	storeweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/store"
)

type ProductResponse struct {
	Id    string                       `json:"id"`
	Name  string                       `json:"name"`
	Price int                          `json:"price"`
	Store storeweb.StoreResponseGlobal `json:"store"`
}
