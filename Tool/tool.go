package Tool

import "time"

func IsWithinOneHour(ipdatetime string) (bool, error) {
	// 取得當前時間
	now := time.Now()

	// 將今年的年份作為解析時的參考
	year := now.Year()

	// 定義輸入時間的格式
	layout := "01/02 15:04"

	// 把字串轉成 "年/月/日 時:分" 格式來解析
	parsedTime, err := time.ParseInLocation(layout, ipdatetime, time.Local)
	if err != nil {
		return false, err
	}

	// 把當前年份加上去
	parsedTime = parsedTime.AddDate(year-parsedTime.Year(), 0, 0)

	// 計算兩個時間的差異
	diff := now.Sub(parsedTime)

	// 判斷是否在一小時內
	return diff <= time.Hour && diff >= 0, nil
}
