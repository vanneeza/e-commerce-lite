package stockrepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type StockRepositoryImpl struct {
}

func NewStockRepository() StockRepository {
	return &StockRepositoryImpl{}
}

func (repository *StockRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, stock entity.Stock) (entity.Stock, error) {

	SQL := "INSERT INTO tbl_stock(id, stock, product_id, store_id) VALUES($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, SQL, stock.Id, stock.Stock, stock.Product.Id, stock.Store.Id)
	helper.PanicError(err)

	return stock, nil
}

func (repository *StockRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, stock entity.Stock) (entity.Stock, error) {
	SQL := "UPDATE tbl_stock SET stock = $1, product_id = $2, store_id = $3 WHERE id = $4"
	_, err := tx.ExecContext(ctx, SQL, stock.Stock, stock.Product.Id, stock.Store.Id, stock.Id)
	helper.PanicError(err)
	return stock, nil
}

func (repository *StockRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, stockId string) error {
	SQL := "UPDATE tbl_stock SET stock = 0 WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, stockId)
	helper.PanicError(err)

	return nil
}

func (repository *StockRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, stockId string) (entity.Stock, error) {
	SQL := "SELECT id, stock, product_id, store_id FROM tbl_stock WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, stockId)
	helper.PanicError(err)
	defer rows.Close()

	stock := entity.Stock{}
	if rows.Next() {
		err := rows.Scan(&stock.Id, &stock.Stock, &stock.Product.Id, &stock.Store.Id)
		helper.PanicError(err)
		return stock, nil
	} else {
		return stock, errors.New("stock is not found")
	}
}
func (repository *StockRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Stock, error) {
	SQL := "SELECT id, stock, product_id, store_id FROM tbl_stock"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)

	var stocks []entity.Stock
	for rows.Next() {
		stock := entity.Stock{}
		err := rows.Scan(&stock.Id, &stock.Stock, &stock.Product.Id, &stock.Store.Id)
		helper.PanicError(err)
		stocks = append(stocks, stock)
	}

	defer rows.Close()
	return stocks, nil
}
