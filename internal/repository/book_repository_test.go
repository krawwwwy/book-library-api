package repository

import (
	"testing"
	"time"

	"github.com/krawwwwy/book-library-api/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BookRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo *BookRepository
}

func (s *BookRepositoryTestSuite) SetupSuite() {
	// Подключение к тестовой базе данных
	dsn := "host=localhost user=postgres password=postgres dbname=book_library_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db

	// Миграция схемы
	err = s.db.AutoMigrate(&model.Book{})
	if err != nil {
		s.T().Fatal(err)
	}

	s.repo = NewBookRepository(s.db)
}

func (s *BookRepositoryTestSuite) TearDownTest() {
	// Очистка таблицы после каждого теста
	s.db.Exec("TRUNCATE TABLE books")
}

func (s *BookRepositoryTestSuite) TestCreateBook() {
	book := &model.Book{
		Title:       "Тестовая книга",
		Author:      "Тестовый автор",
		ISBN:        "1234567890",
		Description: "Тестовое описание",
		Year:        2024,
		Publisher:   "Тестовое издательство",
		Available:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Act
	err := s.repo.Create(book)

	// Assert
	assert.NoError(s.T(), err)
	assert.NotZero(s.T(), book.ID)

	// Проверяем, что книга действительно создана в базе
	var found model.Book
	err = s.db.First(&found, book.ID).Error
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), book.Title, found.Title)
	assert.Equal(s.T(), book.Author, found.Author)
	assert.Equal(s.T(), book.ISBN, found.ISBN)
}

func (s *BookRepositoryTestSuite) TestGetByID() {
	// Arrange
	book := &model.Book{
		Title:     "Тестовая книга",
		Author:    "Тестовый автор",
		ISBN:      "1234567890",
		Available: true,
	}
	s.db.Create(book)

	// Act
	found, err := s.repo.GetByID(book.ID)

	// Assert
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), found)
	assert.Equal(s.T(), book.ID, found.ID)
	assert.Equal(s.T(), book.Title, found.Title)
}

func (s *BookRepositoryTestSuite) TestGetByIDNotFound() {
	// Act
	found, err := s.repo.GetByID(999)

	// Assert
	assert.Error(s.T(), err)
	assert.Nil(s.T(), found)
}

func (s *BookRepositoryTestSuite) TestGetAll() {
	// Arrange
	books := []model.Book{
		{Title: "Книга 1", Author: "Автор 1", ISBN: "1111111111"},
		{Title: "Книга 2", Author: "Автор 2", ISBN: "2222222222"},
		{Title: "Книга 3", Author: "Автор 3", ISBN: "3333333333"},
	}
	for _, book := range books {
		s.db.Create(&book)
	}

	// Act
	found, err := s.repo.GetAll(1, 2)

	// Assert
	assert.NoError(s.T(), err)
	assert.Len(s.T(), found, 2)
}

func (s *BookRepositoryTestSuite) TestSearch() {
	// Arrange
	books := []model.Book{
		{Title: "Война и мир", Author: "Лев Толстой", ISBN: "1111111111"},
		{Title: "Анна Каренина", Author: "Лев Толстой", ISBN: "2222222222"},
		{Title: "Преступление и наказание", Author: "Федор Достоевский", ISBN: "3333333333"},
	}
	for _, book := range books {
		s.db.Create(&book)
	}

	// Act
	found, err := s.repo.Search("Толстой")

	// Assert
	assert.NoError(s.T(), err)
	assert.Len(s.T(), found, 2)
}

func TestBookRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BookRepositoryTestSuite))
} 