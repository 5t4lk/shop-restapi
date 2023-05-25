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

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	v1 := router.Group("/v1")
	{
		h.InitShopRouter(v1)
	}

	return router
}
