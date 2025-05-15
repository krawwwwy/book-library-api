package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krawwwwy/book-library-api/internal/model"
	"github.com/krawwwwy/book-library-api/internal/service"
)

// BookHandler представляет обработчик HTTP-запросов для книг
type BookHandler struct {
	service *service.BookService
}

// NewBookHandler создает новый экземпляр BookHandler
func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// RegisterRoutes регистрирует маршруты для книг
// @Summary Регистрация маршрутов API для книг
// @Description Регистрирует все доступные эндпоинты для работы с книгами
func (h *BookHandler) RegisterRoutes(router *gin.Engine) {
	books := router.Group("/api/books")
	{
		books.POST("", h.CreateBook)
		books.GET("", h.GetBooks)
		books.GET("/:id", h.GetBook)
		books.PUT("/:id", h.UpdateBook)
		books.DELETE("/:id", h.DeleteBook)
		books.GET("/search", h.SearchBooks)
		books.POST("/:id/toggle-availability", h.ToggleAvailability)
	}
}

// CreateBook создает новую книгу
// @Summary Создание новой книги
// @Description Создает новую книгу в библиотеке
// @Tags books
// @Accept json
// @Produce json
// @Param book body model.BookCreate true "Данные новой книги"
// @Success 201 {object} model.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var bookCreate model.BookCreate
	if err := c.ShouldBindJSON(&bookCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат данных"})
		return
	}

	book, err := h.service.CreateBook(&bookCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// GetBooks получает список всех книг
// @Summary Получение списка книг
// @Description Получает список всех книг с пагинацией
// @Tags books
// @Produce json
// @Param page query int false "Номер страницы" default(1)
// @Param page_size query int false "Размер страницы" default(10)
// @Success 200 {array} model.Book
// @Failure 500 {object} map[string]string
// @Router /api/books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	books, err := h.service.GetAllBooks(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetBook получает книгу по ID
// @Summary Получение книги по ID
// @Description Получает детальную информацию о книге по её ID
// @Tags books
// @Produce json
// @Param id path int true "ID книги"
// @Success 200 {object} model.Book
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID"})
		return
	}

	book, err := h.service.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "книга не найдена"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook обновляет информацию о книге
// @Summary Обновление книги
// @Description Обновляет информацию о существующей книге
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "ID книги"
// @Param book body model.BookCreate true "Обновленные данные книги"
// @Success 200 {object} model.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID"})
		return
	}

	var bookUpdate model.BookCreate
	if err := c.ShouldBindJSON(&bookUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат данных"})
		return
	}

	book, err := h.service.UpdateBook(uint(id), &bookUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook удаляет книгу
// @Summary Удаление книги
// @Description Удаляет книгу из библиотеки
// @Tags books
// @Produce json
// @Param id path int true "ID книги"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID"})
		return
	}

	if err := h.service.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// SearchBooks ищет книги по запросу
// @Summary Поиск книг
// @Description Ищет книги по названию или автору
// @Tags books
// @Produce json
// @Param q query string true "Поисковый запрос"
// @Success 200 {array} model.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/books/search [get]
func (h *BookHandler) SearchBooks(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "параметр поиска не указан"})
		return
	}

	books, err := h.service.SearchBooks(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// ToggleAvailability изменяет статус доступности книги
// @Summary Изменение доступности книги
// @Description Переключает статус доступности книги (доступна/недоступна)
// @Tags books
// @Produce json
// @Param id path int true "ID книги"
// @Success 200 {object} model.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/books/{id}/toggle-availability [post]
func (h *BookHandler) ToggleAvailability(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID"})
		return
	}

	book, err := h.service.ToggleBookAvailability(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
} 