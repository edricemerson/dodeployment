package handler

import (
	"net/http"

	"p3-graded-challenge-1-edricemerson/payment-app/service/payment"
	"p3-graded-challenge-1-edricemerson/payment-app/util"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	service payment.Service
}

func NewPaymentHandler(s payment.Service) *PaymentHandler {
	return &PaymentHandler{
		service: s,
	}
}

func (h *PaymentHandler) Create(c echo.Context) error {

	var p payment.Payment

	if err := c.Bind(&p); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	result, err := h.service.Create(p)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}
