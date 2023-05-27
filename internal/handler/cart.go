package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online_shop/internal/types"
)

func (h *Handler) add(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input types.AddToCart
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Cart.Add(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

func (h *Handler) delete(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input types.RemoveFromCart
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Cart.Delete(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
