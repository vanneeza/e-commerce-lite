package storerepository

import (
	"context"
	"database/sql"

	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
)

type StoreRepository interface {
	Create(ctx context.Context, tx *sql.Tx, store entity.Store) (entity.Store, error)
	Update(ctx context.Context, tx *sql.Tx, store entity.Store) (entity.Store, error)
	Delete(ctx context.Context, tx *sql.Tx, storeId string) error
	FindById(ctx context.Context, tx *sql.Tx, storeId string) (entity.Store, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Store, error)
}
