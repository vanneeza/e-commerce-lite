package stockservice

import (
	"context"

	stockweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/stock"
)

type StockService interface {
	Register(ctx context.Context, req stockweb.StockCreateRequest) (stockweb.StockResponse, error)
	GetAll(ctx context.Context) ([]stockweb.StockResponse, error)
	GetById(ctx context.Context, stockId string) (stockweb.StockResponse, error)
	Edit(ctx context.Context, req stockweb.StockUpdateRequest) (stockweb.StockResponse, error)
	Unreg(ctx context.Context, stockId string) (stockweb.StockResponse, error)
}
