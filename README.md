# Book Library API

A modern Go-based REST API for managing a personal book library. This project serves as a demonstration of contemporary backend development practices and technologies.

## 🚀 Technologies

- Go 1.22+
- Gin Web Framework
- GORM
- PostgreSQL
- Docker & Docker Compose
- Swagger/OpenAPI
- JWT Authentication
- Unit & Integration Testing

## 📁 Project Structure

```
.
├── cmd/
│   └── api/            # Application entrypoint
├── internal/
│   ├── api/           # API handlers
│   ├── config/        # Configuration
│   ├── middleware/    # HTTP middleware
│   ├── model/         # Domain models
│   ├── repository/    # Data access layer
│   └── service/       # Business logic
├── pkg/               # Public packages
├── docs/             # Documentation
└── scripts/          # Helper scripts
```

## 🛠️ Setup and Running

### Prerequisites

- Go 1.22 or higher
- Docker and Docker Compose
- Make (optional)

### Local Development

1. Clone the repository:
```bash
git clone https://github.com/yourusername/book-library-api
cd book-library-api
```

2. Start the database:
```bash
docker-compose up -d postgres
```

3. Run the application:
```bash
go run cmd/api/main.go
```

### Using Docker

```bash
docker-compose up --build
```

## 📝 API Documentation

Once the application is running, you can access the Swagger documentation at:
```
http://localhost:8080/swagger/index.html
```

## 🧪 Running Tests

```bash
go test ./...
```

## 📜 License

This project is licensed under the MIT License - see the LICENSE file for details. 