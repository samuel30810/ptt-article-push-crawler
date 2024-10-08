package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	Const "headphone/Const"
	"io"
	"net/http"
)

var (
	API_SendMessage = "https://api.telegram.org/bot%s/sendMessage"
)

type TelegramMessage struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

func SendMessageToTG(msg string) (err error) {
	url := fmt.Sprintf(API_SendMessage, Const.TelegramToken)

	body := TelegramMessage{
		ChatID: Const.TelegramChatID,
		Text:   msg,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("SendMessageToTG(): json marshl fail, err = %s \n", err.Error())
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Printf("SendMessageToTG(): http post fail, err = %s \n", err.Error())
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("SendMessageToTG(): io.ReadAll fail, err = %s \n", err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("SendMessageToTG(): http status fail, resp body = %s, status = %d", string(respBody), resp.StatusCode)
		return
	}
	return
}
