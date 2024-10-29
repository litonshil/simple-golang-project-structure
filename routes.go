package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"simple-project/controller"
	"simple-project/repository"
	"simple-project/service"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	stockRepo := repository.NewStockRepository(db)
	stockService := service.NewStockService(stockRepo)
	stockController := controller.NewStockController(stockService)

	e.POST("/stocks", stockController.CreateStock)

	e.GET("/stocks", stockController.GetStocks)
}
