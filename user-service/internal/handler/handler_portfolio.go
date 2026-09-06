package handler

import (
	"net/http"
	"strconv"

	"user-service/internal/dto/request"
	"user-service/internal/helper"
	portfolioUsecase "user-service/internal/usecase/portfolio"

	"github.com/labstack/echo/v5"
)

type PortfolioHandler struct {
	portfolioUsecase portfolioUsecase.PortfolioUsecase
}

func NewPortfolioHandler(
	portfolioUsecase portfolioUsecase.PortfolioUsecase,
) *PortfolioHandler {
	return &PortfolioHandler{
		portfolioUsecase: portfolioUsecase,
	}
}

func (h *PortfolioHandler) Create(c *echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	var req request.CreatePortfolioRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(
			c,
			"invalid request body",
		)
	}

	result, err := h.portfolioUsecase.Create(
		c.Request().Context(),
		userID,
		req,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusCreated,
		"portfolio berhasil dibuat",
		result,
	)
}

func (h *PortfolioHandler) GetByUserID(c *echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	result, err := h.portfolioUsecase.GetByUserID(
		c.Request().Context(),
		userID,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"portfolio berhasil ditemukan",
		result,
	)
}

func (h *PortfolioHandler) GetByID(c *echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	portfolioID, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		return helper.BadRequest(
			c,
			"invalid portfolio id",
		)
	}

	result, err := h.portfolioUsecase.GetByID(
		c.Request().Context(),
		userID,
		uint(portfolioID),
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"portfolio berhasil ditemukan",
		result,
	)
}

func (h *PortfolioHandler) Update(c *echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	portfolioID, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		return helper.BadRequest(
			c,
			"invalid portfolio id",
		)
	}

	var req request.UpdatePortfolioRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(
			c,
			"invalid request body",
		)
	}

	result, err := h.portfolioUsecase.Update(
		c.Request().Context(),
		userID,
		uint(portfolioID),
		req,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"portfolio berhasil diperbarui",
		result,
	)
}

func (h *PortfolioHandler) Delete(c *echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	portfolioID, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		return helper.BadRequest(
			c,
			"invalid portfolio id",
		)
	}

	if err := h.portfolioUsecase.Delete(
		c.Request().Context(),
		userID,
		uint(portfolioID),
	); err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"portfolio berhasil dihapus",
		nil,
	)
}
