package handler

import (
	"net/http"

	"portaljob/internal/middleware"
	"portaljob/internal/usecase"
	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

type SkillTagHandler struct{ tags *usecase.SkillTagUsecase }

func NewSkillTagHandler(tags *usecase.SkillTagUsecase) *SkillTagHandler {
	return &SkillTagHandler{tags: tags}
}

// ListVocab — GET /api/v1/skill-tags (public): the shared vocabulary.
func (h *SkillTagHandler) ListVocab(c *gin.Context) {
	items, err := h.tags.ListVocab(c.Request.Context())
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "skill tags", items)
}

// GetMine — GET /api/v1/me/skill-tags (jobseeker).
func (h *SkillTagHandler) GetMine(c *gin.Context) {
	items, err := h.tags.GetUserTags(c.Request.Context(), middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "my skill tags", items)
}

// SetMine — PUT /api/v1/me/skill-tags (jobseeker): replace tag list by names.
func (h *SkillTagHandler) SetMine(c *gin.Context) {
	var req SkillTagsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	items, err := h.tags.SetUserTags(c.Request.Context(), middleware.UserID(c), req.Skills)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "skill tags saved", items)
}
