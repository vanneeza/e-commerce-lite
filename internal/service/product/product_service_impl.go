package productservice

import (
	"context"
	"database/sql"
	"log"

	"github.com/segmentio/ksuid"
	"github.com/vanneeza/e-commerce-lite/internal/domain"
	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
	productweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/product"
	storeweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/store"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type ProductServiceImpl struct {
	Db            *sql.DB
	ProductDomain domain.ProductDomain
}

func NewProductService(db *sql.DB, productDomain domain.ProductDomain) ProductService {
	return &ProductServiceImpl{
		Db:            db,
		ProductDomain: productDomain,
	}

}

func (s *ProductServiceImpl) Register(ctx context.Context, req productweb.ProductCreateRequest) (productweb.ProductResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	store, _ := s.ProductDomain.StoreRepository.FindById(ctx, tx, req.Store)
	storeResponse := storeweb.StoreResponseGlobal{
		Id:      store.Id,
		Name:    store.Name,
		Email:   store.Email,
		NoHp:    store.NoHp,
		Address: store.Address,
	}

	product := entity.Product{
		Id:    ksuid.New().String(),
		Name:  req.Name,
		Price: req.Price,
		Store: store,
	}

	product, err = s.ProductDomain.ProductRepository.Create(ctx, tx, product)
	helper.PanicError(err)

	productResponse := productweb.ProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
		Store: storeResponse,
	}

	return productResponse, nil
}
func (s *ProductServiceImpl) GetAll(ctx context.Context) ([]productweb.ProductResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	products, err := s.ProductDomain.ProductRepository.FindAll(ctx, tx)
	helper.PanicError(err)

	log.Println(products)
	var productResponses []productweb.ProductResponse
	for _, product := range products {

		store, _ := s.ProductDomain.StoreRepository.FindById(ctx, tx, product.Store.Id)

		storeResponse := storeweb.StoreResponseGlobal{
			Id:      store.Id,
			Name:    store.Name,
			Email:   store.Email,
			NoHp:    store.NoHp,
			Address: store.Address,
		}

		productResponse := productweb.ProductResponse{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
			Store: storeResponse,
		}

		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}
func (s *ProductServiceImpl) GetById(ctx context.Context, productId string) (productweb.ProductResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	product, err := s.ProductDomain.ProductRepository.FindById(ctx, tx, productId)
	helper.PanicError(err)

	log.Println(product)
	store, _ := s.ProductDomain.StoreRepository.FindById(ctx, tx, product.Store.Id)

	storeResponse := storeweb.StoreResponseGlobal{
		Id:      store.Id,
		Name:    store.Name,
		Email:   store.Email,
		NoHp:    store.NoHp,
		Address: store.Address,
	}

	productResponse := productweb.ProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
		Store: storeResponse,
	}
	return productResponse, nil
}
func (s *ProductServiceImpl) Edit(ctx context.Context, req productweb.ProductUpdateRequest) (productweb.ProductResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	store, _ := s.ProductDomain.StoreRepository.FindById(ctx, tx, req.StoreId)

	product := entity.Product{
		Id:    req.Id,
		Name:  req.Name,
		Price: req.Price,
		Store: store,
	}

	storeResponse := storeweb.StoreResponseGlobal{
		Id:      store.Id,
		Name:    store.Name,
		Email:   store.Email,
		NoHp:    store.NoHp,
		Address: store.Address,
	}

	product, err = s.ProductDomain.ProductRepository.Update(ctx, tx, product)
	helper.PanicError(err)

	productResponse := productweb.ProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
		Store: storeResponse,
	}

	return productResponse, nil
}
func (s *ProductServiceImpl) Unreg(ctx context.Context, productId string) (productweb.ProductResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	product, err := s.ProductDomain.ProductRepository.FindById(ctx, tx, productId)
	helper.PanicError(err)
	store, _ := s.ProductDomain.StoreRepository.FindById(ctx, tx, product.Store.Id)

	storeResponse := storeweb.StoreResponseGlobal{
		Id:      store.Id,
		Name:    store.Name,
		Email:   store.Email,
		NoHp:    store.NoHp,
		Address: store.Address,
	}

	productResponse := productweb.ProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
		Store: storeResponse,
	}

	err = s.ProductDomain.ProductRepository.Delete(ctx, tx, productId)
	helper.PanicError(err)

	return productResponse, nil
}
