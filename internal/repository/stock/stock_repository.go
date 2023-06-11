package stockrepository

import (
	"context"
	"database/sql"

	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
)

type StockRepository interface {
	Create(ctx context.Context, tx *sql.Tx, stock entity.Stock) (entity.Stock, error)
	Update(ctx context.Context, tx *sql.Tx, stock entity.Stock) (entity.Stock, error)
	Delete(ctx context.Context, tx *sql.Tx, stockId string) error
	FindById(ctx context.Context, tx *sql.Tx, stockId string) (entity.Stock, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Stock, error)
}
