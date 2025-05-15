# Book Library API

Проект представляет собой REST API для управления библиотекой книг, разработанный на Go с использованием современных практик и технологий.

## Технологии

- Go 1.21.3
- Gin Web Framework
- GORM с PostgreSQL
- Docker и Docker Compose
- Swagger для документации API
- Testify для тестирования

## Функциональность

- CRUD операции для книг
- Поиск книг по названию и автору
- Управление доступностью книг
- Пагинация результатов
- Полнотекстовый поиск с использованием PostgreSQL

## Структура проекта

```
.
├── cmd/
│   └── api/          # Точка входа в приложение
├── internal/
│   ├── api/          # HTTP обработчики
│   ├── config/       # Конфигурация приложения
│   ├── middleware/   # Промежуточное ПО
│   ├── model/        # Модели данных
│   ├── repository/   # Слой доступа к данным
│   └── service/      # Бизнес-логика
├── pkg/              # Публичные пакеты
├── docs/             # Документация
└── scripts/          # Вспомогательные скрипты
```

## Запуск проекта

1. Клонируйте репозиторий:
```bash
git clone https://github.com/yourusername/book-library-api.git
cd book-library-api
```

2. Запустите базу данных:
```bash
docker-compose up -d postgres
```

3. Запустите приложение:
```bash
go run cmd/api/main.go
```

## API Endpoints

- `POST /api/books` - Создание новой книги
- `GET /api/books` - Получение списка книг
- `GET /api/books/:id` - Получение книги по ID
- `PUT /api/books/:id` - Обновление книги
- `DELETE /api/books/:id` - Удаление книги
- `GET /api/books/search` - Поиск книг
- `POST /api/books/:id/toggle-availability` - Изменение доступности книги

## Тестирование

Для запуска тестов:

```bash
# Запуск тестовой базы данных
docker-compose up -d postgres_test

# Запуск всех тестов
go test ./...
```

## Разработка

Проект следует принципам чистой архитектуры и использует:
- Dependency Injection
- Interface-based design
- Unit и интеграционное тестирование
- Swagger документацию
- Docker для разработки и тестирования

## Лицензия

MIT 