package service

import "URL-Shortener/pkg/storage"

type URLShortener struct {
	storage storage.URLStorage
}

func NewURLShortener(urlStorage storage.URLStorage) *URLShortener {
	return &URLShortener{urlStorage}
}
