package domain

import (
	productrepository "github.com/vanneeza/e-commerce-lite/internal/repository/product"
	storerepository "github.com/vanneeza/e-commerce-lite/internal/repository/store"
)

type ProductDomain struct {
	ProductRepository productrepository.ProductRepository
	StoreRepository   storerepository.StoreRepository
}
