package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/hodytrapl/simple-library/library"
	"github.com/hodytrapl/simple-library/storage"
)

func printMenu() {
	fmt.Println("\nДобро пожаловать в библиотеку.")
	fmt.Println("1 - Добавить книгу")
	fmt.Println("2 - Добавить читателя")
	fmt.Println("3 - Удалить книгу")
	fmt.Println("4 - Удалить читателя")
	fmt.Println("5 - Показать список книг")
	fmt.Println("6 - Показать список читателей")
	fmt.Println("7 - Найти книгу по ID")
	fmt.Println("8 - Найти читателя по ID")
	fmt.Println("9 - Выдать книгу читателю")
	fmt.Println("10 - Вернуть книгу в библиотеку")
	fmt.Println("11 - Сохранить список книг в CSV файле")
	fmt.Println("12 - Загрузить список книг из CSV файла")
	fmt.Println("13 - Сохранить список читателей в CSV файле")
	fmt.Println("14 - Загрузить список читателей из CSV файла")
	fmt.Println("15 - Выход.")
}

func hadlerChoise(userChoise int, myLibrary *library.Library, scanner *bufio.Scanner) {
	switch userChoise {
	//Добавить книгу
	case 1:
		fmt.Println("Введите год издания:")
		scanner.Scan()
		year, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Ошибка, нужно ввести число")
		}
		fmt.Println("Введите название книги:")
		scanner.Scan()
		title := scanner.Text()
		fmt.Println("Введите автора книги:")
		scanner.Scan()
		author := scanner.Text()
		_, err = myLibrary.AddBook(year, title, author)
		if err != nil {
			fmt.Println(err)
		}
	//Добавить читателя
	case 2:
		fmt.Println("Введите имя читателя:")
		scanner.Scan()
		firstName := scanner.Text()
		fmt.Println("Введите фамилию читателя:")
		scanner.Scan()
		lastName := scanner.Text()
		_, err := myLibrary.AddReader(firstName, lastName)
		if err != nil {
			fmt.Println(err)
		}
	//Удалить книгу
	case 3:
		fmt.Println("Книга удалена.(не сделано)")
	//Удалить чиатетля
	case 4:
		fmt.Println("Читатель удален.(не сделано)")
	//Показать список книг
	case 5:
		if len(myLibrary.Books) == 0 {
			fmt.Println("Список книг пуст.")
		} else {
			for _, book := range myLibrary.Books {
				fmt.Printf("\nГод издания: %d, Название: %s, Автор: %s\n", book.Year, book.Title, book.Author)
			}
		}
	//Показать список читателей
	case 6:
		if len(myLibrary.Readers) == 0 {
			fmt.Println("Список читателей пуст.")
		} else {
			for _, reader := range myLibrary.Readers {
				fmt.Printf("\nИмя: %s, Фамилия: %s\n", reader.FirstName, reader.LastName)
			}
		}
	//Найти книгу по ID
	case 7:
		fmt.Println("Введите ID книги:")
		scanner.Scan()
		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Ошибка, нужно ввести число.")
		}
		book, err := myLibrary.FindBookById(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(book)
	//Найти читателя по ID
	case 8:
		fmt.Println("Введите ID читателя:")
		scanner.Scan()
		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Ошибка, нужно ввести число.")
		}
		reader, err := myLibrary.FindReaderById(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(reader)
	//Выдать книгу читателю
	case 9:
		fmt.Println("Введите ID книги:")
		scanner.Scan()
		bookId, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Ошибка, нужно ввести число.")
		}
		fmt.Println("Введите ID читателя:")
		scanner.Scan()
		readerId, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Ошибка, нужно ввести число.")
		}
		err = myLibrary.IssueBookToReader(bookId, readerId)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Книга успешно выдана")
	//Вернуть книгу в библиотеку
	case 10:
		fmt.Println("Введите ID книги:")
		scanner.Scan()
		bookId, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Ошибка, нужно ввести число.")
		}
		err = myLibrary.ReturnBook(bookId)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Книга успешно возвращена.")
	//Сохранить список книг в CSC файле
	case 11:
		fmt.Println("Введите название файла CSV:")
		scanner.Scan()
		nameFile := scanner.Text()
		err := storage.SaveBooksToCSV(nameFile, myLibrary.Books)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Список книг успешно сохранен в файл '%s'.", nameFile)
	//Загрузить список книг из CSV файла
	case 12:
		fmt.Println("Введите название файла CSV:")
		scanner.Scan()
		nameFile := scanner.Text()
		var err error
		myLibrary.Books, err = storage.LoadBooksFromCSV(nameFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Список книг успешно выгружен из файла '%s'.", nameFile)
	//Сохранить список читателей в CSV файл
	case 13:
		fmt.Println("Введите название файла CSV:")
		scanner.Scan()
		nameFile := scanner.Text()
		err := storage.SaveReadersToCSV(nameFile, myLibrary.Readers)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Список читателей успешно сохранен в файл '%s'.", nameFile)
	//Загрузить список книг из CSV файла
	case 14:
		fmt.Println("Введите название файла CSV:")
		scanner.Scan()
		nameFile := scanner.Text()
		var err error
		myLibrary.Readers, err = storage.LoadReadersFromCSV(nameFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Список читателей успешно выгружен из файла '%s'.", nameFile)
	case 15:
		fmt.Println("До свидания!")
	}
}

func Run(myLibrary *library.Library) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		printMenu()
		scanner.Scan()
		userChoise, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Ошибка, нужно ввести число.")
		}
		hadlerChoise(userChoise, myLibrary, scanner)
		if userChoise == 15 {
			err := storage.SaveBooksToCSV("books.csv", myLibrary.Books)
			if err != nil {
				fmt.Println("Произошла ошибка сохранения книг:", err)
			}
			err = storage.SaveReadersToCSV("readers.csv", myLibrary.Readers)
			if err != nil {
				fmt.Println("Произошла ошибка сохранения пользователей:", err)
			}
			break
		}
	}

}
