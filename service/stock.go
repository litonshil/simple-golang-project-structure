package service

import (
	"simple-project/domain"
)

type stockService struct {
	stockRepo domain.StockRepository
}

func NewStockService(stockRepo domain.StockRepository) domain.StockService {
	return &stockService{stockRepo}
}

func (s *stockService) CreateStock(req domain.Stock) (domain.Stock, error) {
	return s.stockRepo.CreateStock(req)
}
