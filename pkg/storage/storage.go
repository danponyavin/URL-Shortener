package storage

import (
	"database/sql"
	"errors"
)

var UrlNotFoundError = errors.New("url not found")

type URLStorage interface {
	SaveUrl(shortUrl string, originalUrl string) error
	GetOriginalUrl(shortUrl string) (string, error)
}

type LocalStorage struct {
	data map[string]string
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{data: make(map[string]string)}
}

func (p *PostgresStorage) SaveUrl(shortUrl string, originalUrl string) error {
	_, err := p.db.Exec("INSERT INTO urls (short_url, original_url) VALUES ($1, $2)", shortUrl, originalUrl)
	if err != nil {
		return err
	}

	return nil
}

func (l *LocalStorage) SaveUrl(shortUrl string, originalUrl string) error {
	l.data[shortUrl] = originalUrl

	return nil
}

func (p *PostgresStorage) GetOriginalUrl(shortUrl string) (string, error) {
	var originalUrl string
	err := p.db.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortUrl).Scan(&originalUrl)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", UrlNotFoundError
		}
		return "", err
	}

	return originalUrl, nil
}

func (l *LocalStorage) GetOriginalUrl(shortUrl string) (string, error) {
	if url, ok := l.data[shortUrl]; ok {
		return url, nil
	}

	return "", UrlNotFoundError
}
