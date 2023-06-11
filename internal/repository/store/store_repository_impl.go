package storerepository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type StoreRepositoryImpl struct {
}

func NewStoreRepository() StoreRepository {
	return &StoreRepositoryImpl{}
}

func (repository *StoreRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, store entity.Store) (entity.Store, error) {

	SQL := "INSERT INTO tbl_store(id, name, email, no_hp, address, balance) VALUES($1, $2, $3, $4, $5, $6)"
	_, err := tx.ExecContext(ctx, SQL, store.Id, store.Name, store.Email, store.NoHp, store.Address, store.Balance)
	helper.PanicError(err)

	return store, nil
}

func (repository *StoreRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, store entity.Store) (entity.Store, error) {
	SQL := "UPDATE tbl_store SET name = $1, email = $2, no_hp = $3, address = $4, balance = $5 WHERE id = $6 AND is_deleted = false"
	_, err := tx.ExecContext(ctx, SQL, store.Name, store.Email, store.NoHp, store.Address, store.Balance, store.Id)
	helper.PanicError(err)
	return store, nil
}

func (repository *StoreRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, storeId string) error {
	SQL := "UPDATE tbl_store SET is_deleted = TRUE WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, storeId)
	helper.PanicError(err)

	return nil
}

func (repository *StoreRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, storeId string) (entity.Store, error) {
	log.Println(storeId, "repo store atas")
	SQL := "SELECT id, name, email, no_hp, address, balance FROM tbl_store WHERE id = $1 AND is_deleted = false"
	rows, err := tx.QueryContext(ctx, SQL, storeId)
	helper.PanicError(err)

	defer rows.Close()

	store := entity.Store{}
	if rows.Next() {
		err := rows.Scan(&store.Id, &store.Name, &store.Email, &store.NoHp, &store.Address, &store.Balance)
		helper.PanicError(err)
		log.Println(store, "repo store tengah")

		return store, nil
	} else {
		return store, errors.New("store is not found")
	}
}
func (repository *StoreRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Store, error) {
	SQL := "SELECT id, name, email, no_hp, address, balance FROM tbl_store WHERE is_deleted = false"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)

	var categories []entity.Store
	for rows.Next() {
		store := entity.Store{}
		err := rows.Scan(&store.Id, &store.Name, &store.Email, &store.NoHp, &store.Address, &store.Balance)
		helper.PanicError(err)
		categories = append(categories, store)
	}

	defer rows.Close()
	return categories, nil
}
