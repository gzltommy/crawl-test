package main

import (
	"github.com/gzltommy/crawl-test/03.SingleTask/engine"
	"github.com/gzltommy/crawl-test/03.SingleTask/parse"
)

func main() {
	//engine.Run(engine.Request{
	//	Url:      "https://book.douban.com/",
	//	ParseFun: parse.Tag,
	//})

	//engine.Run(engine.Request{
	//	Url:      "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C",
	//	ParseFun: parse.BookList,
	//})

	//engine.Run(engine.Request{
	//	Url:      "https://book.douban.com/subject/30192800/",
	//	ParseFun: parse.BookDetail,
	//})

	// 最终：从最开始爬，并嵌套爬
	engine.Run(engine.Request{
		Url:      "https://book.douban.com/",
		ParseFun: parse.Tag,
	})
}
