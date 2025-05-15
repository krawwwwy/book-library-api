package service

import (
	"errors"

	"github.com/krawwwwy/book-library-api/internal/model"
	"github.com/krawwwwy/book-library-api/internal/repository"
)

// BookService представляет сервис для работы с книгами
type BookService struct {
	repo *repository.BookRepository
}

// NewBookService создает новый экземпляр BookService
func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

// CreateBook создает новую книгу
func (s *BookService) CreateBook(bookCreate *model.BookCreate) (*model.Book, error) {
	// Проверяем, существует ли книга с таким ISBN
	existingBook, err := s.repo.GetByISBN(bookCreate.ISBN)
	if err == nil && existingBook != nil {
		return nil, errors.New("книга с таким ISBN уже существует")
	}

	// Создаем новую книгу
	book := &model.Book{
		Title:       bookCreate.Title,
		Author:      bookCreate.Author,
		ISBN:        bookCreate.ISBN,
		Description: bookCreate.Description,
		Year:        bookCreate.Year,
		Publisher:   bookCreate.Publisher,
		Available:   true,
	}

	if err := s.repo.Create(book); err != nil {
		return nil, err
	}

	return book, nil
}

// GetBookByID получает книгу по ID
func (s *BookService) GetBookByID(id uint) (*model.Book, error) {
	return s.repo.GetByID(id)
}

// GetAllBooks получает список всех книг с пагинацией
func (s *BookService) GetAllBooks(page, pageSize int) ([]model.Book, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return s.repo.GetAll(page, pageSize)
}

// UpdateBook обновляет информацию о книге
func (s *BookService) UpdateBook(id uint, bookUpdate *model.BookCreate) (*model.Book, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Проверяем, не пытаемся ли мы обновить ISBN на уже существующий
	if book.ISBN != bookUpdate.ISBN {
		existingBook, err := s.repo.GetByISBN(bookUpdate.ISBN)
		if err == nil && existingBook != nil && existingBook.ID != id {
			return nil, errors.New("книга с таким ISBN уже существует")
		}
	}

	book.Title = bookUpdate.Title
	book.Author = bookUpdate.Author
	book.ISBN = bookUpdate.ISBN
	book.Description = bookUpdate.Description
	book.Year = bookUpdate.Year
	book.Publisher = bookUpdate.Publisher

	if err := s.repo.Update(book); err != nil {
		return nil, err
	}

	return book, nil
}

// DeleteBook удаляет книгу
func (s *BookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}

// SearchBooks ищет книги по названию или автору
func (s *BookService) SearchBooks(query string) ([]model.Book, error) {
	return s.repo.Search(query)
}

// ToggleBookAvailability изменяет статус доступности книги
func (s *BookService) ToggleBookAvailability(id uint) (*model.Book, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	book.Available = !book.Available
	if err := s.repo.Update(book); err != nil {
		return nil, err
	}

	return book, nil
} 