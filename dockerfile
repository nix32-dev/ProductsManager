# Dockerfile
FROM golang:1.25.7-bookworm

WORKDIR /app

# Копируем go.mod и go.sum первыми — для кэширования слоя
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o pmanager .

# Запускаем бинарник напрямую
CMD ["./pmanager"]