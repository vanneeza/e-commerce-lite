package domain

import (
	customerrepository "github.com/vanneeza/e-commerce-lite/internal/repository/customer"
	productrepository "github.com/vanneeza/e-commerce-lite/internal/repository/product"
	storerepository "github.com/vanneeza/e-commerce-lite/internal/repository/store"
	transactionrepostiory "github.com/vanneeza/e-commerce-lite/internal/repository/transaction"
)

type TxRepository struct {
	TxRepository       transactionrepostiory.TrasactionRepository
	CustomerRepository customerrepository.CustomerRepository
	ProductRepository  productrepository.ProductRepository
	StoreRepository    storerepository.StoreRepository
}
