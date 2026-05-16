# Go SQLite CRUD API with Swagger

This is a professional-grade CRUD API built with Go, SQLite, and Swagger documentation.

## Features
- **Clean Architecture**: Decoupled layers for handlers, services, and repositories.
- **SQLite & GORM**: Lightweight database with easy schema management.
- **Swagger Documentation**: Automatic API docs available at `/swagger/index.html`.
- **Middleware**: Includes logging and recovery middleware.

## Getting Started

### 1. Initialize Module
```bash
go mod init github.com/ahsan/go-sqlite-crud
go mod tidy
```

### 2. Install Swagger (Swag)
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 3. Generate Swagger Documentation
```bash
swag init -g cmd/api/main.go
```

### 4. Run the Application
```bash
go run cmd/api/main.go
```

### 5. Access the API
- **API Base URL**: `http://localhost:8080/api/v1`
- **Swagger UI**: `http://localhost:8080/swagger/index.html`

## Directory Structure
- `cmd/api`: Main entry point.
- `internal/handler`: HTTP request handlers and Swagger annotations.
- `internal/service`: Business logic layer.
- `internal/repository`: Database access layer (GORM).
- `internal/model`: Data models/entities.
- `pkg/database`: Database connection and migration.

### Cara Melakukan Update (Generate Ulang)

```bash
~/go/bin/swag init -g cmd/api/main.go
```