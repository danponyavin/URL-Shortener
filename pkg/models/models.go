package models

type UrlModel struct {
	Url string `json:"url" binding:"required"`
}
