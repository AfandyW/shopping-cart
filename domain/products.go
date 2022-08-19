package domain

import (
	"context"
	"database/sql"
)

type Product struct {
	ID          int
	ProductCode string
	ProductName string
	Quantity    int
}

func (p Product) IsAvailable() bool {
	return p.ProductCode != ""
}

type ProductRequest struct {
	ProductCode string `json:"product_code"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}

type ProductResponse struct {
	ProductCode string `json:"product_code"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}

type Filter struct {
	ProductName string
}

type Repository interface {
	Create(ctx context.Context, tx *sql.Tx, product Product) error
	Update(ctx context.Context, tx *sql.Tx, product Product) error
	Delete(ctx context.Context, tx *sql.Tx, productCode string) error
	List(ctx context.Context, db *sql.DB) ([]Product, error)
	ListWithFilter(ctx context.Context, db *sql.DB, filter Filter) ([]Product, error)
	Get(ctx context.Context, db *sql.DB, productCode string) (*Product, error)
	GetByName(ctx context.Context, db *sql.DB, productName string) (*Product, error)
}

type Service interface {
	Create(ctx context.Context, payload ProductRequest) error
	Update(ctx context.Context, payload ProductRequest) error
	Delete(ctx context.Context, productCode string) error
	List(ctx context.Context, filter Filter) ([]ProductResponse, error)
	Get(ctx context.Context, productCode string) (*ProductResponse, error)
}
