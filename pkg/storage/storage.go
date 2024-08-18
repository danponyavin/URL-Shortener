package storage

type URLStorage interface {
}

type LocalStorage struct {
	data map[string]string
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{data: make(map[string]string)}
}
