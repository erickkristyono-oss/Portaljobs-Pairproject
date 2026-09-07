package handler

import (
	"errors"
	"net/http"
	"strconv"

	domainerrors "admin-service/internal/domain/error"
	"admin-service/internal/dto/request"
	"admin-service/internal/dto/response"
	"admin-service/internal/helper"
	"admin-service/internal/mapper"
	"admin-service/internal/usecase/report"

	"github.com/labstack/echo/v5"
)

type ReportHandler struct {
	reportUsecase report.ReportUsecase
}

func NewReportHandler(reportUsecase report.ReportUsecase) *ReportHandler {
	return &ReportHandler{
		reportUsecase: reportUsecase,
	}
}

// Create membuat report baru.
// POST /admin/reports
func (h *ReportHandler) Create(c *echo.Context) error {

	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	var req request.CreateReportRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(
			c,
			"invalid request body",
		)
	}

	if err := c.Validate(&req); err != nil {
		return helper.BadRequest(
			c,
			err.Error(),
		)
	}

	report, err := h.reportUsecase.Create(
		c.Request().Context(),
		userID,
		req,
	)

	if err != nil {
		return handleDomainError(c, err)
	}

	return c.JSON(
		http.StatusCreated,
		helper.Response{
			ResponseCode:    http.StatusCreated,
			ResponseMessage: "report created successfully",
			ResponseData:    mapper.ToReportResponse(report),
		},
	)
}

// FindAll mengambil semua report.
// GET /admin/reports
func (h *ReportHandler) FindAll(c *echo.Context) error {

	reports, err := h.reportUsecase.FindAll(
		c.Request().Context(),
	)

	if err != nil {
		return handleDomainError(c, err)
	}

	data := make(
		[]response.ReportResponse,
		0,
		len(reports),
	)

	for i := range reports {
		data = append(
			data,
			mapper.ToReportResponse(&reports[i]),
		)
	}

	return c.JSON(
		http.StatusOK,
		helper.Response{
			ResponseCode:    http.StatusOK,
			ResponseMessage: "reports retrieved successfully",
			ResponseData:    data,
		},
	)
}

// FindByID mengambil report berdasarkan ID.
// GET /admin/reports/:id
func (h *ReportHandler) FindByID(c *echo.Context) error {

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		return helper.BadRequest(
			c,
			"invalid report id",
		)
	}

	report, err := h.reportUsecase.FindByID(
		c.Request().Context(),
		uint(id),
	)

	if err != nil {
		return handleDomainError(c, err)
	}

	return c.JSON(
		http.StatusOK,
		helper.Response{
			ResponseCode:    http.StatusOK,
			ResponseMessage: "report retrieved successfully",
			ResponseData:    mapper.ToReportResponse(report),
		},
	)
}

// UpdateStatus mengubah status report.
// PATCH /admin/reports/:id/status
func (h *ReportHandler) UpdateStatus(c *echo.Context) error {

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		return helper.BadRequest(
			c,
			"invalid report id",
		)
	}

	var req request.UpdateReportStatusRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(
			c,
			"invalid request body",
		)
	}

	if err := c.Validate(&req); err != nil {
		return helper.BadRequest(
			c,
			err.Error(),
		)
	}

	if err := h.reportUsecase.UpdateStatus(
		c.Request().Context(),
		uint(id),
		req.Status,
	); err != nil {
		return handleDomainError(c, err)
	}

	return c.JSON(
		http.StatusOK,
		helper.Response{
			ResponseCode:    http.StatusOK,
			ResponseMessage: "report status updated successfully",
		},
	)
}

func getUserID(c *echo.Context) (uint, error) {

	userID, ok := c.Get("user_id").(uint)

	if !ok {
		return 0, echo.NewHTTPError(
			http.StatusUnauthorized,
			"unauthorized",
		)
	}

	return userID, nil
}

func handleDomainError(c *echo.Context, err error) error {

	switch {

	case errors.Is(err, domainerrors.ErrNotFound):
		return helper.NotFound(
			c,
			"report not found",
		)

	case errors.Is(err, domainerrors.ErrForbidden):
		return helper.Forbidden(
			c,
			"forbidden",
		)

	case errors.Is(err, domainerrors.ErrConflict):
		return helper.Conflict(
			c,
			"report conflict",
		)

	default:
		return helper.InternalServerError(
			c,
			"internal server error",
		)
	}
}
