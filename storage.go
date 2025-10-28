package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Storable interface {
	Save() error
	Load() error
}

func SaveBooksToCSV(books []*Book, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Title", "Author", "Year", "IsIssue", "ReaderTakerID"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("ошибка записи заголовка: %v", err)
	}

	for _, book := range books {
		record := []string{
			strconv.Itoa(book.ID),
			book.Title,
			book.Author,
			strconv.Itoa(book.Year),
			strconv.FormatBool(book.isIssue),
			strconv.Itoa(book.ReaderTakerID),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("ошибка записи данных книги: %v", err)
		}
	}

	return nil
}

func LoadBooksFromCSV(filename string) ([]*Book, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, fmt.Errorf("error:%v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, 0, fmt.Errorf("error:%v", err)
	}

	if len(records) < 2 {
		return nil, 0, fmt.Errorf("error:%v", "нихрена тута нету")
	}

	var books []*Book
	maxID := 0

	for i, record := range records[1:] {
		if len(record) != 6 {
			return nil, 0, fmt.Errorf("error:%d", i+2)
		}

		id, err := strconv.Atoi(record[0])

		if err != nil {
			return nil, 0, fmt.Errorf("ошибка парсинга ID в строке %d: %v", i+2, err)
		}

		year, err := strconv.Atoi(record[3])

		if err != nil {
			return nil, 0, fmt.Errorf("ошибка парсинга Year в строке %d: %v", i+2, err)
		}

		isIssue, err := strconv.ParseBool(record[4])

		if err != nil {
			return nil, 0, fmt.Errorf("ошибка парсинга Year в строке %d: %v", i+2, err)
		}

		readerTakerID, err := strconv.Atoi(record[5])
		if err != nil {
			return nil, 0, fmt.Errorf("ошибка парсинга ReaderTakerID в строке %d: %v", i+2, err)
		}

		if id > maxID {
			maxID = id
		}

		book := &Book{
			ID:            id,
			Title:         record[1],
			Author:        record[2],
			Year:          year,
			isIssue:       isIssue,
			ReaderTakerID: readerTakerID,
		}

		books = append(books, book)
	}

	return books, maxID, nil
}
