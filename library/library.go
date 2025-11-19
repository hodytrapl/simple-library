package library

import (
	"fmt"

	"github.com/hodytrapl/simple-library/domain"
)

// Структура библиотеки
type Library struct {
	Books   []*domain.Book
	Readers []*domain.Reader

	lastBookID   int
	lastReaderID int
}

func New() *Library {
	return &Library{Books: []*domain.Book{}, Readers: []*domain.Reader{}}
}

// Метод добавляющтй читателя в библиотеку
func (l *Library) AddReader(firstName, lastName string) (*domain.Reader, error) {
	if firstName == "" || lastName == "" {
		return nil, fmt.Errorf("имя или фамилия не могут быть пустыми")
	}
	l.lastReaderID++
	newReader := &domain.Reader{
		ID:        l.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}
	l.Readers = append(l.Readers, newReader)
	return newReader, nil
}

// Метод добавляющий книгу в библиотеку
func (l *Library) AddBook(year int, title, author string) (*domain.Book, error) {
	for _, value := range l.Books {
		if value.Author == author && value.Title == title {
			return nil, fmt.Errorf("такая книга уже есть в библиотеке")
		}
	}
	l.lastBookID++
	newBook := &domain.Book{
		ID:       l.lastBookID,
		Year:     year,
		Title:    title,
		Author:   author,
		IsIssued: false,
	}
	l.Books = append(l.Books, newBook)
	return newBook, nil
}

// Метод ищущий читателя по ID
func (l *Library) FindBookById(id int) (*domain.Book, error) {
	flag := false
	for i := 0; i < len(l.Books); i++ {
		if i == id-1 {
			flag = true
		}
	}
	if flag {
		return l.Books[id-1], nil
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

// Метод ищущий книгу по ID
func (l *Library) FindReaderById(id int) (*domain.Reader, error) {
	flag := false
	for i := 0; i < len(l.Readers); i++ {
		if i == id-1 {
			flag = true
		}
	}
	if flag {
		return l.Readers[id-1], nil
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

// Метод для выдачи книги читателю
func (l *Library) IssueBookToReader(bookId, readerId int) error {
	book, err := l.FindBookById(bookId)
	if book == nil {
		return err
	}
	reader, err := l.FindReaderById(readerId)
	if reader == nil {
		return err
	}
	err = book.IssueBook(reader)
	if err != nil {
		return err
	}
	return nil
}

// Метод возвращающий книгу по ID
func (l *Library) ReturnBook(bookId int) error {
	book, err := l.FindBookById(bookId)
	if err != nil {
		return err
	}
	test := book.ReturnBook()
	if test != nil {
		return test
	}
	return nil
}
