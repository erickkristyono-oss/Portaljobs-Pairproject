package handler

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v5"

	"application-service/internal/domain/constant"
	domainerrors "application-service/internal/domain/error"
	"application-service/internal/dto/request"
	"application-service/internal/dto/response"
	"application-service/internal/helper"
	application "application-service/internal/usecase/app"
)

type ApplicationHandler struct {
	usecase application.ApplicationUsecase
}

func NewApplicationHandler(
	usecase application.ApplicationUsecase,
) *ApplicationHandler {
	return &ApplicationHandler{
		usecase: usecase,
	}
}

// Apply godoc
// POST /applications
func (h *ApplicationHandler) Apply(c *echo.Context) error {

	userID, err := getUserID(c)
	if err != nil {
		return helper.Unauthorized(
			c,
			"unauthorized",
		)
	}

	var req request.CreateApplicationRequest

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

	applicationData, err := h.usecase.Apply(
		c.Request().Context(),
		userID,
		req,
	)
	if err != nil {
		return h.handleError(c, err)
	}

	return helper.Created(
		c,
		"application created successfully",
		response.FromEntity(applicationData),
	)
}

// GetMyApplications godoc
// GET /applications/me
func (h *ApplicationHandler) GetMyApplications(c *echo.Context) error {

	userID, err := getUserID(c)
	if err != nil {
		return helper.Unauthorized(
			c,
			"unauthorized",
		)
	}

	applications, err := h.usecase.GetMyApplications(
		c.Request().Context(),
		userID,
	)
	if err != nil {
		return h.handleError(c, err)
	}

	data := make([]response.ApplicationResponse, 0, len(applications))

	for _, applicationData := range applications {
		data = append(
			data,
			response.FromEntity(applicationData),
		)
	}

	return helper.Success(
		c,
		"applications retrieved successfully",
		data,
	)
}

// GetByID godoc
// GET /applications/:id
func (h *ApplicationHandler) GetByID(c *echo.Context) error {

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		return helper.BadRequest(
			c,
			"invalid application id",
		)
	}

	applicationData, err := h.usecase.GetByID(
		c.Request().Context(),
		uint(id),
	)
	if err != nil {
		return h.handleError(c, err)
	}

	return helper.Success(
		c,
		"application retrieved successfully",
		response.FromEntity(applicationData),
	)
}

// GetApplicationsByJobID godoc
// GET /jobs/:job_id/applications
func (h *ApplicationHandler) GetApplicationsByJobID(c *echo.Context) error {

	userID, err := getUserID(c)
	if err != nil {
		return helper.Unauthorized(
			c,
			"unauthorized",
		)
	}

	jobID, err := strconv.ParseUint(
		c.Param("job_id"),
		10,
		64,
	)
	if err != nil {
		return helper.BadRequest(
			c,
			"invalid job id",
		)
	}

	applications, err := h.usecase.GetApplicationsByJobID(
		c.Request().Context(),
		userID,
		uint(jobID),
	)
	if err != nil {
		return h.handleError(c, err)
	}

	data := make(
		[]response.ApplicationResponse,
		0,
		len(applications),
	)

	for _, applicationData := range applications {
		data = append(
			data,
			response.FromEntity(applicationData),
		)
	}

	return helper.Success(
		c,
		"applications retrieved successfully",
		data,
	)
}

// UpdateStatus godoc
// PATCH /applications/:id/status
func (h *ApplicationHandler) UpdateStatus(c *echo.Context) error {

	userID, err := getUserID(c)
	if err != nil {
		return helper.Unauthorized(
			c,
			"unauthorized",
		)
	}

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		return helper.BadRequest(
			c,
			"invalid application id",
		)
	}

	var req struct {
		Status string `json:"status" validate:"required"`
	}

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

	if !constant.IsValidApplicationStatus(req.Status) {
		return helper.BadRequest(
			c,
			"invalid application status",
		)
	}

	err = h.usecase.UpdateStatus(
		c.Request().Context(),
		userID,
		uint(id),
		req.Status,
	)
	if err != nil {
		return h.handleError(c, err)
	}

	return helper.Success(
		c,
		"application status updated successfully",
		nil,
	)
}

// Delete godoc
// DELETE /applications/:id
func (h *ApplicationHandler) Delete(c *echo.Context) error {

	userID, err := getUserID(c)
	if err != nil {
		return helper.Unauthorized(
			c,
			"unauthorized",
		)
	}

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		return helper.BadRequest(
			c,
			"invalid application id",
		)
	}

	err = h.usecase.Delete(
		c.Request().Context(),
		userID,
		uint(id),
	)
	if err != nil {
		return h.handleError(c, err)
	}

	return helper.Success(
		c,
		"application deleted successfully",
		nil,
	)
}

func getUserID(c *echo.Context) (uint, error) {

	value := c.Get("user_id")

	userID, ok := value.(uint)
	if !ok {
		return 0, errors.New("user id not found")
	}

	return userID, nil
}

func (h *ApplicationHandler) handleError(c *echo.Context, err error) error {

	switch {
	case errors.Is(err, domainerrors.ErrNotFound):
		return helper.NotFound(
			c,
			"application not found",
		)

	case errors.Is(err, domainerrors.ErrForbidden):
		return helper.Forbidden(
			c,
			"you are not allowed to access this application",
		)

	case errors.Is(err, domainerrors.ErrConflict):
		return helper.Conflict(
			c,
			"application conflict",
		)

	default:
		return helper.InternalServerError(
			c,
			"internal server error",
		)
	}
}
