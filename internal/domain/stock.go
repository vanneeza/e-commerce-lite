package domain

import (
	productrepository "github.com/vanneeza/e-commerce-lite/internal/repository/product"
	stockrepository "github.com/vanneeza/e-commerce-lite/internal/repository/stock"
	storerepository "github.com/vanneeza/e-commerce-lite/internal/repository/store"
)

type StockDomain struct {
	StockRepository   stockrepository.StockRepository
	ProductRepository productrepository.ProductRepository
	StoreRepository   storerepository.StoreRepository
}
