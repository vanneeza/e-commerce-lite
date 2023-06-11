package storeservice

import (
	"context"
	"database/sql"

	"github.com/segmentio/ksuid"
	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
	storeweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/store"
	storerepository "github.com/vanneeza/e-commerce-lite/internal/repository/store"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type StoreServiceImpl struct {
	Db              *sql.DB
	StoreRepository storerepository.StoreRepository
}

func NewStoreService(db *sql.DB, storeRepository storerepository.StoreRepository) StoreService {
	return &StoreServiceImpl{
		Db:              db,
		StoreRepository: storeRepository,
	}

}

func (s *StoreServiceImpl) Register(ctx context.Context, req storeweb.StoreCreateRequest) (storeweb.StoreResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	store := entity.Store{
		Id:      ksuid.New().String(),
		Name:    req.Name,
		Email:   req.Email,
		NoHp:    req.NoHp,
		Address: req.Address,
		Balance: 0,
	}

	store, err = s.StoreRepository.Create(ctx, tx, store)
	helper.PanicError(err)

	storeResponse := storeweb.StoreResponse{
		Id:      store.Id,
		Name:    store.Name,
		Email:   store.Email,
		NoHp:    store.NoHp,
		Address: store.Address,
		Balance: store.Balance,
	}

	return storeResponse, nil
}
func (s *StoreServiceImpl) GetAll(ctx context.Context) ([]storeweb.StoreResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	stores, err := s.StoreRepository.FindAll(ctx, tx)
	helper.PanicError(err)

	var storeResponses []storeweb.StoreResponse
	for _, store := range stores {
		storeResponses = append(storeResponses, storeweb.StoreResponse(store))
	}

	return storeResponses, nil
}
func (s *StoreServiceImpl) GetById(ctx context.Context, storeId string) (storeweb.StoreResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	store, err := s.StoreRepository.FindById(ctx, tx, storeId)
	helper.PanicError(err)

	storeResponse := storeweb.StoreResponse{
		Id:      store.Id,
		Name:    store.Name,
		Email:   store.Email,
		NoHp:    store.NoHp,
		Address: store.Address,
		Balance: store.Balance,
	}

	return storeResponse, nil
}
func (s *StoreServiceImpl) Edit(ctx context.Context, req storeweb.StoreUpdateRequest) (storeweb.StoreResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	store := entity.Store{
		Id:      req.Id,
		Name:    req.Name,
		Email:   req.Email,
		NoHp:    req.NoHp,
		Address: req.Address,
		Balance: req.Balance,
	}

	store, err = s.StoreRepository.Update(ctx, tx, store)
	helper.PanicError(err)

	storeResponse := storeweb.StoreResponse{
		Id:      store.Id,
		Name:    store.Name,
		Email:   store.Email,
		NoHp:    store.NoHp,
		Address: store.Address,
		Balance: store.Balance,
	}

	return storeResponse, nil
}
func (s *StoreServiceImpl) Unreg(ctx context.Context, storeId string) (storeweb.StoreResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	store, err := s.StoreRepository.FindById(ctx, tx, storeId)
	helper.PanicError(err)

	storeResponse := storeweb.StoreResponse{
		Id:      store.Id,
		Name:    store.Name,
		Email:   store.Email,
		NoHp:    store.NoHp,
		Address: store.Address,
		Balance: store.Balance,
	}

	err = s.StoreRepository.Delete(ctx, tx, storeId)
	helper.PanicError(err)

	return storeResponse, nil
}
