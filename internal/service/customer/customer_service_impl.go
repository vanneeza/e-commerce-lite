package customerservice

import (
	"context"
	"database/sql"

	"github.com/segmentio/ksuid"
	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
	customerweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/customer"
	customerrepository "github.com/vanneeza/e-commerce-lite/internal/repository/customer"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
	"golang.org/x/crypto/bcrypt"
)

type CustomerServiceImpl struct {
	Db                 *sql.DB
	CustomerRepository customerrepository.CustomerRepository
}

func NewCustomerService(db *sql.DB, customerRepository customerrepository.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{
		Db:                 db,
		CustomerRepository: customerRepository,
	}

}

func (s *CustomerServiceImpl) Register(ctx context.Context, req customerweb.CustomerCreateRequest) (customerweb.CustomerResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	customer := entity.Customer{
		Id:       ksuid.New().String(),
		Name:     req.Name,
		Email:    req.Email,
		NoHp:     req.NoHp,
		Address:  req.Address,
		Password: string(encryptedPassword),
	}

	customer, err = s.CustomerRepository.Create(ctx, tx, customer)
	helper.PanicError(err)

	customerResponse := customerweb.CustomerResponse{
		Id:      customer.Id,
		Name:    customer.Name,
		Email:   customer.Email,
		NoHp:    customer.NoHp,
		Address: customer.Address,
	}

	return customerResponse, nil
}
func (s *CustomerServiceImpl) GetAll(ctx context.Context) ([]customerweb.CustomerResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	customers, err := s.CustomerRepository.FindAll(ctx, tx)
	helper.PanicError(err)

	var customerResponses []customerweb.CustomerResponse
	for _, customer := range customers {
		customerResponse := customerweb.CustomerResponse{
			Id:      customer.Id,
			Name:    customer.Name,
			NoHp:    customer.NoHp,
			Email:   customer.Email,
			Address: customer.Address,
		}
		customerResponses = append(customerResponses, customerResponse)
	}

	return customerResponses, nil
}
func (s *CustomerServiceImpl) GetByParams(ctx context.Context, customerId, noHp string) (customerweb.CustomerResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := s.CustomerRepository.FindByParams(ctx, tx, customerId, noHp)
	helper.PanicError(err)

	customerResponse := customerweb.CustomerResponse{
		Id:      customer.Id,
		Name:    customer.Name,
		Email:   customer.Email,
		NoHp:    customer.NoHp,
		Address: customer.Address,
	}

	return customerResponse, nil
}
func (s *CustomerServiceImpl) Edit(ctx context.Context, req customerweb.CustomerUpdateRequest) (customerweb.CustomerResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	customer := entity.Customer{
		Id:       req.Id,
		Name:     req.Name,
		Email:    req.Email,
		NoHp:     req.NoHp,
		Address:  req.Address,
		Password: string(encryptedPassword),
	}

	customer, err = s.CustomerRepository.Update(ctx, tx, customer)
	helper.PanicError(err)

	customerResponse := customerweb.CustomerResponse{
		Id:      customer.Id,
		Name:    customer.Name,
		Email:   customer.Email,
		NoHp:    customer.NoHp,
		Address: customer.Address,
	}

	return customerResponse, nil
}
func (s *CustomerServiceImpl) Unreg(ctx context.Context, customerId, noHp string) (customerweb.CustomerResponse, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := s.CustomerRepository.FindByParams(ctx, tx, customerId, noHp)
	helper.PanicError(err)

	customerResponse := customerweb.CustomerResponse{
		Id:      customer.Id,
		Name:    customer.Name,
		Email:   customer.Email,
		NoHp:    customer.NoHp,
		Address: customer.Address,
	}

	err = s.CustomerRepository.Delete(ctx, tx, customerId)
	helper.PanicError(err)

	return customerResponse, nil
}

func (s *CustomerServiceImpl) GetPassword(ctx context.Context, customerId, noHp string) (entity.Customer, error) {
	tx, err := s.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := s.CustomerRepository.FindByParams(ctx, tx, customerId, noHp)
	helper.PanicError(err)

	return customer, nil
}
