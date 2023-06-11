package customerrepository

import (
	"context"
	"database/sql"

	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
)

type CustomerRepository interface {
	Create(ctx context.Context, tx *sql.Tx, customer entity.Customer) (entity.Customer, error)
	Update(ctx context.Context, tx *sql.Tx, customer entity.Customer) (entity.Customer, error)
	Delete(ctx context.Context, tx *sql.Tx, customerId string) error
	FindByParams(ctx context.Context, tx *sql.Tx, customerId, noHp string) (entity.Customer, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Customer, error)
}
