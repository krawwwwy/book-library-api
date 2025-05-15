package service

import (
	"errors"
	"testing"

	"github.com/krawwwwy/book-library-api/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBookRepository - мок для репозитория книг
type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) Create(book *model.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) GetByID(id uint) (*model.Book, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Book), args.Error(1)
}

func (m *MockBookRepository) GetAll(page, pageSize int) ([]model.Book, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]model.Book), args.Error(1)
}

func (m *MockBookRepository) Update(book *model.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBookRepository) GetByISBN(isbn string) (*model.Book, error) {
	args := m.Called(isbn)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Book), args.Error(1)
}

func (m *MockBookRepository) Search(query string) ([]model.Book, error) {
	args := m.Called(query)
	return args.Get(0).([]model.Book), args.Error(1)
}

func TestCreateBook(t *testing.T) {
	// Arrange
	mockRepo := new(MockBookRepository)
	service := NewBookService(mockRepo)
	
	testCases := []struct {
		name          string
		input         *model.BookCreate
		setupMock     func()
		expectedError bool
	}{
		{
			name: "Успешное создание книги",
			input: &model.BookCreate{
				Title:       "Война и мир",
				Author:      "Лев Толстой",
				ISBN:        "1234567890",
				Description: "Великий роман-эпопея",
				Year:        1869,
				Publisher:   "Русский вестник",
			},
			setupMock: func() {
				mockRepo.On("GetByISBN", "1234567890").Return(nil, errors.New("not found"))
				mockRepo.On("Create", mock.AnythingOfType("*model.Book")).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "Книга с таким ISBN уже существует",
			input: &model.BookCreate{
				Title:       "Дубликат",
				Author:      "Автор",
				ISBN:        "1234567890",
				Description: "Описание",
				Year:        2024,
				Publisher:   "Издательство",
			},
			setupMock: func() {
				existingBook := &model.Book{ID: 1, ISBN: "1234567890"}
				mockRepo.On("GetByISBN", "1234567890").Return(existingBook, nil)
			},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			tc.setupMock()

			// Act
			book, err := service.CreateBook(tc.input)

			// Assert
			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, book)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, book)
				assert.Equal(t, tc.input.Title, book.Title)
				assert.Equal(t, tc.input.Author, book.Author)
				assert.Equal(t, tc.input.ISBN, book.ISBN)
			}
		})
	}
}

func TestGetBookByID(t *testing.T) {
	// Arrange
	mockRepo := new(MockBookRepository)
	service := NewBookService(mockRepo)

	testCases := []struct {
		name          string
		bookID        uint
		setupMock     func()
		expectedError bool
	}{
		{
			name:   "Успешное получение книги",
			bookID: 1,
			setupMock: func() {
				book := &model.Book{
					ID:     1,
					Title:  "Тестовая книга",
					Author: "Тестовый автор",
				}
				mockRepo.On("GetByID", uint(1)).Return(book, nil)
			},
			expectedError: false,
		},
		{
			name:   "Книга не найдена",
			bookID: 999,
			setupMock: func() {
				mockRepo.On("GetByID", uint(999)).Return(nil, errors.New("not found"))
			},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			tc.setupMock()

			// Act
			book, err := service.GetBookByID(tc.bookID)

			// Assert
			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, book)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, book)
				assert.Equal(t, tc.bookID, book.ID)
			}
		})
	}
}

func TestSearchBooks(t *testing.T) {
	// Arrange
	mockRepo := new(MockBookRepository)
	service := NewBookService(mockRepo)

	testCases := []struct {
		name           string
		searchQuery    string
		setupMock      func()
		expectedCount  int
		expectedError  bool
	}{
		{
			name:        "Успешный поиск книг",
			searchQuery: "Толстой",
			setupMock: func() {
				books := []model.Book{
					{ID: 1, Title: "Война и мир", Author: "Лев Толстой"},
					{ID: 2, Title: "Анна Каренина", Author: "Лев Толстой"},
				}
				mockRepo.On("Search", "Толстой").Return(books, nil)
			},
			expectedCount: 2,
			expectedError: false,
		},
		{
			name:        "Поиск без результатов",
			searchQuery: "Несуществующий автор",
			setupMock: func() {
				mockRepo.On("Search", "Несуществующий автор").Return([]model.Book{}, nil)
			},
			expectedCount: 0,
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			tc.setupMock()

			// Act
			books, err := service.SearchBooks(tc.searchQuery)

			// Assert
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, books, tc.expectedCount)
			}
		})
	}
} 