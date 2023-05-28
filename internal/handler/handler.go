package handler

import (
	"github.com/gin-gonic/gin"
	"online_shop/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		product := api.Group("/product")
		{
			product.POST("/", h.createProduct)
			product.GET("/", h.getProducts)
			product.GET("/:id", h.getProductById)
			product.DELETE("/:id", h.deleteProductById)
			product.PUT("/:id", h.updateProductById)
		}
		cart := api.Group("/cart")
		{
			cart.POST("/", h.add)
			cart.DELETE("/", h.delete)
			cart.GET("/", h.getAll)
		}
		order := api.Group("/order")
		{
			order.GET("/", h.placeOrder)
			order.DELETE("/:id", h.deleteOrder)
		}
	}

	return router
}
