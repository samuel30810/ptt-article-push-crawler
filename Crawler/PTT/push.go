package crawler

import (
	"fmt"
	"headphone/Const"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GetPTTPushData(url string) (pushDataList []Const.PTTPushData, err error) {
	// 發送 HTTP GET 請求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("GetPTTPushData(): http.Get fail, err = %s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	// 檢查 HTTP 狀態碼是否正確
	if resp.StatusCode != 200 {
		fmt.Printf("GetPTTPushData(): status code error, status code = %d, status = %s\n", resp.StatusCode, resp.Status)
		return
	}

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("GetPTTPushData(): goquery.NewDocumentFromReader fail, err = %s\n", err.Error())
		return
	}

	// 找到所有 class 為 "push" 的 div
	doc.Find("div.push").Each(func(i int, s *goquery.Selection) {
		// 抓取 push-content 和 push-ipdatetime
		content := s.Find(".push-content").Text()
		ipdatetime := cleanIpdatetimeToString(s.Find(".push-ipdatetime").Text())

		pushTime, _ := cleanIpdatetimeToTime(ipdatetime)

		// 去除 content 前面的 ": "
		content = content[2:]

		// 將資料存入 PushData struct
		pushData := Const.PTTPushData{
			Content:    content,
			IPDatetime: ipdatetime,
			PushTime:   pushTime,
		}

		// 將 PushData 加入 slice
		pushDataList = append(pushDataList, pushData)
	})
	return
}

func GetPTTPushDataAfterTime(url string, afterTime time.Time) (pushDataList []Const.PTTPushData, err error) {
	// 發送 HTTP GET 請求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("GetPTTPushData(): http.Get fail, err = %s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	// 檢查 HTTP 狀態碼是否正確
	if resp.StatusCode != 200 {
		fmt.Printf("GetPTTPushData(): status code error, status code = %d, status = %s\n", resp.StatusCode, resp.Status)
		return
	}

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("GetPTTPushData(): goquery.NewDocumentFromReader fail, err = %s\n", err.Error())
		return
	}

	// 找到所有 class 為 "push" 的 div
	doc.Find("div.push").Each(func(i int, s *goquery.Selection) {
		// 抓取 push-content 和 push-ipdatetime
		content := s.Find(".push-content").Text()
		ipdatetime := cleanIpdatetimeToString(s.Find(".push-ipdatetime").Text())

		pushTime, _ := cleanIpdatetimeToTime(ipdatetime)

		// 只抓指定時間後的推文
		if !pushTime.After(afterTime) {
			return
		}

		// 去除 content 前面的 ": "
		content = content[2:]

		// 將資料存入 PushData struct
		pushData := Const.PTTPushData{
			Content:    content,
			IPDatetime: ipdatetime,
			PushTime:   pushTime,
		}

		// 將 PushData 加入 slice
		pushDataList = append(pushDataList, pushData)
	})
	return
}

func cleanIpdatetimeToString(ipdatetime string) (cleanIpdatetime string) {
	ipdatetime = strings.Replace(ipdatetime, "\x0a", "", -1)
	ipdatetime = strings.Replace(ipdatetime, " ", "", 1)

	cleanIpdatetime = ipdatetime
	return
}

func cleanIpdatetimeToTime(ipdatetime string) (parsedTime time.Time, err error) {

	year := time.Now().Year()
	layout := "01/02 15:04"

	parsedTime, err = time.ParseInLocation(layout, ipdatetime, time.Local)
	if err != nil {
		fmt.Printf("cleanIpdatetimeToTime(): time.ParseInLocation fail, err = %s\n", err.Error())
		return
	}

	parsedTime = parsedTime.AddDate(year-parsedTime.Year(), 0, 0)

	return
}
