package productrepository

import (
	"context"
	"database/sql"

	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
)

type ProductRepository interface {
	Create(ctx context.Context, tx *sql.Tx, product entity.Product) (entity.Product, error)
	Update(ctx context.Context, tx *sql.Tx, product entity.Product) (entity.Product, error)
	Delete(ctx context.Context, tx *sql.Tx, productId string) error
	FindById(ctx context.Context, tx *sql.Tx, productId string) (entity.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Product, error)
}
