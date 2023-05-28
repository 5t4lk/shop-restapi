package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online_shop/internal/types"
)

func (h *Handler) placeOrder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	cartItems, err := h.services.Cart.GetAll(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	cart := types.ShoppingCart{
		UserID: userId,
		Items:  cartItems,
	}

	orderID, err := h.services.Order.Place(userId, cart)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"order_id": orderID,
	})
}

func (h *Handler) deleteOrder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id := c.Param("id")

	err = h.services.Order.Delete(userId, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
