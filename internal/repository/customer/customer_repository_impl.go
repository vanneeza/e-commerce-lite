package customerrepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repository *CustomerRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, customer entity.Customer) (entity.Customer, error) {

	SQL := "INSERT INTO tbl_customer(id, name, no_hp, email, address, password) VALUES($1, $2, $3, $4, $5, $6)"
	_, err := tx.ExecContext(ctx, SQL, customer.Id, customer.Name, customer.NoHp, customer.Email, customer.Address, customer.Password)
	helper.PanicError(err)

	return customer, nil
}

func (repository *CustomerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customer entity.Customer) (entity.Customer, error) {
	SQL := "UPDATE tbl_customer SET name = $1, no_hp = $2, email = $3, address = $4, password = $5 WHERE id = $6"
	_, err := tx.ExecContext(ctx, SQL, customer.Name, customer.NoHp, customer.Email, customer.Address, customer.Password, customer.Id)
	helper.PanicError(err)
	return customer, nil
}

func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customerId string) error {
	SQL := "UPDATE tbl_customer SET is_deleted = TRUE WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, customerId)
	helper.PanicError(err)

	return nil
}

func (repository *CustomerRepositoryImpl) FindByParams(ctx context.Context, tx *sql.Tx, customerId, noHp string) (entity.Customer, error) {
	SQL := "SELECT id, name, no_hp, email, address, password FROM tbl_customer WHERE is_deleted = false AND id = $1 OR no_hp = $2"
	rows, err := tx.QueryContext(ctx, SQL, customerId, noHp)
	helper.PanicError(err)

	defer rows.Close()

	customer := entity.Customer{}
	if rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Name, &customer.NoHp, &customer.Email, &customer.Address, &customer.Password)
		helper.PanicError(err)
		return customer, nil
	} else {
		return customer, errors.New("customer is not found")
	}
}
func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Customer, error) {
	SQL := "SELECT id, name, no_hp, email, address, password FROM tbl_customer WHERE is_deleted = false"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)

	var categories []entity.Customer
	for rows.Next() {
		customer := entity.Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.NoHp, &customer.Email, &customer.Address, &customer.Password)
		helper.PanicError(err)
		categories = append(categories, customer)
	}

	defer rows.Close()
	return categories, nil
}
