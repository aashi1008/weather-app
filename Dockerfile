FROM golang:1.23.7 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o weather-app ./cmd/main.go

# static image has nothing except the kernel helpers
FROM gcr.io/distroless/static
COPY --from=builder /app/weather-app /weather-app

EXPOSE 8080
CMD ["/weather-app"]