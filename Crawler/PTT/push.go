package crawler

import (
	"log"
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
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 檢查 HTTP 狀態碼是否正確
	if resp.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
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
