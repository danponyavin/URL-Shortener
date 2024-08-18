package handler

import (
	"URL-Shortener/docs"
	"URL-Shortener/pkg/models"
	"URL-Shortener/pkg/service"
	"URL-Shortener/pkg/storage"
	"errors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

var serviceError = "Service error"

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

const BasePath = "/"

// @BasePath /
func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = BasePath
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.POST("/create", h.CreateShortURL)
	r.GET("/:short_url", h.GetOriginalURL)

	return r
}

type Error struct {
	Message string `json:"message"`
}

// RefreshTokens godoc
// @Summary Создание короткой ссылки
// @Schemes
// @Description Создание короткой ссылки
// @Accept json
// @Produce json
// @Param data body models.UrlModel true "Входные параметры"
// @Success 200 {object} models.UrlModel
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /create [post]
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

// GetTokens godoc
// @Summary Получение оригинальной ссылки
// @Schemes
// @Description Получение оригинальной ссылки
// @Accept json
// @Produce json
// @Param short_url path string true "Short URL" Example(dkmNJ2x)
// @Success 200 {object} models.UrlModel
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /{short_url} [get]
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
