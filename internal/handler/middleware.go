package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	emptyAuthHeader     = "empty auth header"
	invalidAuthHeader   = "invalid auth header"
	userNotFound        = "user is not found"
	invalidUserId       = "user id is of invalid type"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, emptyAuthHeader)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, invalidAuthHeader)
		return
	}

	userId, err := h.services.Validate(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, userNotFound)
	}

	idStr, ok := id.(string)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, invalidUserId)
	}

	return idStr, nil
}
