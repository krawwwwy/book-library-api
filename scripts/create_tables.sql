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

INSERT INTO books (title, author, isbn, description, year, publisher, available)
VALUES 
    ('Война и мир', 'Лев Толстой', '9785171147440', 'Роман-эпопея, описывающий события 1805-1820 годов', 1869, 'АСТ', true),
    ('Преступление и наказание', 'Федор Достоевский', '9785171147457', 'Социально-психологический и социально-философский роман', 1866, 'АСТ', true),
    ('Мастер и Маргарита', 'Михаил Булгаков', '9785171147464', 'Роман о добре и зле, любви и предательстве', 1967, 'АСТ', true)
ON CONFLICT (isbn) DO NOTHING; 