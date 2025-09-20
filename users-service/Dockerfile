FROM golang:1.21-alpine AS builder

WORKDIR /app

# Устанавливаем необходимые пакеты
RUN apk add --no-cache git

# Копируем go mod файлы
COPY go.mod go.sum ./

# Загружаем зависимости с отключенной проверкой сумм
RUN GOSUMDB=off go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN GOSUMDB=off go build -o users-service ./cmd/server

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарный файл
COPY --from=builder /app/users-service .

# Открываем порт
EXPOSE 50051

# Запускаем приложение
CMD ["./users-service"]
