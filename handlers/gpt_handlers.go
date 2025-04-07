package handlers

import (
	"Vlad_chatGPT_project/api"
	"Vlad_chatGPT_project/config"
	"Vlad_chatGPT_project/models"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	currentModel    string
	availableModels []string
)

func init() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Ошибка загрузки конфигурации:", err)
		return
	}

	availableModels = cfg.AvailableModels
	currentModel = cfg.AvailableModels[0]
}

// AskHandler обрабатывает запросы к модели OpenAI

func AskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST метод поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var message models.MessageFromPostman
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Ошибка в JSON", http.StatusBadRequest)
		return
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		http.Error(w, "Ошибка загрузки конфигурации", http.StatusInternalServerError)
		return
	}

	// Отправляем запрос к OpenAI API
	response, err := api.CallOpenAI(cfg.APIKey, currentModel, cfg.APIURL, message)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка OpenAI: %v", err), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ с текущей моделью
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"response": response,
		"model":    currentModel, // Включаем текущую модель в ответ
	})
}

// GetAvailableModelsHandler возвращает список доступных моделей
func GetAvailableModelsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]string{"available_models": availableModels})
}

// SetModelHandler изменяет текущую модель
func SetModelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST метод поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var SetModelRequest models.SetModelRequest
	if err := json.NewDecoder(r.Body).Decode(&SetModelRequest); err != nil {
		http.Error(w, "Ошибка в JSON", http.StatusBadRequest)
		return
	}

	// Проверяем, существует ли выбранная модель в списке доступных
	for _, m := range availableModels {
		if m == SetModelRequest.Model {
			currentModel = m
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Модель успешно изменена"})
			return
		}
	}

	http.Error(w, "Модель не найдена в списке доступных", http.StatusBadRequest)
}
