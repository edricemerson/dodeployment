package handler

import (
	"net/http"

	"p3-graded-challenge-1-edricemerson/shopping-app/service/product"
	"p3-graded-challenge-1-edricemerson/shopping-app/util"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service product.Service
}

func NewProductHandler(s product.Service) *ProductHandler {
	return &ProductHandler{
		service: s,
	}
}

func (h *ProductHandler) Create(c echo.Context) error {

	var p product.Product

	if err := c.Bind(&p); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	result, err := h.service.Create(p)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *ProductHandler) GetAll(c echo.Context) error {

	data, err := h.service.GetAll()
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h *ProductHandler) GetByID(c echo.Context) error {

	id := c.Param("id")

	data, err := h.service.GetByID(id)
	if err != nil {
		return util.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h *ProductHandler) Update(c echo.Context) error {

	id := c.Param("id")

	var p product.Product

	if err := c.Bind(&p); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	result, err := h.service.Update(id, p)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *ProductHandler) Delete(c echo.Context) error {

	id := c.Param("id")

	err := h.service.Delete(id)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "product deleted",
	})
}
