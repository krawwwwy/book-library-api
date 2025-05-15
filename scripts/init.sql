-- Создание расширения для поиска
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Создание таблицы books
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    isbn VARCHAR(13) UNIQUE NOT NULL,
    description TEXT,
    year INTEGER,
    publisher VARCHAR(255),
    available BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание индексов для поиска
CREATE INDEX IF NOT EXISTS idx_books_title_trgm ON books USING gin (title gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_books_author_trgm ON books USING gin (author gin_trgm_ops);

-- Создание индекса для ISBN
CREATE UNIQUE INDEX IF NOT EXISTS idx_books_isbn ON books (isbn);

-- Добавление тестовых данных
INSERT INTO books (title, author, isbn, description, year, publisher, available, created_at, updated_at)
VALUES 
    ('Война и мир', 'Лев Толстой', '9785171147440', 'Роман-эпопея, описывающий события 1805-1820 годов', 1869, 'АСТ', true, NOW(), NOW()),
    ('Преступление и наказание', 'Федор Достоевский', '9785171147457', 'Социально-психологический и социально-философский роман', 1866, 'АСТ', true, NOW(), NOW()),
    ('Мастер и Маргарита', 'Михаил Булгаков', '9785171147464', 'Роман о добре и зле, любви и предательстве', 1967, 'АСТ', true, NOW(), NOW())
ON CONFLICT (isbn) DO NOTHING; 