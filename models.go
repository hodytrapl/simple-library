package main

import "fmt"

type Reader struct {
	ID        int
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

type Book struct {
	ID      int
	Title   string
	Year    int
	Author  string
	isIssue bool
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
	b.isIssue = true
	fmt.Printf("Книга %s была выдана\n", b.Title)
}

func (b *Book) ReturnBook() {
	if !b.isIssue {
		fmt.Printf("Книга %s и так в библиотеке", b.Title)
		return
	}
	b.isIssue = false
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}
