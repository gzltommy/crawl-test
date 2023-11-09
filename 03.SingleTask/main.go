package main

import (
	"github.com/gzltommy/crawl-test/03.SingleTask/engine"
	"github.com/gzltommy/crawl-test/03.SingleTask/parse"
)

func main() {
	//engine.Run(engine.Request{
	//	Url:      "https://book.douban.com/",
	//	ParseFun: parse.ParseTag,
	//})

	engine.Run(engine.Request{
		Url:      "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C",
		ParseFun: parse.ParseBookList,
	})
}
