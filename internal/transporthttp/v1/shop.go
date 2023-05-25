package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online_shop/internal/types"
)

func (h *Handler) InitShopRouter(v1 *gin.RouterGroup) {
	auth := v1.Group("/shop")
	{
		auth.POST("/", h.Create)
		auth.GET("/:id", h.GetByID)
		auth.GET("/", h.GetAll)
		auth.DELETE("/:id", h.Delete)
		auth.PUT("/:id", h.Update)
	}
}

func (h *Handler) Create(c *gin.Context) {
	var v types.ProductCreateInput

	if err := c.BindJSON(&v); err != nil {
		newResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := h.services.Shop.Create(c, v); err != nil {
		newResponse(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": v,
	})
}

func (h *Handler) GetAll(c *gin.Context) {

}

func (h *Handler) GetByID(c *gin.Context) {

}

func (h *Handler) Delete(c *gin.Context) {

}

func (h *Handler) Update(c *gin.Context) {

}
