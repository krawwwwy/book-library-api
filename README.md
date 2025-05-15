# Библиотека книг API

Современный REST API на Go для управления персональной библиотекой книг. Этот проект демонстрирует современные практики и технологии backend-разработки.

## 🚀 Технологии

- Go 1.22+
- Gin Web Framework
- GORM
- PostgreSQL
- Docker & Docker Compose
- Swagger/OpenAPI
- JWT Аутентификация
- Модульное и интеграционное тестирование

## 📁 Структура проекта

```
.
├── cmd/
│   └── api/            # Точка входа приложения
├── internal/
│   ├── api/           # Обработчики API
│   ├── config/        # Конфигурация
│   ├── middleware/    # HTTP промежуточное ПО
│   ├── model/         # Модели данных
│   ├── repository/    # Слой доступа к данным
│   └── service/       # Бизнес-логика
├── pkg/               # Публичные пакеты
├── docs/             # Документация
└── scripts/          # Вспомогательные скрипты
```

## 🛠️ Установка и запуск

### Предварительные требования

- Go 1.22 или выше
- Docker и Docker Compose
- Make (опционально)

### Локальная разработка

1. Клонировать репозиторий:
```bash
git clone https://github.com/krawwwwy/book-library-api
cd book-library-api
```

2. Запустить базу данных:
```bash
docker-compose up -d postgres
```

3. Запустить приложение:
```bash
go run cmd/api/main.go
```

### Использование Docker

```bash
docker-compose up --build
```

## 📝 Документация API

После запуска приложения документация Swagger доступна по адресу:
```
http://localhost:8080/swagger/index.html
```

## 🧪 Запуск тестов

```bash
go test ./...
```

## 📜 Лицензия

Этот проект распространяется под лицензией MIT - подробности см. в файле LICENSE. 