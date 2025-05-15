package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krawwwwy/book-library-api/internal/api"
	"github.com/krawwwwy/book-library-api/internal/config"
	"github.com/krawwwwy/book-library-api/internal/model"
	"github.com/krawwwwy/book-library-api/internal/repository"
	"github.com/krawwwwy/book-library-api/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Загрузка конфигурации
	cfg := config.GetConfig()

	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(cfg.DB.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Автоматическая миграция моделей
	if err := db.AutoMigrate(&model.Book{}); err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}

	// Инициализация репозитория
	bookRepo := repository.NewBookRepository(db)

	// Инициализация сервиса
	bookService := service.NewBookService(bookRepo)

	// Инициализация обработчика
	bookHandler := api.NewBookHandler(bookService)

	// Инициализация роутера Gin
	router := gin.Default()

	// Обслуживание статических файлов
	router.Static("/css", "./public/css")
	router.Static("/js", "./public/js")
	router.StaticFile("/", "./public/index.html")
	router.StaticFile("/books.html", "./public/books.html")

	// Регистрация API маршрутов
	bookHandler.RegisterRoutes(router)

	// Настройка сервера
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	// Запуск сервера в горутине
	go func() {
		log.Printf("Сервер запущен на http://localhost:%s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска сервера: %s\n", err)
		}
	}()

	// Ожидание сигнала для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Выключение сервера...")

	// Контекст для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Ошибка при выключении сервера:", err)
	}

	log.Println("Сервер успешно остановлен")
} 