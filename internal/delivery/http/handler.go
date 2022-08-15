package http

import (
	"github.com/gin-gonic/gin"
	"github.com/onemgvv/wb-l0/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	api := router.Group("/api")
	{
		api.GET("/order/:id", h.GetByID)
	}

	return router
}
