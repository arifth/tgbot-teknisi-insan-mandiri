package utils

import (
	"arifthalhah/sigesit-bot/v2/config"
	"arifthalhah/sigesit-bot/v2/templates"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func GetKeyValue(str string) (string, string) {
	keyValue := strings.Split(str, "=")
	return keyValue[0], keyValue[1]
}

type TelegramResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"`
	ErrorCode   int             `json:"error_code,omitempty"`
	Description string          `json:"description,omitempty"`
}

func RequestToChannel(chatID string, message string, messageID string) string {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.Config("TELEGRAM_API_TOKEN"))

	// Create the request body.
	data := url.Values{
		"chat_id":           {chatID},
		"text":              {message},
		"message_thread_id": {messageID},
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return fmt.Sprintf("Message sent successfully!")
	}

	// Set the Content-Type header.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create an HTTP client and send the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return fmt.Sprintf("Message sent successfully!")

	}
	defer resp.Body.Close()

	// Decode the response.
	var telegramResp TelegramResponse
	err = json.NewDecoder(resp.Body).Decode(&telegramResp)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return fmt.Sprintf("Message sent successfully!")

	}

	// Check the response status.
	if !telegramResp.Ok {
		fmt.Printf("Telegram API error: %d - %s\n", telegramResp.ErrorCode, telegramResp.Description)
		return fmt.Sprintf("Message sent successfully!", err)

	}

	return fmt.Sprintf("Message sent successfully!")
}

func IsMatchFormat(str string) (bool, []string, string) {
	keyValue := strings.Split(str, "\n")
	if len(keyValue) < 15 {
		return false, nil, ""
	}
	template := strings.Split(templates.RepliesToCreateNewTask(), "\n")
	if len(keyValue) == 15 {
		//	TODO: add validation for rest fields
		for key, value := range keyValue {
			if value == "" {
				fmt.Println("apa ini ", template[10], key)
				return false, nil, fmt.Sprintf("Field' %v 'tidak boleh kosong!", template[key])
			}
		}
	}
	return true, append([]string{}, keyValue[2:]...), ""
}
