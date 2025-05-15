// Конфигурация
const API_URL = '/api';
let currentPage = 1;
const pageSize = 10;

// DOM элементы
document.addEventListener('DOMContentLoaded', function() {
    // Инициализация только на странице каталога книг
    if (window.location.pathname.includes('books.html')) {
        initBooksCatalog();
    }
});

// Инициализация каталога книг
function initBooksCatalog() {
    loadBooks();

    // Поиск книг
    document.getElementById('search-button').addEventListener('click', searchBooks);
    document.getElementById('search-input').addEventListener('keypress', function(e) {
        if (e.key === 'Enter') {
            searchBooks();
        }
    });

    // Пагинация
    document.getElementById('prev-page-btn').addEventListener('click', function() {
        if (currentPage > 1) {
            currentPage--;
            loadBooks();
        }
    });
    
    document.getElementById('next-page-btn').addEventListener('click', function() {
        currentPage++;
        loadBooks();
    });

    // Добавление новой книги
    document.getElementById('save-book-btn').addEventListener('click', saveBook);

    // Обновление книги
    document.getElementById('update-book-btn').addEventListener('click', updateBook);
}

// Загрузка списка книг
function loadBooks() {
    fetch(`${API_URL}/books?page=${currentPage}&page_size=${pageSize}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка при загрузке книг');
            }
            return response.json();
        })
        .then(books => {
            displayBooks(books);
            document.getElementById('current-page').textContent = currentPage;
            
            // Управление кнопками пагинации
            document.getElementById('prev-page-btn').disabled = currentPage === 1;
            document.getElementById('next-page-btn').disabled = books.length < pageSize;
        })
        .catch(error => {
            showMessage(error.message, 'danger');
        });
}

// Отображение книг в таблице
function displayBooks(books) {
    const tableBody = document.getElementById('books-table');
    tableBody.innerHTML = '';
    
    if (books.length === 0) {
        const row = document.createElement('tr');
        row.innerHTML = '<td colspan="8" class="text-center">Книги не найдены</td>';
        tableBody.appendChild(row);
        return;
    }

    books.forEach(book => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${book.id}</td>
            <td>${book.title}</td>
            <td>${book.author}</td>
            <td>${book.isbn}</td>
            <td>${book.year}</td>
            <td>${book.publisher || '-'}</td>
            <td><span class="${book.available ? 'book-available' : 'book-unavailable'}">${book.available ? 'Доступна' : 'Недоступна'}</span></td>
            <td class="action-buttons">
                <button class="btn btn-sm btn-outline-primary edit-book" data-id="${book.id}">Редактировать</button>
                <button class="btn btn-sm btn-outline-warning toggle-availability" data-id="${book.id}">${book.available ? 'Сделать недоступной' : 'Сделать доступной'}</button>
                <button class="btn btn-sm btn-outline-danger delete-book" data-id="${book.id}">Удалить</button>
            </td>
        `;
        tableBody.appendChild(row);
    });

    // Добавляем обработчики событий для кнопок действий
    document.querySelectorAll('.edit-book').forEach(button => {
        button.addEventListener('click', function() {
            const bookId = this.getAttribute('data-id');
            editBook(bookId);
        });
    });

    document.querySelectorAll('.toggle-availability').forEach(button => {
        button.addEventListener('click', function() {
            const bookId = this.getAttribute('data-id');
            toggleAvailability(bookId);
        });
    });

    document.querySelectorAll('.delete-book').forEach(button => {
        button.addEventListener('click', function() {
            const bookId = this.getAttribute('data-id');
            deleteBook(bookId);
        });
    });
}

