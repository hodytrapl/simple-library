package main

import "fmt"

func main() {

	myLibrary := &Library{}

	fmt.Println("1. НАПОЛНЕНИЕ БИБЛИОТЕКИ")

	reader1 := myLibrary.AddReader("Агунда", "Кокойты")
	reader2 := myLibrary.AddReader("Сергей", "Меняйло")
	fmt.Printf("   Добавлен: %s\n", reader1)
	fmt.Printf("   Добавлен: %s\n", reader2)

	book1 := myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
	book2 := myLibrary.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)
	fmt.Printf("   Добавлена: %s\n", book1)
	fmt.Printf("   Добавлена: %s\n", book2)

	fmt.Printf("\n   Итого: %d читателей, %d книг\n", len(myLibrary.Readers), len(myLibrary.Books))

	fmt.Println("\n2. ТЕСТИРОВАНИЕ ВЫДАЧИ КНИГ")

	fmt.Println("\n   Сценарий 1: Успешная выдача книги")
	err := myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Printf("   ОШИБКА: %v\n", err)
	} else {
		fmt.Println("   УСПЕХ: Книга '1984' выдана читателю Агунда Кокойты")
		
		book, findErr := myLibrary.FindBookByID(1)
		if findErr != nil {
			fmt.Printf("   ОШИБКА поиска книги: %v\n", findErr)
		} else {
			fmt.Printf("   Статус книги: %s\n", book.String())
		}
	}

	fmt.Println("\n   Сценарий 2: Попытка выдать уже выданную книгу")
	err = myLibrary.IssueBookToReader(1, 2)
	if err != nil {
		fmt.Printf("   ОЖИДАЕМАЯ ОШИБКА: %v\n", err)
	} else {
		fmt.Println("   НЕОЖИДАННЫЙ УСПЕХ: Книга выдана")
	}

	fmt.Println("\n   Сценарий 3: Попытка выдать книгу несуществующему читателю")
	err = myLibrary.IssueBookToReader(2, 99)
	if err != nil {
		fmt.Printf("   ОЖИДАЕМАЯ ОШИБКА: %v\n", err)
	} else {
		fmt.Println("   НЕОЖИДАННЫЙ УСПЕХ: Книга выдана")
	}

	fmt.Println("\n   Сценарий 4: Попытка выдать несуществующую книгу")
	err = myLibrary.IssueBookToReader(99, 1)
	if err != nil {
		fmt.Printf("   ОЖИДАЕМАЯ ОШИБКА: %v\n", err)
	} else {
		fmt.Println("   НЕОЖИДАННЫЙ УСПЕХ: Книга выдана")
	}

	fmt.Println("\n3. ТЕСТИРОВАНИЕ ВОЗВРАТА КНИГ")

	fmt.Println("\n   Сценарий 5: Успешный возврат книги")
	err = myLibrary.ReturnBook(1)
	if err != nil {
		fmt.Printf("   ОШИБКА: %v\n", err)
	} else {
		fmt.Println("   УСПЕХ: Книга '1984' возвращена в библиотеку")
		
		book, findErr := myLibrary.FindBookByID(1)
		if findErr != nil {
			fmt.Printf("   ОШИБКА поиска книги: %v\n", findErr)
		} else {
			fmt.Printf("   Статус книги: %s\n", book.String())
		}
	}

	fmt.Println("\n   Сценарий 6: Попытка вернуть книгу, которая уже в библиотеке")
	err = myLibrary.ReturnBook(1)
	if err != nil {
		fmt.Printf("   ОЖИДАЕМАЯ ОШИБКА: %v\n", err)
	} else {
		fmt.Println("   НЕОЖИДАННЫЙ УСПЕХ: Книга возвращена")
	}

	fmt.Println("\n   Сценарий 7: Попытка вернуть несуществующую книгу")
	err = myLibrary.ReturnBook(99)
	if err != nil {
		fmt.Printf("   ОЖИДАЕМАЯ ОШИБКА: %v\n", err)
	} else {
		fmt.Println("   НЕОЖИДАННЫЙ УСПЕХ: Книга возвращена")
	}

	fmt.Println("\n4. ДОПОЛНИТЕЛЬНЫЕ СЦЕНАРИИ")

	fmt.Println("\n   Сценарий 8: Выдача книги неактивному читателю")
	
	deactivateErr := reader2.Deactive()
	if deactivateErr != nil {
		fmt.Printf("   ОШИБКА деактивации: %v\n", deactivateErr)
	} else {
		fmt.Println("   Читатель Сергей Меняйло деактивирован")
	}

	err = myLibrary.IssueBookToReader(2, 2)
	if err != nil {
		fmt.Printf("   ОЖИДАЕМАЯ ОШИБКА: %v\n", err)
	} else {
		fmt.Println("   НЕОЖИДАННЫЙ УСПЕХ: Книга выдана неактивному читателю")
	}

	fmt.Println("\n5. ФИНАЛЬНОЕ СОСТОЯНИЕ БИБЛИОТЕКИ")
	fmt.Printf("   Читатели: %d\n", len(myLibrary.Readers))
	fmt.Printf("   Книги: %d\n", len(myLibrary.Books))
	fmt.Println("\n   Список всех книг:")
	myLibrary.ListAllBooks()
}
