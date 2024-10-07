package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	token  = ""
	chatID = 0
)

var (
	API_SendMessage = "https://api.telegram.org/bot%s/sendMessage"
)

type TelegramMessage struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

func SendMessageToTG(msg string) (err error) {
	url := fmt.Sprintf(API_SendMessage, token)

	body := TelegramMessage{
		ChatID: chatID,
		Text:   msg,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("SendMessageToTG(): json marshl fail, err = %s", err.Error())
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Printf("SendMessageToTG(): http post fail, err = %s", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("SendMessageToTG(): http status fail, err = %s, status = %d", err.Error(), resp.StatusCode)
		return
	}
	return
}
