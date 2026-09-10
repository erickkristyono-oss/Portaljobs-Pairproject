package handler

import (
	"net/http"

	"portaljob/internal/usecase"
	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{ auth *usecase.AuthUsecase }

func NewAuthHandler(auth *usecase.AuthUsecase) *AuthHandler { return &AuthHandler{auth: auth} }

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.auth.Register(c.Request.Context(), req.Nama, req.Email, req.Password, req.Role)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusCreated, "registered", user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	token, user, err := h.auth.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, http.StatusOK, "login success", gin.H{"token": token, "user": user})
}
