package parse

import (
	"github.com/gzltommy/crawl-test/03.SingleTask/engine"
	"regexp"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^<]+)</a>`

func ParseContent(content []byte) engine.ParseResult {
	//<a href="/tag/小说" class="tag">小说</a>
	re := regexp.MustCompile(regexpStr)
	match := re.FindAllSubmatch(content, -1)

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
