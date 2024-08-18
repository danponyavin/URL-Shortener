package handler

import (
	"URL-Shortener/pkg/models"
	"URL-Shortener/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var serviceError = "Service error"

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/create", h.CreateShortURL)
	r.GET("/:short_url", h.GetLongURL)

	return r
}

type Error struct {
	Message string `json:"message"`
}

func (h *Handler) CreateShortURL(c *gin.Context) {
	var req models.UrlModel

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
		return
	}

	shortURL, err := h.services.Shortener.ShortenUrl(req.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{Message: serviceError})
		return
	}

	c.JSON(http.StatusOK, models.UrlModel{Url: shortURL})
}

func (h *Handler) GetLongURL(c *gin.Context) {

}
