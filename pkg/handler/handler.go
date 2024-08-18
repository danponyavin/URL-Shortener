package handler

import (
	"URL-Shortener/pkg/models"
	"URL-Shortener/pkg/service"
	"URL-Shortener/pkg/storage"
	"errors"
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
	r.GET("/:short_url", h.GetOriginalURL)

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

func (h *Handler) GetOriginalURL(c *gin.Context) {
	shortURL := c.Param("short_url")
	if shortURL == "" {
		c.JSON(http.StatusBadRequest, Error{Message: "Empty short_url"})
		return
	}

	origURL, err := h.services.Shortener.GetOriginalUrl(shortURL)
	if err != nil {
		switch {
		case errors.Is(err, storage.UrlNotFoundError):
			c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, Error{Message: serviceError})
			return
		}
	}

	c.JSON(http.StatusOK, models.UrlModel{Url: origURL})
}