// Поиск книг
function searchBooks() {
    const query = document.getElementById('search-input').value.trim();
    
    if (!query) {
        currentPage = 1;
        loadBooks();
        return;
    }
    
    fetch(`${API_URL}/books/search?q=${encodeURIComponent(query)}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка при поиске книг');
            }
            return response.json();
        })
        .then(books => {
            displayBooks(books);
            // Скрываем элементы пагинации при поиске
            document.getElementById('prev-page-btn').disabled = true;
            document.getElementById('next-page-btn').disabled = true;
            document.getElementById('current-page').textContent = 1;
        })
        .catch(error => {
            showMessage(error.message, 'danger');
        });
}

// Сохранение новой книги
function saveBook() {
    const bookData = {
        title: document.getElementById('title').value,
        author: document.getElementById('author').value,
        isbn: document.getElementById('isbn').value,
        description: document.getElementById('description').value,
        year: parseInt(document.getElementById('year').value),
        publisher: document.getElementById('publisher').value
    };
    
    fetch(`${API_URL}/books`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(bookData)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Ошибка при добавлении книги');
        }
        return response.json();
    })
    .then(newBook => {
        // Закрываем модальное окно и сбрасываем форму
        const modal = bootstrap.Modal.getInstance(document.getElementById('addBookModal'));
        modal.hide();
        document.getElementById('add-book-form').reset();
        
        // Обновляем список книг и показываем сообщение
        loadBooks();
        showMessage('Книга успешно добавлена', 'success');
    })
    .catch(error => {
        showMessage(error.message, 'danger');
    });
}

// Редактирование книги
function editBook(bookId) {
    // Загрузка данных книги
    fetch(`${API_URL}/books/${bookId}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка при загрузке данных книги');
            }
            return response.json();
        })
        .then(book => {
            // Заполняем форму данными
            document.getElementById('edit-book-id').value = book.id;
            document.getElementById('edit-title').value = book.title;
            document.getElementById('edit-author').value = book.author;
            document.getElementById('edit-isbn').value = book.isbn;
            document.getElementById('edit-description').value = book.description || '';
            document.getElementById('edit-year').value = book.year;
            document.getElementById('edit-publisher').value = book.publisher || '';
            
            // Открываем модальное окно
            const modal = new bootstrap.Modal(document.getElementById('editBookModal'));
            modal.show();
        })
        .catch(error => {
            showMessage(error.message, 'danger');
        });
}

// Обновление книги
function updateBook() {
    const bookId = document.getElementById('edit-book-id').value;
    const bookData = {
        title: document.getElementById('edit-title').value,
        author: document.getElementById('edit-author').value,
        isbn: document.getElementById('edit-isbn').value,
        description: document.getElementById('edit-description').value,
        year: parseInt(document.getElementById('edit-year').value),
        publisher: document.getElementById('edit-publisher').value
    };
    
    fetch(`${API_URL}/books/${bookId}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(bookData)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Ошибка при обновлении книги');
        }
        return response.json();
    })
    .then(updatedBook => {
        // Закрываем модальное окно
        const modal = bootstrap.Modal.getInstance(document.getElementById('editBookModal'));
        modal.hide();
        
        // Обновляем список книг и показываем сообщение
        loadBooks();
        showMessage('Книга успешно обновлена', 'success');
    })
    .catch(error => {
        showMessage(error.message, 'danger');
    });
}

// Изменение доступности книги
function toggleAvailability(bookId) {
    fetch(`${API_URL}/books/${bookId}/toggle-availability`, {
        method: 'POST'
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Ошибка при изменении доступности книги');
        }
        return response.json();
    })
    .then(updatedBook => {
        // Обновляем список книг и показываем сообщение
        loadBooks();
        const status = updatedBook.available ? 'доступной' : 'недоступной';
        showMessage(`Книга отмечена как ${status}`, 'success');
    })
    .catch(error => {
        showMessage(error.message, 'danger');
    });
}

// Удаление книги
function deleteBook(bookId) {
    if (confirm('Вы уверены, что хотите удалить эту книгу?')) {
        fetch(`${API_URL}/books/${bookId}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка при удалении книги');
            }
            // Обновляем список книг и показываем сообщение
            loadBooks();
            showMessage('Книга успешно удалена', 'success');
        })
        .catch(error => {
            showMessage(error.message, 'danger');
        });
    }
}

// Показать сообщение
function showMessage(message, type) {
    const messageBox = document.getElementById('message-box');
    messageBox.textContent = message;
    messageBox.className = `alert alert-${type}`;
    messageBox.classList.remove('d-none');
    
    // Автоматически скрываем сообщение через 3 секунды
    setTimeout(() => {
        messageBox.classList.add('fade-out');
        setTimeout(() => {
            messageBox.classList.add('d-none');
            messageBox.classList.remove('fade-out');
        }, 500);
    }, 3000);
} 