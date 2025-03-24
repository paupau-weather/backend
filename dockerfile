# Используем официальный образ Go в качестве базового
FROM golang:1.23.5-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum, чтобы воспользоваться кэшированием зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь исходный код приложения в контейнер
COPY . .

# Собираем приложение
RUN go build -o main .

# Создаем минимальный образ для запуска приложения
FROM alpine:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем исполняемый файл из предыдущего этапа сборки
COPY --from=builder /app/main .

# Указываем порт, который приложение будет слушать (если нужно)
EXPOSE 5000

# Команда для запуска приложения
CMD ["./main"]