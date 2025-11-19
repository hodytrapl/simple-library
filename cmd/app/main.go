package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Регистрируем обработчик для пути "/"
	//Все запросы на адрес http://localhost/8080 будут обрабатываться функцией homeHandler
	http.HandleFunc("/", healthCheckHandler)

	//Определяем порт, на котором будет работать сервер
	port := ":8080"
	fmt.Printf("Запускаем сервер на порту %s", port)

	//Запускаем сервер
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err) //Если сервер не запустился - логируем и выходим
	}

}

// Наш первый обработчик
// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello from Library"))
// }

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	//1. Создаем структуру для ответа, чтобы ее было легко превратить в JSON
	response := struct {
		Status      string `json:"status"`
		ProjectName string `json:"project_name"`
	}{
		Status:      "available",
		ProjectName: "LibraryAPI",
	}

	//2. Сериаулизуем структуру в JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//3. Устанавливаем заголовок
	w.Header().Set("Content-Type", "application/json")

	//4. Устанавливаем статус код 200 ОК
	w.WriteHeader(http.StatusOK)

	//5. Пишем JSON для ответа
	w.Write(jsonData)
}
