package transactionrepository

import (
	"context"
	"database/sql"

	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TrasactionRepository {
	return &TransactionRepositoryImpl{}
}

// CreateOrder implements TrasactionRepository.
func (*TransactionRepositoryImpl) CreateOrder(ctx context.Context, tx *sql.Tx, order entity.Order) (entity.Order, error) {
	SQL := "INSERT INTO tx_order (id, qty, product_id, customer_id, store_id, tx_detail_id) VALUES($1, $2, $3, $4, $5, $6)"
	_, err := tx.ExecContext(ctx, SQL, order.Id, order.Qty, order.ProductId.Id, order.CustomerId.Id, order.StoreId.Id, order.TxDetailId.Id)
	helper.PanicError(err)

	return order, nil
}

// CreateOrderDetail implements TrasactionRepository.
func (*TransactionRepositoryImpl) CreateOrderDetail(ctx context.Context, tx *sql.Tx, orderDetail entity.OrderDetail) (entity.OrderDetail, error) {
	SQL := "INSERT INTO tx_detail (id, total_price, status, buy_date) VALUES($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, SQL, orderDetail.Id, orderDetail.TotalPrice, orderDetail.Status, orderDetail.BuyDate)
	helper.PanicError(err)

	return orderDetail, nil
}

func (*TransactionRepositoryImpl) CreateCallback(ctx context.Context, tx *sql.Tx, callBack entity.Callback) (entity.Callback, error) {

	SQL := `INSERT INTO tx_callback (id, detail_id, user_id, is_high, payment_method, status, merchant_name, amount, paid_amount, bank_code,
		paid_date, payer_email, description, created, updated, currency, payment_channel, payment_destination)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12 ,$13, $14, $15, $16, $17, $18)`
	_, err := tx.ExecContext(ctx, SQL,
		callBack.Id,
		callBack.DetailId,
		callBack.UserId,
		callBack.IsHigh,
		callBack.PaymentMethod,
		callBack.Status,
		callBack.MerchantName,
		callBack.Amount,
		callBack.PaidAmount,
		callBack.BankCode,
		callBack.PaidDate,
		callBack.PayerEmail,
		callBack.Description,
		callBack.Created,
		callBack.Updated,
		callBack.Currency,
		callBack.PaymentChannel,
		callBack.PaymentDestination)
	helper.PanicError(err)

	return callBack, nil
}

func (*TransactionRepositoryImpl) UpdateStatusDetail(ctx context.Context, tx *sql.Tx, callback entity.Callback) error {
	SQL := "UPDATE tx_detail SET status = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, SQL, callback.Status, callback.DetailId)
	helper.PanicError(err)

	return nil
}
