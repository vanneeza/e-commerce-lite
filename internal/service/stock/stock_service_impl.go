package stockservice

import (
	"context"
	"database/sql"
	"log"

	"github.com/segmentio/ksuid"
	"github.com/vanneeza/e-commerce-lite/internal/domain"
	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
	stockweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/stock"
	storeweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/store"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type StockServiceImpl struct {
	Db          *sql.DB
	StockDomain domain.StockDomain
}

func NewStockService(db *sql.DB, stockDomain domain.StockDomain) StockService {
	return &StockServiceImpl{
		Db:          db,
		StockDomain: stockDomain,
	}

}

func (s *StockServiceImpl) Register(ctx context.Context, req stockweb.StockCreateRequest) (stockweb.StockResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	product, _ := s.StockDomain.ProductRepository.FindById(ctx, tx, req.ProductId)
	store, _ := s.StockDomain.StoreRepository.FindById(ctx, tx, product.Store.Id)

	stock := entity.Stock{
		Id:      ksuid.New().String(),
		Stock:   req.Stock,
		Product: product,
		Store:   store,
	}

	storeResponse := storeweb.StoreResponseGlobal{
		Id:      store.Id,
		Name:    store.Name,
		Email:   store.Email,
		NoHp:    store.NoHp,
		Address: store.Address,
	}
	stock, err = s.StockDomain.StockRepository.Create(ctx, tx, stock)
	helper.PanicError(err)

	stockResponse := stockweb.StockResponse{
		Id:      stock.Id,
		Stock:   stock.Stock,
		Product: product,
		Store:   storeResponse,
	}

	return stockResponse, nil
}
func (s *StockServiceImpl) GetAll(ctx context.Context) ([]stockweb.StockResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	stocks, err := s.StockDomain.StockRepository.FindAll(ctx, tx)
	helper.PanicError(err)

	log.Println(stocks)
	var stockResponses []stockweb.StockResponse
	for _, stock := range stocks {

		product, _ := s.StockDomain.ProductRepository.FindById(ctx, tx, stock.Product.Id)
		store, _ := s.StockDomain.StoreRepository.FindById(ctx, tx, stock.Store.Id)

		storeResponse := storeweb.StoreResponseGlobal{
			Id:      store.Id,
			Name:    store.Name,
			Email:   store.Email,
			NoHp:    store.NoHp,
			Address: store.Address,
		}

		stockResponse := stockweb.StockResponse{
			Id:      stock.Id,
			Stock:   stock.Stock,
			Product: product,
			Store:   storeResponse,
		}

		stockResponses = append(stockResponses, stockResponse)
	}

	return stockResponses, nil
}
func (s *StockServiceImpl) GetById(ctx context.Context, stockId string) (stockweb.StockResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	stock, err := s.StockDomain.StockRepository.FindById(ctx, tx, stockId)
	helper.PanicError(err)

	log.Println(stock)
	product, _ := s.StockDomain.ProductRepository.FindById(ctx, tx, stock.Product.Id)
	store, _ := s.StockDomain.StoreRepository.FindById(ctx, tx, stock.Store.Id)

	storeResponse := storeweb.StoreResponseGlobal{
		Id:      store.Id,
		Name:    store.Name,
		Email:   store.Email,
		NoHp:    store.NoHp,
		Address: store.Address,
	}
	stockResponse := stockweb.StockResponse{
		Id:      stock.Id,
		Stock:   stock.Stock,
		Product: product,
		Store:   storeResponse,
	}

	return stockResponse, nil
}
func (s *StockServiceImpl) Edit(ctx context.Context, req stockweb.StockUpdateRequest) (stockweb.StockResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	product, _ := s.StockDomain.ProductRepository.FindById(ctx, tx, req.ProductId)
	store, _ := s.StockDomain.StoreRepository.FindById(ctx, tx, product.Store.Id)

	stock := entity.Stock{
		Id:      req.Id,
		Stock:   req.Stock,
		Product: product,
		Store:   store,
	}

	storeResponse := storeweb.StoreResponseGlobal{
		Id:      store.Id,
		Name:    store.Name,
		Email:   store.Email,
		NoHp:    store.NoHp,
		Address: store.Address,
	}

	stock, err = s.StockDomain.StockRepository.Update(ctx, tx, stock)
	helper.PanicError(err)

	stockResponse := stockweb.StockResponse{
		Id:      stock.Id,
		Stock:   stock.Stock,
		Product: product,
		Store:   storeResponse,
	}

	return stockResponse, nil
}
func (s *StockServiceImpl) Unreg(ctx context.Context, stockId string) (stockweb.StockResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	stock, err := s.StockDomain.StockRepository.FindById(ctx, tx, stockId)
	helper.PanicError(err)

	product, _ := s.StockDomain.ProductRepository.FindById(ctx, tx, stock.Product.Id)
	store, _ := s.StockDomain.StoreRepository.FindById(ctx, tx, stock.Store.Id)

	storeResponse := storeweb.StoreResponseGlobal{
		Id:      store.Id,
		Name:    store.Name,
		Email:   store.Email,
		NoHp:    store.NoHp,
		Address: store.Address,
	}

	stockResponse := stockweb.StockResponse{
		Id:      stock.Id,
		Stock:   stock.Stock,
		Product: product,
		Store:   storeResponse,
	}

	err = s.StockDomain.StockRepository.Delete(ctx, tx, stockId)
	helper.PanicError(err)

	return stockResponse, nil
}
