# API Documentation

This directory contains API documentation for the Book Library API.

## Endpoints

### Books API

#### GET /api/books
- Description: Get a list of books with pagination
- Parameters:
  - page: page number (default: 1)
  - page_size: number of items per page (default: 10)
- Response: Array of Book objects

#### GET /api/books/:id
- Description: Get a specific book by ID
- Parameters:
  - id: Book ID
- Response: Book object

#### POST /api/books
- Description: Create a new book
- Body: BookCreate object
- Response: Created Book object

#### PUT /api/books/:id
- Description: Update a book
- Parameters:
  - id: Book ID
- Body: BookCreate object
- Response: Updated Book object

#### DELETE /api/books/:id
- Description: Delete a book
- Parameters:
  - id: Book ID
- Response: No content

#### GET /api/books/search
- Description: Search for books by query
- Parameters:
  - q: Search query
- Response: Array of Book objects

#### POST /api/books/:id/toggle-availability
- Description: Toggle book availability
- Parameters:
  - id: Book ID
- Response: Updated Book object

## Models

### Book
```json
{
  "id": 1,
  "title": "War and Peace",
  "author": "Leo Tolstoy",
  "isbn": "9785171147440",
  "description": "Epic novel",
  "year": 1869,
  "publisher": "Publisher",
  "available": true,
  "created_at": "2025-05-15T21:00:00Z",
  "updated_at": "2025-05-15T21:00:00Z"
}
```

### BookCreate
```json
{
  "title": "War and Peace",
  "author": "Leo Tolstoy",
  "isbn": "9785171147440",
  "description": "Epic novel",
  "year": 1869,
  "publisher": "Publisher"
}
``` 