package handler

import (
	"fmt"
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

func (h *Handler) getProducts(c *gin.Context) {
	products, err := h.services.GetAll()
	if err != nil {
		fmt.Println("it fails here h/pr")
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, types.Products{
		Products: products,
	})
}

func (h *Handler) getProductById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id := c.Param("id")

	list, err := h.services.Product.GetById(userId, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) deleteProductById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id := c.Param("id")

	err = h.services.Product.Delete(userId, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

func (h *Handler) updateProductById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id := c.Param("id")

	var input types.UpdateProduct
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Product.Update(userId, id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
