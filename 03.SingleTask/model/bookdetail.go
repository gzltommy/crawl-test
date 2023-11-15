package model

import "encoding/json"

type BookDetail struct {
	BookName  string `json:"book_name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	BookPages string `json:"book_pages"`
	Price     string `json:"price"`
	Score     string `json:"score"`
	Intro     string `json:"intro"`
}

func (b *BookDetail) String() string {
	buf, _ := json.Marshal(b)
	return string(buf)
}
