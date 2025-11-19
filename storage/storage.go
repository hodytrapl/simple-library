package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"encoding/json"
	"github.com/hodytrapl/simple-library/domain"
	"github.com/hodytrapl/simple-library/library"
)

type Storable interface {
	Save() error
	Load() error
}

type storageData struct {
	Books   []*domain.Book   `json:"books"`
	Readers []*domain.Reader `json:"readers"`
}


func SaveBooksToCSV(fileName string, books []*domain.Book) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("не удалось создать файл %s: %w", fileName, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"ID", "Название", "Автор", "Год", "Используется", "ID читателя"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("не удалось записать заголовок: %w", err)
	}

	for _, book := range books {
		if book.ReaderId != nil {
			record := []string{
				strconv.Itoa(book.ID),
				book.Title,
				book.Author,
				strconv.Itoa(book.Year),
				strconv.FormatBool(book.IsIssued),
				strconv.Itoa(*book.ReaderId),
			}
			if err := writer.Write(record); err != nil {
				return fmt.Errorf("не удалось записать данные книг: %w", err)
			}
		} else {
			record := []string{
				strconv.Itoa(book.ID),
				book.Title,
				book.Author,
				strconv.Itoa(book.Year),
				strconv.FormatBool(book.IsIssued),
				"nil",
			}
			if err := writer.Write(record); err != nil {
				return fmt.Errorf("не удалось записать данные книг: %w", err)
			}
		}
	}
	return nil
}

func LoadBooksFromCSV(fileName string) ([]*domain.Book, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var books []*domain.Book

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, record := range records[1:] {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("не удалось сконвертировать ID %s", record[0])
		}

		year, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, fmt.Errorf("не удалось сконвертировать год %s", record[3])
		}

		issue, err := strconv.ParseBool(record[4])
		if err != nil {
			return nil, fmt.Errorf("не удалось сконвертировать использование книги %s", record[4])
		}

		if record[5] != "nil"{
			readerid, err := strconv.Atoi(record[5])
			if err != nil {
				return nil, fmt.Errorf("не удалось сконввертировать ID читателя %s", record[5])
			}
			book := &domain.Book{
				ID:       id,
				Title:    record[1],
				Author:   record[2],
				Year:     year,
				IsIssued: issue,
				ReaderId: &readerid,
			}
			books = append(books, book)
		} else {
			book := &domain.Book{
				ID:       id,
				Title:    record[1],
				Author:   record[2],
				Year:     year,
				IsIssued: issue,
			}
			books = append(books, book)
		}
	}
	return books, nil
}

func SaveReadersToCSV(fileName string, readers []*domain.Reader) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("не удалось создать файл %s: %w", fileName, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"ID", "Имя", "Фамилия", "Активен"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("не удалось записать заголовок: %w", err)
	}

	for _, reader := range readers {
		record := []string{
			strconv.Itoa(reader.ID),
			reader.FirstName,
			reader.LastName,
			strconv.FormatBool(reader.IsActive),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("не удалось записать данные пользователя: %w", err)
		}
	}
	return nil
}

func LoadReadersFromCSV(fileName string) ([]*domain.Reader, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var readers []*domain.Reader

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, record := range records[1:] {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("не удалось сконвертировать ID %s", record[0])
		}
		isactive, err := strconv.ParseBool(record[3])
		if err != nil {
			return nil, fmt.Errorf("не удалось сконвертировать активность читателя %s", record[4])
		}
		reader1 := &domain.Reader{
			ID:       id,
			FirstName: record[1],
			LastName: record[2],
			IsActive: isactive,
			
		}
		readers = append(readers, reader1)
	}
	return readers, nil
}

func SaveLibraryToJSON(filepath string, lib *library.Library) error{
	data := storageData{
		Books: lib.Books,
		Readers: lib.Readers,
	}

	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, jsonData, 0644) 
}

func LoadLibraryFromJSON(filePath string) (*library.Library, error){
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data storageData

	if err := json.Unmarshal(jsonData, &data); err != nil{
		return nil, err
	}

	lib := library.New()
	lib.Books = data.Books
	lib.Readers = data.Readers

	return lib, nil
}