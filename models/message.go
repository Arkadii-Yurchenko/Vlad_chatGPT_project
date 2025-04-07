package models

// Message описывает сообщение для API
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// RequestBody описывает тело запроса к API
type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// ResponseChoice описывает вариант ответа от API
type ResponseChoice struct {
	Message Message `json:"message"`
}

// ResponseBody описывает тело ответа от API
type ResponseBody struct {
	Choices []ResponseChoice `json:"choices"`
}

type MessageFromPostman struct {
	Message string `json:"message"`
}

type SetModelRequest struct {
	Model string `json:"model"`
}
