package api

import (
	"Vlad_chatGPT_project/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// CallOpenAI отправляет запрос к OpenAI API и возвращает ответ
func CallOpenAI(apiKey string, model string, apiURL string, userMessage models.MessageFromPostman) (string, error) {
	requestBody := models.RequestBody{
		Model: model,
		Messages: []models.Message{
			{Role: "user", Content: userMessage.Message},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ошибка API: %s", string(bodyBytes))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responseBody models.ResponseBody
	if err := json.Unmarshal(bodyBytes, &responseBody); err != nil {
		return "", err
	}

	if len(responseBody.Choices) > 0 {
		return responseBody.Choices[0].Message.Content, nil
	}

	return "", errors.New("не получено ни одного ответа")
}
