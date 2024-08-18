**Задание:**

Требуется создать микросервис сокращения url.

**Запуск приложения:**

**1. Запуск приложения с базой данных Postgres:**

`docker-compose up`

ИЛИ

`go mod tidy`

`go run cmd/main.go -d`

**Применение миграций:**

`migrate -path ./migrations -database postgres://postgres:mypass@localhost:5436/test?sslmode=disable up`

**2. Запуск приложения с локальным хранилищем:**

`go mod tidy`

`go run cmd/main.go`

**Документация приложения:**

http://localhost:8080/swagger/index.html
