# Используем образ Golang 1.23 на базе Alpine
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Устанавливаем необходимые зависимости
RUN apk add --no-cache curl bash git

# Устанавливаем swag для генерации документации Swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Загружаем и устанавливаем migrate для миграций
RUN apk add --no-cache curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xz -C /usr/local/bin

# Копируем go.mod и go.sum в контейнер и выполняем go mod tidy
COPY go.mod go.sum ./
RUN go mod tidy

# Копируем остальные файлы приложения
COPY . .

# Генерируем Swagger-документацию
RUN swag init

# Делаем миграции базы данных, если необходимо (при наличии скрипта migrate.sh)
RUN chmod +x ./migrate.sh

# Собираем Go-приложение
RUN go build -o myapp .

# Запускаем приложение
CMD ["./myapp"]
