
# Weather Service API

This project is a **Go-based service** that provides weather-related functionality with validation, service layer, and API handlers.  
It follows a **clean architecture approach** with separate layers for configuration, models, services, validators, and handlers.  

## 📂 Project Structure

---
```
.
├── cmd/                   # Application entrypoint
│   └── main.go
│
├── config/                # Configuration files
│   └── config.go
│
├── helm/                  # Environment configs
│   ├── local.yaml
│   └── prod.yaml
│
├── internal/              # Core application logic
│   ├── handler/           # HTTP handlers (API endpoints)
│   │   ├── handler.go
│   │   └── handler_test.go
│   │
│   ├── metrics/           # Metrics exposition (Prometheus, etc.)
│   │   └── metrics.go
│   │
│   ├── models/            # Data models (structs, DB models, etc.)
│   │   └── model.go
│   │
│   ├── routes/            # HTTP route registration
│   │   └── routes.go
│   │
│   ├── service/           # Business logic (core services)
│   │   ├── service.go
│   │   └── service_test.go
│   │
│   └── validator/         # Request validation logic
│       └── weather_validation.go
|       └── weather_validation_test.go
│
├── app/                   # API specification
│   └── swagger.yaml
│
├── tests/                 # Integration tests
│   └── api_integration_test.go
│
├── Dockerfile             # Docker build config
├── docker-compose.yml     # Compose file for local dev
├── Makefile               # Build/test automation
├── go.mod / go.sum        # Go dependencies
└── README.md              # Project documentation

```
---

## 🚀 Getting Started

### Prerequisites
- [Go 1.20+](https://go.dev/dl/)
- [Docker](https://www.docker.com/)
- [Make](https://www.gnu.org/software/make/)

### Setup & Run
---
```bash
# Clone repo
git clone <repo-url>
cd <repo-name>

# Run locally
go run ./cmd/main.go

# Or use docker-compose
rm -rf bin/                                                    
        docker compose down --rmi all -v

docker compose build

docker compose up -d

docker run -it --entrypoint sh weather-app-weather-app
```
---

The service will start on the port defined in `config/local.yaml`.

## 🧪 Testing

Unit and integration tests are included. Run them with:

---
```bash
# Run all tests
go test ./...

# Run integration tests
go test ./tests/...
```
---

## 📖 API Documentation

The API is documented using **Swagger**.  
File: [`app/swagger.yaml`](app/swagger.yaml)  

You can view it locally with:
---
```bash
docker run -p 8080:8080 -e SWAGGER_JSON=/swagger.yaml     -v $(pwd)/app/swagger.yaml:/swagger.yaml swaggerapi/swagger-ui
```
---
Then open [http://localhost:8080](http://localhost:8080) in your browser.

## ⚙️ Features

- Weather data validation layer
- Service-oriented architecture
- Configurable via YAML (`local.yaml`, `prod.yaml`)
- Unit tests + integration tests
- Docker & Makefile support
- Swagger API specification

## 📌 Notes

- `internal/` contains the core application logic, isolated from external frameworks.  
- `tests/` focuses on integration-level testing of API endpoints.  
- The project is structured for **scalability and maintainability**, making it easy to extend with new services or APIs.
