package main

import (
	"p3-graded-challenge-1-edricemerson/payment-app/handler"
	"p3-graded-challenge-1-edricemerson/payment-app/repository"
	"p3-graded-challenge-1-edricemerson/payment-app/service/payment"
	"p3-graded-challenge-1-edricemerson/payment-app/util"

	"github.com/labstack/echo/v4"
)

func main() {

	util.ConnectDB()

	util.StartScheduler()

	paymentRepo := repository.NewMongoRepository(util.DB)

	paymentService := payment.NewService(paymentRepo)

	paymentHandler := handler.NewPaymentHandler(paymentService)

	e := echo.New()

	e.POST("/payments", paymentHandler.Create)

	e.Logger.Fatal(e.Start(":8081"))
}
