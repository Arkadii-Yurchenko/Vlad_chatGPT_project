package server

import (
	"Vlad_chatGPT_project/handlers" // Добавлен правильный импорт
	"log"
	"net/http"
)

// StartServer запускает HTTP-сервер
func StartServer() {
	// Подключаем обработчики
	http.HandleFunc("/ask", handlers.AskHandler)
	http.HandleFunc("/available-models", handlers.GetAvailableModelsHandler)
	http.HandleFunc("/set-model", handlers.SetModelHandler)
	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
