package main

import (
	"fmt"
	"time"

	PTTCrawler "headphone/Crawler/PTT"
	TG "headphone/TG"
)

var (
	urlMapLastPushTime = map[string]time.Time{}
)

func main() {

	url := "https://www.ptt.cc/bbs/Headphone/M.1530392323.A.695.html"

	SendAllPTTPushToTG(url)

	for {
		SendNewPTTPushToTG(url)
		time.Sleep(15 * time.Minute)
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
