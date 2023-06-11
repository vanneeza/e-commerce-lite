package productrepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, product entity.Product) (entity.Product, error) {

	SQL := "INSERT INTO tbl_product(id, name, price, store_id) VALUES($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, SQL, product.Id, product.Name, product.Price, product.Store.Id)
	helper.PanicError(err)

	return product, nil
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product entity.Product) (entity.Product, error) {
	SQL := "UPDATE tbl_product SET name = $1, price = $2, store_id = $3 WHERE id = $4"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Store.Id, product.Id)
	helper.PanicError(err)
	return product, nil
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productId string) error {
	SQL := "UPDATE tbl_product SET is_deleted = TRUE WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, productId)
	helper.PanicError(err)

	return nil
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId string) (entity.Product, error) {
	SQL := "SELECT id, name, price, store_id FROM tbl_product WHERE id = $1 AND is_deleted = FALSE"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicError(err)
	defer rows.Close()

	product := entity.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Store.Id)
		helper.PanicError(err)
		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}
func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Product, error) {
	SQL := "SELECT id, name, price, store_id FROM tbl_product WHERE is_deleted = FALSE"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)

	var products []entity.Product
	for rows.Next() {
		product := entity.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Store.Id)
		helper.PanicError(err)
		products = append(products, product)
	}

	defer rows.Close()
	return products, nil
}
