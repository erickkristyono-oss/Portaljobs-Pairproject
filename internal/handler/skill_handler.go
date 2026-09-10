package handler

import (
	"net/http"
	"strconv"

	"portaljob/internal/domain"
	"portaljob/internal/middleware"
	"portaljob/internal/usecase"
	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

type SkillHandler struct{ skills *usecase.SkillUsecase }

func NewSkillHandler(skills *usecase.SkillUsecase) *SkillHandler {
	return &SkillHandler{skills: skills}
}

func (h *SkillHandler) Add(c *gin.Context) {
	var req SkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	s := domain.Skill{
		NameLicense: req.NameLicense, Level: domain.SkillLevel(req.Level),
		Organization: req.Organization, Grade: req.Grade,
		ExpiredDate: req.ExpiredDate, Description: req.Description,
	}
	saved, err := h.skills.Add(c.Request.Context(), middleware.UserID(c), s)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusCreated, "skill added", saved)
}

func (h *SkillHandler) List(c *gin.Context) {
	items, err := h.skills.List(c.Request.Context(), middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "skills", items)
}

func (h *SkillHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid skill id")
		return
	}
	if err := h.skills.Remove(c.Request.Context(), middleware.UserID(c), uint(id)); err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "skill deleted", nil)
}
