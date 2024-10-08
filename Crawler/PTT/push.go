package crawler

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type PTTPushData struct {
	Content    string
	IPDatetime string
}

func GetPTTPushData(url string) (pushDataList []PTTPushData, err error) {
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
		ipdatetime := s.Find(".push-ipdatetime").Text()

		// 去除 content 前面的 ": "
		content = content[2:]

		// 將資料存入 PushData struct
		pushData := PTTPushData{
			Content:    content,
			IPDatetime: ipdatetime,
		}

		// 將 PushData 加入 slice
		pushDataList = append(pushDataList, pushData)
	})
	return
}
