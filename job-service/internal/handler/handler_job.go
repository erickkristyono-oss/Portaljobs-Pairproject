package handler

import (
	"errors"
	"net/http"
	"strconv"

	domainerrors "job-service/internal/domain/error"
	"job-service/internal/dto/request"
	"job-service/internal/dto/response"
	"job-service/internal/helper"
	"job-service/internal/middleware"
	jobUsecase "job-service/internal/usecase/job"

	"github.com/labstack/echo/v5"
)

type JobHandler struct {
	jobUsecase jobUsecase.JobUsecase
}

func NewJobHandler(jobUsecase jobUsecase.JobUsecase) *JobHandler {
	return &JobHandler{
		jobUsecase: jobUsecase,
	}
}

// Create membuat lowongan baru.
// POST /jobs
func (h *JobHandler) Create(c *echo.Context) error {
	var req request.CreateJobRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(c, "invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return helper.BadRequest(c, err.Error())
	}

	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	job, err := h.jobUsecase.Create(
		c.Request().Context(),
		userID,
		req,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return c.JSON(http.StatusCreated, helper.Response{
		ResponseCode:    http.StatusCreated,
		ResponseMessage: "job created successfully",
		ResponseData:    response.FromEntity(job),
	})
}

// GetAll mengambil semua lowongan.
// GET /jobs
func (h *JobHandler) GetAll(c *echo.Context) error {

	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	// userID digunakan untuk memastikan user sudah login.
	// Semua user yang memiliki JWT dapat melihat job published.
	_ = userID

	jobs, err := h.jobUsecase.GetAll(
		c.Request().Context(),
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	data := make([]response.JobResponse, 0, len(jobs))

	for _, job := range jobs {
		data = append(
			data,
			response.FromEntity(job),
		)
	}

	return c.JSON(http.StatusOK, helper.Response{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "published jobs retrieved successfully",
		ResponseData:    data,
	})
}

// GetByID mengambil detail job berdasarkan ID.
// GET /jobs/:id
func (h *JobHandler) GetByID(c *echo.Context) error {
	id, err := getIDParam(c)
	if err != nil {
		return helper.BadRequest(c, "invalid job id")
	}

	job, err := h.jobUsecase.GetByID(
		c.Request().Context(),
		id,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return c.JSON(http.StatusOK, helper.Response{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "success",
		ResponseData:    response.FromEntity(job),
	})
}

// Update mengubah data job.
// PUT /jobs/:id
func (h *JobHandler) Update(c *echo.Context) error {
	id, err := getIDParam(c)
	if err != nil {
		return helper.BadRequest(c, "invalid job id")
	}

	var req request.UpdateJobRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(c, "invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return helper.BadRequest(c, err.Error())
	}

	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	job, err := h.jobUsecase.Update(
		c.Request().Context(),
		userID,
		id,
		req,
	)
	if err != nil {
		return handleDomainError(c, err)
	}

	return c.JSON(http.StatusOK, helper.Response{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "job updated successfully",
		ResponseData:    response.FromEntity(job),
	})
}

// Delete menghapus job.
// DELETE /jobs/:id
func (h *JobHandler) Delete(c *echo.Context) error {
	id, err := getIDParam(c)
	if err != nil {
		return helper.BadRequest(c, "invalid job id")
	}

	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	if err := h.jobUsecase.Delete(
		c.Request().Context(),
		userID,
		id,
	); err != nil {
		return handleDomainError(c, err)
	}

	return c.JSON(http.StatusOK, helper.Response{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "job deleted successfully",
	})
}

// Publish mengubah status job menjadi published.
// PATCH /jobs/:id/publish
func (h *JobHandler) Publish(c *echo.Context) error {
	id, err := getIDParam(c)
	if err != nil {
		return helper.BadRequest(c, "invalid job id")
	}

	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	if err := h.jobUsecase.Publish(
		c.Request().Context(),
		userID,
		id,
	); err != nil {
		return handleDomainError(c, err)
	}

	return c.JSON(http.StatusOK, helper.Response{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "job published successfully",
	})
}

// Close mengubah status job menjadi closed.
// PATCH /jobs/:id/close
func (h *JobHandler) Close(c *echo.Context) error {
	id, err := getIDParam(c)
	if err != nil {
		return helper.BadRequest(c, "invalid job id")
	}

	userID, err := getUserID(c)
	if err != nil {
		return err
	}

	if err := h.jobUsecase.Close(
		c.Request().Context(),
		userID,
		id,
	); err != nil {
		return handleDomainError(c, err)
	}

	return c.JSON(http.StatusOK, helper.Response{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "job closed successfully",
	})
}

// getUserID mengambil user ID dari JWT middleware.
func getUserID(c *echo.Context) (uint, error) {
	userIDValue := c.Get(middleware.UserIDKey)

	userID, ok := userIDValue.(uint)
	if !ok || userID == 0 {
		return 0, echo.NewHTTPError(
			http.StatusUnauthorized,
			"unauthorized",
		)
	}

	return userID, nil
}

// getIDParam mengambil ID dari URL.
func getIDParam(c *echo.Context) (uint, error) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id == 0 {
		return 0, err
	}

	return uint(id), nil
}

// handleDomainError mengubah domain error menjadi HTTP response.
func handleDomainError(c *echo.Context, err error) error {
	switch {
	case errors.Is(err, domainerrors.ErrNotFound):
		return helper.NotFound(c, "job not found")

	case errors.Is(err, domainerrors.ErrForbidden):
		return helper.Forbidden(c, "forbidden")

	case errors.Is(err, domainerrors.ErrConflict):
		return helper.Conflict(c, "job conflict")

	default:
		return helper.InternalServerError(c, "internal server error")
	}
}
