package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AfandyW/shopping-cart/domain"
)

type service struct {
	Repo domain.Repository
	DB   *sql.DB
}

func NewService(repo domain.Repository, db *sql.DB) domain.Service {
	return &service{
		Repo: repo,
		DB:   db,
	}
}

func (s *service) Create(ctx context.Context, payload domain.ProductRequest) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return fmt.Errorf("service.Create: Failed to begin db")
	}
	defer tx.Rollback()

	products, err := s.Repo.List(ctx, s.DB)
	if err != nil {
		return fmt.Errorf("service.Create.List: return err %s", err.Error())
	}

	totalProduct := len(products)

	product, err := s.Repo.GetByName(ctx, s.DB, payload.ProductName)
	if err != nil {
		return fmt.Errorf("return err %s", err.Error())
	}

	if product.IsAvailable() {
		return fmt.Errorf("product name already exists")
	}

	err = s.Repo.Create(ctx, tx, domain.Product{
		ProductCode: fmt.Sprintf("p-%d", totalProduct+1),
		ProductName: payload.ProductName,
		Quantity:    payload.Quantity,
	})

	if err != nil {
		return fmt.Errorf("return err %s", err.Error())
	}

	tx.Commit()

	return nil
}

func (s *service) Update(ctx context.Context, payload domain.ProductRequest) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin db")
	}
	defer tx.Rollback()

	product, err := s.Repo.Get(ctx, s.DB, payload.ProductCode)
	if err != nil {
		return fmt.Errorf("return err %s", err.Error())
	}

	if !product.IsAvailable() {
		return fmt.Errorf("product not found")
	}

	err = s.Repo.Update(ctx, tx, domain.Product{
		ProductCode: payload.ProductCode,
		Quantity:    payload.Quantity,
	})

	if err != nil {
		return fmt.Errorf("return err %s", err.Error())
	}

	tx.Commit()

	return nil
}

func (s *service) Delete(ctx context.Context, productCode string) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin db")
	}
	defer tx.Rollback()

	product, err := s.Repo.Get(ctx, s.DB, productCode)
	if err != nil {
		return fmt.Errorf("return err %s", err.Error())
	}

	if !product.IsAvailable() {
		return fmt.Errorf("product not found")
	}

	err = s.Repo.Delete(ctx, tx, product.ProductCode)

	if err != nil {
		return fmt.Errorf("return err %s", err.Error())
	}

	tx.Commit()

	return nil
}

func (s *service) Get(ctx context.Context, productCode string) (*domain.ProductResponse, error) {
	product, err := s.Repo.Get(ctx, s.DB, productCode)
	if err != nil {
		return nil, fmt.Errorf("return err %s", err.Error())
	}

	if !product.IsAvailable() {
		return nil, fmt.Errorf("product not found")
	}

	return &domain.ProductResponse{
		ProductCode: product.ProductCode,
		ProductName: product.ProductName,
		Quantity:    product.Quantity,
	}, nil
}

func (s *service) List(ctx context.Context) ([]domain.ProductResponse, error) {
	resp, err := s.Repo.List(ctx, s.DB)
	if err != nil {
		return nil, fmt.Errorf("return err %s", err.Error())
	}

	var products []domain.ProductResponse
	for _, v := range resp {
		product := domain.ProductResponse{
			ProductCode: v.ProductCode,
			ProductName: v.ProductName,
			Quantity:    v.Quantity,
		}

		products = append(products, product)
	}

	return products, nil
}
