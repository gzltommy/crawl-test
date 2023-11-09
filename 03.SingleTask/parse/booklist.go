package parse

import (
	"github.com/gzltommy/crawl-test/03.SingleTask/engine"
	"regexp"
)

// <a href="https://book.douban.com/subject/30192800/" title="Python神经网络编程" onclick="moreurl(this,{i:'0',query:”,subject_id:'30192800',from:'book_subject_search'})">
var bookListRegexp = regexp.MustCompile(`<a href="([^"]+)" title="([^"]+)"`)

func ParseBookList(content []byte) engine.ParseResult {
	match := bookListRegexp.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:      "https://book.douban.com" + string(m[1]),
			ParseFun: engine.NilParse,
		})
	}
	return result
}
