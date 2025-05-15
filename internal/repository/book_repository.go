package repository

import (
	"github.com/krawwwwy/book-library-api/internal/model"
	"gorm.io/gorm"
)

// BookRepository представляет репозиторий для работы с книгами
type BookRepository struct {
	db *gorm.DB
}

// NewBookRepository создает новый экземпляр BookRepository
func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

// Create создает новую книгу
func (r *BookRepository) Create(book *model.Book) error {
	return r.db.Create(book).Error
}

// GetByID получает книгу по ID
func (r *BookRepository) GetByID(id uint) (*model.Book, error) {
	var book model.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// GetAll получает все книги с пагинацией
func (r *BookRepository) GetAll(page, pageSize int) ([]model.Book, error) {
	var books []model.Book
	offset := (page - 1) * pageSize
	err := r.db.Offset(offset).Limit(pageSize).Find(&books).Error
	return books, err
}

// Update обновляет информацию о книге
func (r *BookRepository) Update(book *model.Book) error {
	return r.db.Save(book).Error
}

// Delete удаляет книгу по ID
func (r *BookRepository) Delete(id uint) error {
	return r.db.Delete(&model.Book{}, id).Error
}

// GetByISBN получает книгу по ISBN
func (r *BookRepository) GetByISBN(isbn string) (*model.Book, error) {
	var book model.Book
	err := r.db.Where("isbn = ?", isbn).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// Search ищет книги по названию или автору
func (r *BookRepository) Search(query string) ([]model.Book, error) {
	var books []model.Book
	err := r.db.Where("title ILIKE ? OR author ILIKE ?", "%"+query+"%", "%"+query+"%").Find(&books).Error
	return books, err
} 