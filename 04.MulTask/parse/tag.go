package parse

import (
	"github.com/gzltommy/crawl-test/04.MulTask/types"
	"regexp"
)

// <a href="/tag/小说" class="tag">小说</a>
var tagRegexp = regexp.MustCompile(`<a href="([^"]+)" class="tag">([^<]+)</a>`)

func Tag(content []byte) types.ParseResult {
	match := tagRegexp.FindAllSubmatch(content, -1)
	result := types.ParseResult{}
	for _, m := range match {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, types.Request{
			Url:      "https://book.douban.com" + string(m[1]),
			ParseFun: BookList,
		})
	}
	return result
}
