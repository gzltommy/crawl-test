package model

import "encoding/json"

type BookDetail struct {
	Author    string
	Publisher string
	BookPages string
	Price     string
	Score     string
	Intro     string
}

func (b *BookDetail) String() string {
	buf, _ := json.Marshal(b)
	return string(buf)
}
