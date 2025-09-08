# ---------- build stage ----------
FROM golang:1.22-alpine AS build
WORKDIR /src

# (опционально) если будут приватные модули — нужен git
RUN apk add --no-cache git

# сначала зависимости (лучше кешируется)
COPY go.mod ./
RUN go mod download

# теперь исходники
COPY . .

# собираем статический бинарник
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o app .

# ---------- run stage ----------
# distroless уже содержит CA-сертификаты и юзера nonroot
FROM scratch
WORKDIR /app
# нужны корневые сертификаты для HTTPS-запросов
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/app /app/app
EXPOSE 8080
ENTRYPOINT ["/app/app"]