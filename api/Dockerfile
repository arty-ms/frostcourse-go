# ---------- build stage ----------
FROM golang:1.22-alpine AS build
WORKDIR /src

# (опционально) если будут приватные модули — нужен git
RUN apk add --no-cache git

# сначала зависимости (лучше кешируется)
COPY go.mod ./
RUN go mod download

COPY api/ ./api/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o server ./api

FROM scratch
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/server /app/server
EXPOSE 8080
ENTRYPOINT ["/app/server"]