package main

import "fmt"

type Library struct {
	Books   []*Book
	Readers []*Reader

	lastBookID   int
	lastReaderID int
}

func (lib *Library) AddReader(firstName, lastName string) *Reader {
	lib.lastReaderID++

	newReader := &Reader{
		ID:        &lib.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		isActive:  true,
	}

	lib.Readers = append(lib.Readers, newReader)
	return newReader
}
func (lib *Library) AddBook(title, author string, year int) *Book {
	lib.lastBookID++

	newBook := &Book{
		ID:      lib.lastBookID,
		Title:   title,
		Author:  author,
		Year:    year,
		isIssue: false,
	}

	lib.Books = append(lib.Books, newBook)

	return newBook
}

func (lib *Library) ListAllBooks() {
	for _, book := range lib.Books {
		fmt.Println(book.String())
	}
}

func (lib *Library) FindBookByID(id int) (*Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (lib *Library) FindReaderByID(id int) (*Reader, error) {
	for _, reader := range lib.Readers {
		if *reader.ID == id {
			return reader, nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найдена", id)
}

func (lib *Library) IssueBookToReader(bookID int, readerID int) error {
	book_result, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	reader_result, err := lib.FindReaderByID(readerID)

	if err != nil {
		return err
	}
	
	return book_result.IssueBook(reader_result)
}

func (lib *Library) ReturnBook(bookid int) error{
	book, err := lib.FindBookByID(bookid)
	if err != nil {
		return err
	}
	
	return book.ReturnBook()
}

type Reader struct {
	ID        *int
	FirstName string
	LastName  string
	isActive  bool
}



func (r *Reader) Deactive() error {
	if !r.isActive {
		return fmt.Errorf("читатель %s %s уже неактивен", r.FirstName, r.LastName)
	}
	r.isActive = false
	return nil
}

func (r Reader) String() string {
	status := "активен"
	if !r.isActive {
		status = "неактивен"
	}
	return fmt.Sprintf("Читатель: %s %s (ID: %d, статус: %s)", r.FirstName, r.LastName, r.ID, status)
}


func (r *Reader) AssignBook(b *Book) error {
	if !b.isIssue {
		return fmt.Errorf("Книга '%s' не выдана\n", b.Title)
	}
	if b.ReaderTakerID != *r.ID {
		
		return fmt.Errorf("Книга '%s' не выдана ему\n", b.Title)
	}
	return fmt.Errorf("Читатель %s %s взял книгу '%s' (%s, %d)\n", r.FirstName, r.LastName, b.Title, b.Author, b.Year)
}

type Book struct {
	ID            int
	Title         string
	Year          int
	Author        string
	isIssue       bool
	ReaderTakerID int
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.isIssue {
		status = fmt.Sprintf("выдана читателю ID: %d", b.ReaderTakerID)
	}
	return fmt.Sprintf("Книга: %s (%s, %d) - %s", b.Title, b.Author, b.Year, status)
}

func (b *Book) IssueBook(reader *Reader) error {
	if b.isIssue {	
		return fmt.Errorf("Книга %s уже используется\n", b.Title)
	}
	if !reader.isActive {
		return fmt.Errorf("Читатель %s %s не активен и не может получить книгу.", reader.FirstName, reader.LastName)
	}

	if reader.ID == nil {		
		return fmt.Errorf("У читателя %s %s не указан ID\n", reader.FirstName, reader.LastName)
	}

	b.isIssue = true
	b.ReaderTakerID = *reader.ID
	return fmt.Errorf("Книга %s была выдана читателю %s %s\n", b.Title, reader.FirstName, reader.LastName)
}

func (b *Book) ReturnBook() error {
	if !b.isIssue {		
		return fmt.Errorf("Книга %s и так в библиотеке", b.Title)
	}
	b.isIssue = false
	b.ReaderTakerID = 0
	return nil
}
