package productservice

import (
	"context"

	productweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/product"
)

type ProductService interface {
	Register(ctx context.Context, req productweb.ProductCreateRequest) (productweb.ProductResponse, error)
	GetAll(ctx context.Context) ([]productweb.ProductResponse, error)
	GetById(ctx context.Context, productId string) (productweb.ProductResponse, error)
	Edit(ctx context.Context, req productweb.ProductUpdateRequest) (productweb.ProductResponse, error)
	Unreg(ctx context.Context, productId string) (productweb.ProductResponse, error)
}
