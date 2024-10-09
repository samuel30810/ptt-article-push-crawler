package Tool

import "time"

func GetNowString() string {
	now := time.Now()
	layout := "2006/01/02 15:04:05"
	return now.Format(layout)
}
