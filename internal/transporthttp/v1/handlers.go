package v1

import (
	"github.com/gin-gonic/gin"
	"online_shop/internal/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.InitShopRouter(v1)
	}
}
