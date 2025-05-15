@echo off
set DB_HOST=localhost
set DB_PORT=5432
set DB_USER=postgres
set DB_PASSWORD=postgres
set DB_NAME=book_library
set DB_SSLMODE=disable
set SERVER_PORT=8080

go run cmd/api/main.go 