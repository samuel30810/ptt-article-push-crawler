package main

import (
	"fmt"
	"time"

	Const "headphone/Const"
	PTTCrawler "headphone/Crawler/PTT"
	TG "headphone/TG"
	Tool "headphone/Tool"
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

	articles, err := Const.ReadArticleFile()
	if err != nil {
		fmt.Printf("Const.ReadArticleFile fail, err = %s\n", err.Error())
		return
	}

	for _, article := range articles {
		SendAllPTTPushToTG(article)
	}

	for {

		for _, article := range articles {
			SendNewPTTPushToTG(article)
		}

		time.Sleep(15 * time.Minute)
	}
}

func SendAllPTTPushToTG(article Const.PTTArticle) (err error) {
	pushDataList, err := PTTCrawler.GetPTTPushData(article.ArticleURL)
	if err != nil {
		fmt.Printf("SendAllPTTPushToTG(): Get PushData fail, err = %s\n", err.Error())
		return
	}

	msg := ""

	for _, push := range pushDataList {
		msg = msg + "\n" + push.Content + "_" + push.IPDatetime
	}

	if msg == "" {
		fmt.Printf("SendAllPTTPushToTG(): %s no push. %s\n" + Tool.GetNowString())
		return
	}

	msg = article.ArticleName + " 結果：" + "\n" + msg

	err = TG.SendMessageToTG(msg)
	if err != nil {
		fmt.Printf("SendAllPTTPushToTG(): Send Message To TG fail, err = %s \n", err.Error())
		return
	}

	urlMapLastPushTime[article.ArticleURL] = pushDataList[len(pushDataList)-1].PushTime
	return
}

func SendNewPTTPushToTG(article Const.PTTArticle) (err error) {
	pushDataList, err := PTTCrawler.GetPTTPushDataAfterTime(article.ArticleURL, urlMapLastPushTime[article.ArticleURL])
	if err != nil {
		fmt.Printf("SendNewPTTPushToTG(): Get PushData fail, err = %s\n", err.Error())
		return
	}

	if len(pushDataList) == 0 {
		fmt.Printf("SendNewPTTPushToTG(): %s no new push. %s\n", article.ArticleName, Tool.GetNowString())
		return
	}

	// 更新最後推文時間
	urlMapLastPushTime[article.ArticleURL] = pushDataList[len(pushDataList)-1].PushTime

	msg := article.ArticleName + " 結果：" + "\n"
	for _, push := range pushDataList {
		msg += push.Content + "_" + push.IPDatetime + "\n"
	}

	err = TG.SendMessageToTG(msg)
	if err != nil {
		fmt.Printf("SendNewPTTPushToTG(): Send Message To TG fail, err = %s \n", err.Error())
		return
	}
	return
}
