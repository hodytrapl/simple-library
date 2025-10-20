package main

import "fmt"

type Reader struct {
	ID        *int
	FirstName string
	LastName  string
	isActive  bool
}

func (r Reader) DisplayReader() {
	fmt.Printf("читатель:%s %s (ID: %d)\n", r.FirstName, r.LastName, r.ID)
}

func (r *Reader) Deactive() {
	r.isActive = false
}

func (r Reader) String() string {
	status := ""
	if r.isActive {
		status = "active"
	} else {
		status = "noactive"
	}
	return fmt.Sprintf("пользователь %s %s, id: %d, %s\n", r.FirstName, r.LastName, r.ID, status)
}

func (r *Reader) AssignBook(b *Book) {
	if !b.isIssue {
		fmt.Printf("Книга '%s' не выдана\n", b.Title)
		return
	}
	if b.ReaderTakerID != *r.ID {
		fmt.Printf("Книга '%s' не выдана ему\n", b.Title)
		return
	}
	fmt.Printf("Читатель %s %s взял книгу '%s' (%s, %d)\n", r.FirstName, r.LastName, b.Title, b.Author, b.Year)
}

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

    fmt.Printf("Зарегистрирован новый читатель: %s\n", newReader)
    return newReader
}
func (lib *Library) AddBook(title, author string, year int) *Book {
    lib.lastBookID++

    newBook := &Book{
        ID:       lib.lastBookID,
        Title:    title,
        Author:   author,
        Year:     year,
        isIssue: false,
    }

    lib.Books = append(lib.Books, newBook)

    fmt.Printf("Добавлена новая книга: %s\n", newBook)
    return newBook
}

func (lib *Library) ListAllBooks(){
	for _, book :=range lib.Books{
		fmt.Println(book.String())
	}	
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
	return fmt.Sprintf("%s (%s, %d) \n", b.Title, b.Author, b.Year)
}
func (b *Book) IssueBook(reader *Reader) {
	if b.isIssue {
		fmt.Printf("Книга %s уже используется\n", b.Title)
		return
	}
	if !reader.isActive {
		fmt.Printf("Читатель %s %s не активен и не может получить книгу.", reader.FirstName, reader.LastName)
		return
	}

	if reader.ID == nil {
		fmt.Printf("У читателя %s %s не указан ID\n", reader.FirstName, reader.LastName)
		return
	}

	b.isIssue = true
	b.ReaderTakerID = *reader.ID
	fmt.Printf("Книга %s была выдана читателю %s %s\n", b.Title, reader.FirstName, reader.LastName)
}

func (b *Book) ReturnBook() {
	if !b.isIssue {
		fmt.Printf("Книга %s и так в библиотеке", b.Title)
		return
	}
	b.isIssue = false
	b.ReaderTakerID = 0
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}
