package service

import (
	"URL-Shortener/pkg/server"
	"URL-Shortener/pkg/storage"
	"crypto/rand"
	"errors"
	"log"
	"math/big"
)

const (
	chars     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	urlLength = 7
)

type URLShortener struct {
	storage storage.URLStorage
}

func NewURLShortener(urlStorage storage.URLStorage) *URLShortener {
	return &URLShortener{urlStorage}
}

func (u *URLShortener) ShortenUrl(url string) (string, error) {
	var shortenedURL string

	for {
		shortenedURL = generateURL()
		_, err := u.storage.GetOriginalUrl(shortenedURL)
		if err != nil {
			if errors.Is(err, storage.UrlNotFoundError) {
				break
			} else {
				log.Println("GetOriginalUrl Error:", err)
				return "", err
			}
		}
	}

	err := u.storage.SaveUrl(shortenedURL, url)
	if err != nil {
		log.Println("SaveUrl Error:", err)
		return "", err
	}

	fullURL := getFullShortenedUrl(shortenedURL)

	return fullURL, nil
}

func generateURL() string {
	shortURL := make([]byte, urlLength)
	for i := range shortURL {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			panic(err)
		}
		shortURL[i] = chars[randomIndex.Int64()]
	}
	return string(shortURL)
}

func getFullShortenedUrl(shortTag string) string {
	return "http://localhost" + server.PORT + "/" + shortTag
}

func (u *URLShortener) GetOriginalUrl(shortTag string) (string, error) {
	url, err := u.storage.GetOriginalUrl(shortTag)
	if err != nil {
		log.Println("GetOriginalUrl Error:", err)
		return "", err
	}

	return url, nil
}
