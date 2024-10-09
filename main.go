package main

import (
	"fmt"
	"time"

	Const "headphone/Const"
	PTTCrawler "headphone/Crawler/PTT"
	TG "headphone/TG"
)

var (
	urlMapLastPushTime = map[string]time.Time{}
)

func main() {

	err := Const.ReadAndSetSettings("setting.json")
	if err != nil {
		fmt.Printf("Const.ReadAndSetSettings fail, err = %s\n", err.Error())
		return
	}

	url := Const.ArticleURL

	SendAllPTTPushToTG(url)

	for {
		SendNewPTTPushToTG(url)
		time.Sleep(time.Duration(Const.CheckFrequency) * time.Minute)
	}
}

func SendAllPTTPushToTG(url string) (err error) {
	pushDataList, err := PTTCrawler.GetPTTPushData(url)
	if err != nil {
		fmt.Printf("SendAllPTTPushToTG(): Get PushData fail, err = %s", err.Error())
		return
	}

	msg := ""

	for _, push := range pushDataList {
		msg = msg + "\n" + push.Content + "_" + push.IPDatetime
	}

	if msg == "" {
		fmt.Println("SendAllPTTPushToTG(): no new push")
		return
	}

	msg = "結果：" + "\n" + msg

	err = TG.SendMessageToTG(msg)
	if err != nil {
		fmt.Printf("SendAllPTTPushToTG(): Send Message To TG fail, err = %s \n", err.Error())
		return
	}

	urlMapLastPushTime[url] = pushDataList[len(pushDataList)-1].PushTime
	return
}

func SendNewPTTPushToTG(url string) (err error) {
	pushDataList, err := PTTCrawler.GetPTTPushData(url)
	if err != nil {
		fmt.Printf("SendNewPTTPushToTG(): Get PushData fail, err = %s", err.Error())
		return
	}

	msg := ""

	for _, push := range pushDataList {

		if push.PushTime.After(urlMapLastPushTime[url]) {
			msg = msg + "\n" + push.Content + "_" + push.IPDatetime
		}
	}

	if msg == "" {
		fmt.Println("SendNewPTTPushToTG(): no new push")
		return
	}

	msg = "結果：" + "\n" + msg

	err = TG.SendMessageToTG(msg)
	if err != nil {
		fmt.Printf("SendNewPTTPushToTG(): Send Message To TG fail, err = %s \n", err.Error())
		return
	}

	urlMapLastPushTime[url] = pushDataList[len(pushDataList)-1].PushTime
	return
}
