package main

import (
	"p3-graded-challenge-1-edricemerson/shopping-app/handler"
	"p3-graded-challenge-1-edricemerson/shopping-app/repository"
	"p3-graded-challenge-1-edricemerson/shopping-app/service/product"
	"p3-graded-challenge-1-edricemerson/shopping-app/service/transaction"
	"p3-graded-challenge-1-edricemerson/shopping-app/util"

	"github.com/labstack/echo/v4"
)

func main() {

	util.ConnectDB()

	util.StartScheduler()

	transactionRepo := repository.NewTransactionRepository(util.DB)
	transactionService := transaction.NewService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	productRepo := repository.NewProductRepository(util.DB)
	productService := product.NewService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	e := echo.New()

	e.POST("/transactions", transactionHandler.Create)
	e.GET("/transactions", transactionHandler.GetAll)
	e.GET("/transactions/:id", transactionHandler.GetByID)
	e.PUT("/transactions/:id", transactionHandler.Update)
	e.DELETE("/transactions/:id", transactionHandler.Delete)

	e.POST("/products", productHandler.Create)
	e.GET("/products", productHandler.GetAll)
	e.GET("/products/:id", productHandler.GetByID)
	e.PUT("/products/:id", productHandler.Update)
	e.DELETE("/products/:id", productHandler.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
