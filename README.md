# Book Library API

Проект представляет собой полноценное приложение для управления библиотекой книг с REST API и пользовательским интерфейсом, разработанное на Go с использованием современных практик и технологий.

## Технологии

### Backend
- Go 1.21.3
- Gin Web Framework
- GORM с PostgreSQL
- Docker и Docker Compose
- Swagger для документации API
- Testify для тестирования

### Frontend
- HTML5, CSS3, JavaScript
- Bootstrap 5
- Fetch API для работы с REST API

## Функциональность

- CRUD операции для книг
- Поиск книг по названию и автору
- Управление доступностью книг
- Пагинация результатов
- Полнотекстовый поиск с использованием PostgreSQL
- Удобный веб-интерфейс для работы с библиотекой

## Структура проекта

```
.
├── cmd/
│   └── api/            # Точка входа в приложение
├── internal/
│   ├── api/            # HTTP обработчики
│   ├── config/         # Конфигурация приложения
│   ├── middleware/     # Промежуточное ПО
│   ├── model/          # Модели данных
│   ├── repository/     # Слой доступа к данным
│   └── service/        # Бизнес-логика
├── public/             # Статические файлы для frontend
│   ├── css/            # CSS стили
│   ├── js/             # JavaScript файлы
│   ├── index.html      # Главная страница
│   └── books.html      # Страница управления книгами
├── pkg/                # Публичные пакеты
├── docs/               # Документация
└── scripts/            # Вспомогательные скрипты
```

## Запуск проекта

Самый простой способ запустить проект - использовать Docker Compose:

```bash
# Клонировать репозиторий
git clone https://github.com/krawwwwy/book-library-api.git
cd book-library-api

# Запустить все контейнеры
docker-compose up -d
```

После запуска приложение будет доступно по адресу [http://localhost:8080](http://localhost:8080)

### Запуск без Docker

Для запуска без Docker вам потребуется:

1. Go 1.21.3 или выше
2. PostgreSQL 15

```bash
# Настройка базы данных
psql -U postgres -c "CREATE DATABASE book_library"
psql -U postgres -d book_library -f scripts/create_tables.sql

# Запуск приложения
go run cmd/api/main.go
```

## API Endpoints

Все API доступны по базовому пути `/api`:

| Метод | Путь | Описание |
|-------|------|----------|
| GET | /api/books | Получение списка книг с пагинацией |
| GET | /api/books/:id | Получение книги по ID |
| POST | /api/books | Создание новой книги |
| PUT | /api/books/:id | Обновление книги |
| DELETE | /api/books/:id | Удаление книги |
| GET | /api/books/search | Поиск книг по запросу |
| POST | /api/books/:id/toggle-availability | Изменение доступности книги |

## Веб-интерфейс

Проект включает в себя удобный веб-интерфейс для работы с библиотекой:

- Главная страница: информация о проекте
- Каталог книг: просмотр, поиск, добавление, редактирование и удаление книг

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

## Автор

Krawwwwy - [GitHub](https://github.com/krawwwwy)

## Лицензия

MIT 