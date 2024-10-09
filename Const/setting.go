package Const

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Settings struct {
	ArticleURL     string `json:"articleURL"`
	TelegramToken  string `json:"telegramToken"`
	TelegramChatID int    `json:"telegramChatID"`
	CheckFrequency int    `json:"checkFrequency"`
}

var (
	ArticleURL     = ""
	TelegramToken  = ""
	TelegramChatID = 0
	CheckFrequency = 0
)

func ReadAndSetSettings(filename string) (err error) {
	// 打開檔案
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("ReadAndSetSettings(): os.Open fail, err = %s\n", err.Error())
		return
	}
	defer file.Close()

	// 讀取檔案內容
	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("ReadAndSetSettings(): io.ReadAll fail, err = %s\n", err.Error())
		return
	}

	// 將 JSON 資料解析到結構體
	var settings Settings
	err = json.Unmarshal(bytes, &settings)
	if err != nil {
		fmt.Printf("ReadAndSetSettings(): json.Unmarshal fail, err = %s\n", err.Error())
		return
	}

	if settings.ArticleURL == "" || settings.TelegramToken == "" || settings.TelegramChatID == 0 || CheckFrequency == 0 {
		err = errors.New("setting broken, please check setting file")
		fmt.Printf("ReadAndSetSettings(): setting broken, please check setting file\n")
		return
	}

	ArticleURL = settings.ArticleURL
	TelegramToken = settings.TelegramToken
	TelegramChatID = settings.TelegramChatID
	CheckFrequency = settings.CheckFrequency
	return
}
