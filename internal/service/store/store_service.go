package storeservice

import (
	"context"

	storeweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/store"
)

type StoreService interface {
	Register(ctx context.Context, req storeweb.StoreCreateRequest) (storeweb.StoreResponse, error)
	GetAll(ctx context.Context) ([]storeweb.StoreResponse, error)
	GetById(ctx context.Context, storeId string) (storeweb.StoreResponse, error)
	Edit(ctx context.Context, req storeweb.StoreUpdateRequest) (storeweb.StoreResponse, error)
	Unreg(ctx context.Context, storeId string) (storeweb.StoreResponse, error)
}
