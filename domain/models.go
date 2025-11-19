package domain

import "fmt"

//Структура книги
type Book struct {
	ID       int
	Year     int
	Title    string
	Author   string
	IsIssued bool
	ReaderId *int
}

//Структура читателя
type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

//Метод для красивого выводы читателя
func (r Reader) String() string {
	status := ""
	if r.IsActive {
		status = "активен."
	} else {
		status = "неактивен."
	}
	return fmt.Sprintf("Пользователь %s %s, ID: %d, пользователь %s", r.FirstName, r.LastName, r.ID, status)
}

//Метод для красивого вывода книга
func (b Book) String() string {
	status := ""
	if b.IsIssued {
		status = "используется."
		return fmt.Sprintf("ID: %d, %s (%s %d), книга %s читателем %d", b.ID, b.Title, b.Author, b.Year, status, *b.ReaderId)
	} else {
		status = "не используется."
		return fmt.Sprintf("ID: %d, %s (%s %d), книга %s", b.ID, b.Title, b.Author, b.Year, status)
	}
}

//Метод для деактивации читателя
func (r *Reader) Deactivate() error{
	if !r.IsActive {
		return fmt.Errorf("%v и так неактивен", r)
	} else{
		r.IsActive = false
	}
	return nil
}

//Метод выдающий книгу читателю
func (b *Book) IssueBook(r *Reader) error {
	if b.IsIssued {
		return fmt.Errorf("книга уже используется")
	} else {
		b.IsIssued = true
		b.ReaderId = &r.ID
	}
	return nil
}

//Метод возвращающий книгу
func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	} else {
		b.IsIssued = false
		b.ReaderId = nil
	}
	return nil
}