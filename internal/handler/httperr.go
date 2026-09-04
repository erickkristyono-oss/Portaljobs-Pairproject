package handler

import (
	"net/http"

	"portaljob/internal/domain"
	"portaljob/pkg/response"

	"github.com/gin-gonic/gin"
)

// fail maps domain sentinel errors to the correct HTTP status code so every
// handler stays consistent (REST best-practice: right status for right error).
func fail(c *gin.Context, err error) {
	switch err {
	case domain.ErrNotFound:
		response.Error(c, http.StatusNotFound, err.Error())
	case domain.ErrEmailTaken, domain.ErrConflict:
		response.Error(c, http.StatusConflict, err.Error())
	case domain.ErrInvalidCredential:
		response.Error(c, http.StatusUnauthorized, err.Error())
	case domain.ErrForbidden, domain.ErrSuspended:
		response.Error(c, http.StatusForbidden, err.Error())
	case domain.ErrInvalidRole, domain.ErrInvalidLevel, domain.ErrNoCompanyProfile:
		response.Error(c, http.StatusBadRequest, err.Error())
	default:
		response.Error(c, http.StatusInternalServerError, "internal error")
	}
}
