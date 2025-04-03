# Используем официальный образ Go в качестве базового
FROM golang:1.23.2 AS builder

# Создаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы приложения
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Используем минимальный образ для запуска
FROM alpine:latest

# Устанавливаем необходимые библиотеки
RUN apk --no-cache add ca-certificates

# Копируем собранное приложение из предыдущего этапа
COPY --from=builder /app/main /usr/local/bin/app/main

# Устанавливаем права на выполнение
RUN chmod +x /usr/local/bin/app/main

# Указываем команду для запуска приложения
CMD ["/usr/local/bin/app/main"]