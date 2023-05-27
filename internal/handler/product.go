package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online_shop/internal/types"
)

func (h *Handler) createProduct(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input types.CreateProduct
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	productId, err := h.services.Product.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": productId,
	})
}
