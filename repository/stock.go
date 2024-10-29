package repository

import (
	"gorm.io/gorm"
	"simple-project/domain"
)

//var StockResp []domain.Stock

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) domain.StockRepository {
	return &stockRepository{
		db: db,
	}
}

func (r *stockRepository) CreateStock(req domain.Stock) (domain.Stock, error) {
	return domain.Stock{}, nil
}
