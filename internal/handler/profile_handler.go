package handler

import (
	"net/http"

	"portaljob/internal/domain"
	"portaljob/internal/middleware"
	"portaljob/internal/usecase"
	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct{ profiles *usecase.ProfileUsecase }

func NewProfileHandler(profiles *usecase.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{profiles: profiles}
}

func (h *ProfileHandler) Save(c *gin.Context) {
	var req ProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	p := domain.Profile{
		Name: req.Name, PhoneNumber: req.PhoneNumber, Address: req.Address,
		Faculty: req.Faculty, Major: req.Major, EducationLevel: req.EducationLevel,
		Started: req.Started, Graduated: req.Graduated,
	}
	saved, err := h.profiles.Save(c.Request.Context(), middleware.UserID(c), p)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "profile saved", saved)
}

func (h *ProfileHandler) Get(c *gin.Context) {
	p, err := h.profiles.Get(c.Request.Context(), middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "profile", p)
}
