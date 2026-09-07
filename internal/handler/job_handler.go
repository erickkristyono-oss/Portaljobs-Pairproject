package handler

import (
	"net/http"
	"strconv"

	"portaljob/internal/middleware"
	"portaljob/internal/usecase"
	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

type JobHandler struct{ jobs *usecase.JobUsecase }

func NewJobHandler(jobs *usecase.JobUsecase) *JobHandler { return &JobHandler{jobs: jobs} }

func (h *JobHandler) Create(c *gin.Context) {
	var req CreateJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	job, err := h.jobs.Create(c.Request.Context(), middleware.UserID(c),
		req.Judul, req.AboutRole, req.Responsibilities, req.Deskripsi, req.Lokasi, req.Gaji, req.RequiredSkills)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusCreated, "job created", job)
}

func (h *JobHandler) List(c *gin.Context) {
	lokasi := c.Query("lokasi")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	jobs, err := h.jobs.List(c.Request.Context(), lokasi, limit, offset)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "jobs", jobs)
}

func (h *JobHandler) Detail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid job id")
		return
	}
	job, err := h.jobs.Detail(c.Request.Context(), uint(id))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "job detail", job)
}

func (h *JobHandler) ListMine(c *gin.Context) {
	jobs, err := h.jobs.ListMine(c.Request.Context(), middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "my jobs", jobs)
}
