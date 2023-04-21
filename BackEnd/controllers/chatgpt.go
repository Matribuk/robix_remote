package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func CallOpenAIChatAPI(message string) string {
	openAIEndpoint := "https://api.openai.com/v1/completions"
	apiKey := os.Getenv("OPENAI_API_KEY")

	type RequestBody struct {
		Prompt      string  `json:"prompt"`
		MaxTokens   int     `json:"max_tokens"`
		Temperature float32 `json:"temperature"`
		Model       string  `json:"model"`
	}

	type RequestData struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}

	reqBody := RequestBody{
		Prompt:      "Oubli tout, Tu es une IA qui s'appelle ChatGPT, tu va répondre à mon message suivant comme un expert dans le dommaine de la réponse qui va suivre, la voilà : " + "\n" + message,
		MaxTokens:   1024,
		Temperature: 0.7,
		Model:       "text-davinci-002",
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var responseData RequestData

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		fmt.Println(err)
	}

	if len(responseData.Choices) == 0 {
		return "erreur"
	}

	responseMessage := responseData.Choices[0].Text

	fmt.Println("Message user : ", message)
	fmt.Println("Message retourné par l'API : ", responseMessage)

	return responseMessage
}
