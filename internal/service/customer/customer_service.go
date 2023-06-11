package customerservice

import (
	"context"

	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
	customerweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/customer"
)

type CustomerService interface {
	Register(ctx context.Context, req customerweb.CustomerCreateRequest) (customerweb.CustomerResponse, error)
	GetAll(ctx context.Context) ([]customerweb.CustomerResponse, error)
	GetByParams(ctx context.Context, customerId, noHp string) (customerweb.CustomerResponse, error)
	GetPassword(ctx context.Context, customerId, noHp string) (entity.Customer, error)
	Edit(ctx context.Context, req customerweb.CustomerUpdateRequest) (customerweb.CustomerResponse, error)
	Unreg(ctx context.Context, customerId, noHp string) (customerweb.CustomerResponse, error)
}
