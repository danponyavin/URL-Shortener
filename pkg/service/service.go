package service

import "URL-Shortener/pkg/storage"

type Shortener interface {
}

type Services struct {
	Shortener
}

func NewServices(storage storage.URLStorage) *Services {
	return &Services{Shortener: NewURLShortener(storage)}
}
