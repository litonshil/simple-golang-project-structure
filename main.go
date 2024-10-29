package main

import (
	"github.com/labstack/echo/v4"
	"simple-project/connection"
)

func main() {
	e := echo.New()

	connection.ConnectDB()
	dbClient := connection.GetDB()

	InitRoutes(e, dbClient)

	//stockRepo := repository.NewStockRepository()
	//stockService := service.NewStockService(stockRepo)
	//stockController := controller.NewStockController(stockService)
	//
	//e.POST("/stocks", stockController.CreateStock)
	//
	//e.GET("/stocks", stockController.GetStocks)

	e.Logger.Fatal(e.Start(":8080"))
}
