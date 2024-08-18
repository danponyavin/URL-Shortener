package handler

import (
	"URL-Shortener/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/create", h.CreateShortURL)
	r.GET("/:short_url", h.GetLongURL)

	return r
}

func (h *Handler) CreateShortURL(c *gin.Context) {

}

func (h *Handler) GetLongURL(c *gin.Context) {

}
