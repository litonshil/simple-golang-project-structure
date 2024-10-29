package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"simple-project/domain"
)

type StockController struct {
	stockService domain.StockService
}

func NewStockController(stockService domain.StockService) *StockController {
	return &StockController{stockService}
}

func (c *StockController) CreateStock(ctx echo.Context) error {

	stockReq := domain.Stock{}

	if err := ctx.Bind(&stockReq); err != nil {
		return err
	}

	_, err := c.stockService.CreateStock(stockReq)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, "successfully added")
}

func (c *StockController) GetStocks(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, []domain.Stock{})
}
