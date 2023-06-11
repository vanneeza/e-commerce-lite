package transactionservice

import (
	"context"

	transactionweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/transaction"
)

type TransactionService interface {
	MakeOrder(ctx context.Context, req transactionweb.CreateOrderRequest) (transactionweb.ResponseOrder, error)
	SaveCallBack(ctx context.Context, req transactionweb.CallbackResponse) error
}
