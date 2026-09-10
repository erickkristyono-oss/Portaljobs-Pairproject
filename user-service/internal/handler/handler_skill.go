package handler

import (
	"net/http"
	"strconv"

	"user-service/internal/dto/request"
	"user-service/internal/helper"
	skillUsecase "user-service/internal/usecase/skill"

	"github.com/labstack/echo/v5"
)

type SkillHandler struct {
	skillUsecase skillUsecase.SkillUsecase
}

func NewSkillHandler(
	skillUsecase skillUsecase.SkillUsecase,
) *SkillHandler {
	return &SkillHandler{
		skillUsecase: skillUsecase,
	}
}

func (h *SkillHandler) Create(c *echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	var req request.CreateSkillRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(
			c,
			"invalid request body",
		)
	}

	result, err := h.skillUsecase.Create(
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
		"skill berhasil ditambahkan",
		result,
	)
}

func (h *SkillHandler) GetByUserID(c *echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	result, err := h.skillUsecase.GetByUserID(
		c.Request().Context(),
		userID,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"skill berhasil ditemukan",
		result,
	)
}

func (h *SkillHandler) Update(c *echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	skillID, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		return helper.BadRequest(
			c,
			"invalid skill id",
		)
	}

	var req request.UpdateSkillRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(
			c,
			"invalid request body",
		)
	}

	result, err := h.skillUsecase.Update(
		c.Request().Context(),
		userID,
		uint(skillID),
		req,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"skill berhasil diperbarui",
		result,
	)
}

func (h *SkillHandler) Delete(c *echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	skillID, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		return helper.BadRequest(
			c,
			"invalid skill id",
		)
	}

	if err := h.skillUsecase.Delete(
		c.Request().Context(),
		userID,
		uint(skillID),
	); err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"skill berhasil dihapus",
		nil,
	)
}

func (h *SkillHandler) GetByUserIDInternal(c *echo.Context) error {
	userID, err := strconv.ParseUint(
		c.Param("user_id"),
		10,
		64,
	)
	if err != nil {
		return helper.BadRequest(
			c,
			"invalid user id",
		)
	}

	result, err := h.skillUsecase.GetByUserID(
		c.Request().Context(),
		uint(userID),
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return helper.Success(
		c,
		http.StatusOK,
		"skill berhasil ditemukan",
		result,
	)
}
