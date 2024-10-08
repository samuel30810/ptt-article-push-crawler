package main

import (
	"fmt"
	"time"

	PTTCrawler "headphone/Crawler/PTT"
	TG "headphone/TG"
	Tool "headphone/TooL"
)

func main() {

	url := "https://www.ptt.cc/bbs/Headphone/M.1530392323.A.695.html"

	for {
		SendPTTPushToTG(url)
		time.Sleep(15 * time.Minute)
	}
}

func SendPTTPushToTG(url string) (err error) {
	pushDataList, err := PTTCrawler.GetPTTPushData(url)
	if err != nil {
		fmt.Printf("Get PushData fail, err = %s", err.Error())
		return
	}

	msg := ""

	for _, push := range pushDataList {
		inTime, err := Tool.IsWithinOneHour(push.IPDatetime)
		if err != nil {
			continue
		}
		if inTime {
			fmt.Printf("Content: %s\nIPDatetime: %s\n", push.Content, push.IPDatetime)
			msg = msg + "\n" + push.Content
		}
	}

	if msg == "" {
		fmt.Println("no new push")
		return
	}

	msg = "結果：" + "\n" + msg

	err = TG.SendMessageToTG(msg)
	if err != nil {
		fmt.Printf("Send Message To TG fail, err = %s \n", err.Error())
		return
	}

	return
}
