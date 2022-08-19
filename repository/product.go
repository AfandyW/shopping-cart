package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AfandyW/shopping-cart/domain"
)

type repo struct {
}

func NewRepository() domain.Repository {
	return &repo{}
}
func (r *repo) Create(ctx context.Context, tx *sql.Tx, product domain.Product) error {
	sql := "Insert into products(product_code,product_name,quantity) values($1,$2,$3) returning id"

	rows, err := tx.ExecContext(ctx, sql,
		product.ProductCode,
		product.ProductName,
		product.Quantity,
	)

	if err != nil {
		return err
	}

	row, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if row != 1 {
		return fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}
	return nil
}

func (r *repo) Update(ctx context.Context, tx *sql.Tx, product domain.Product) error {
	sql := "update products set quantity = $2 where product_code = $1"
	rows, err := tx.ExecContext(ctx, sql,
		product.ProductCode,
		product.Quantity,
	)

	if err != nil {
		return err
	}

	row, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if row != 1 {
		return fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}
	return nil
}

func (r *repo) Delete(ctx context.Context, tx *sql.Tx, productCode string) error {
	sql := "delete from products where product_code = $1"
	rows, err := tx.ExecContext(ctx, sql,
		productCode,
	)

	if err != nil {
		return err
	}

	row, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if row != 1 {
		return fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}
	return nil
}

func (r *repo) List(ctx context.Context, db *sql.DB) ([]domain.Product, error) {
	sql := "select product_code, product_name, quantity from products"
	rows, err := db.QueryContext(ctx, sql)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(
			&product.ProductCode,
			&product.ProductName,
			&product.Quantity,
		); err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

func (r *repo) Get(ctx context.Context, db *sql.DB, productCode string) (*domain.Product, error) {
	sql := "select product_code, product_name, quantity from products where product_code = $1"
	rows, err := db.QueryContext(ctx, sql, productCode)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var product domain.Product
	if rows.Next() {
		if err := rows.Scan(
			&product.ProductCode,
			&product.ProductName,
			&product.Quantity,
		); err != nil {
			return nil, err
		}
	}
	return &product, nil
}

func (r *repo) GetByName(ctx context.Context, db *sql.DB, productName string) (*domain.Product, error) {
	sql := "select product_code, product_name, quantity from products where product_name = $1"
	rows, err := db.QueryContext(ctx, sql, productName)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var product domain.Product
	if rows.Next() {
		if err := rows.Scan(
			&product.ProductCode,
			&product.ProductName,
			&product.Quantity,
		); err != nil {
			return nil, err
		}
	}
	return &product, nil
}

func (r *repo) ListWithFilter(ctx context.Context, db *sql.DB, filter domain.Filter) ([]domain.Product, error) {
	sql := "select product_code, product_name, quantity from products where product_name=$1"
	rows, err := db.QueryContext(ctx, sql, filter.ProductName)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(
			&product.ProductCode,
			&product.ProductName,
			&product.Quantity,
		); err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}
