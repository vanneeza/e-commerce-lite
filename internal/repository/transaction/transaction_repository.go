package transactionrepository

import (
	"context"
	"database/sql"

	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
)

type TrasactionRepository interface {
	CreateOrder(ctx context.Context, tx *sql.Tx, order entity.Order) (entity.Order, error)
	CreateOrderDetail(ctx context.Context, tx *sql.Tx, orderDetail entity.OrderDetail) (entity.OrderDetail, error)
	CreateCallback(ctx context.Context, tx *sql.Tx, callBack entity.Callback) (entity.Callback, error)
	UpdateStatusDetail(ctx context.Context, tx *sql.Tx, callback entity.Callback) error
}
