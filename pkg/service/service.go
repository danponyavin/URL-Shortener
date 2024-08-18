package service

import "URL-Shortener/pkg/storage"

type Shortener interface {
	ShortenUrl(url string) (string, error)
}

type Services struct {
	Shortener
}

func NewServices(storage storage.URLStorage) *Services {
	return &Services{Shortener: NewURLShortener(storage)}
}
