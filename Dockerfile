FROM golang:1.23.7 AS builder
WORKDIR /app

# Copy and build
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weather-app ./cmd
#RUN go build -o weather-app ./cmd

# Step 2: Run the binary in a lightweight container
FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/weather-app .
EXPOSE 8080

CMD ["./weather-app"]