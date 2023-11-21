package douban

import (
	"github.com/gzltommy/crawl-test/04.MulTask/types"
	"regexp"
)

// <a href="https://book.douban.com/subject/30192800/" title="Python神经网络编程" onclick="moreurl(this,{i:'0',query:”,subject_id:'30192800',from:'book_subject_search'})">
var bookListRegexp = regexp.MustCompile(`<a href="([^"]+)" title="([^"]+)"`)

func BookList(content []byte) types.ParseResult {
	match := bookListRegexp.FindAllSubmatch(content, -1)
	result := types.ParseResult{}
	for _, m := range match {
		bookName := string(m[2])
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, types.Request{
			Url: string(m[1]),
			ParseFun: func(c []byte) types.ParseResult {
				return BookDetail(c, bookName)
			},
		})
	}
	return result
}
