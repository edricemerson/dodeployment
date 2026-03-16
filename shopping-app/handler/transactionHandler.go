package handler

import (
	"net/http"

	"p3-graded-challenge-1-edricemerson/shopping-app/service/transaction"
	"p3-graded-challenge-1-edricemerson/shopping-app/util"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(s transaction.Service) *TransactionHandler {
	return &TransactionHandler{
		service: s,
	}
}

func (h *TransactionHandler) Create(c echo.Context) error {

	var t transaction.Transaction

	if err := c.Bind(&t); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	result, err := h.service.Create(t)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *TransactionHandler) GetAll(c echo.Context) error {

	data, err := h.service.GetAll()
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h *TransactionHandler) GetByID(c echo.Context) error {

	id := c.Param("id")

	data, err := h.service.GetByID(id)
	if err != nil {
		return util.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h *TransactionHandler) Update(c echo.Context) error {

	id := c.Param("id")

	var t transaction.Transaction

	if err := c.Bind(&t); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	result, err := h.service.Update(id, t)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *TransactionHandler) Delete(c echo.Context) error {

	id := c.Param("id")

	err := h.service.Delete(id)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "transaction deleted",
	})
}
