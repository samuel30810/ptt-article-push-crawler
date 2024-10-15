package Const

import "time"

type PTTArticle struct {
	ArticleName string        `json:"articleName"`
	ArticleURL  string        `json:"articleURL"`
}

type PTTPushData struct {
	Content    string
	IPDatetime string
	PushTime   time.Time
}
