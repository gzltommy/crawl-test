package parse

import (
	"regexp"
)

// <a href="/tag/小说" class="tag">小说</a>
var (
	authorRegexp    = regexp.MustCompile(`<span class="pl"> 作者</span>:[\s\S]*?<a.*?>([^<]+)</a>`)
	publisherRegexp = regexp.MustCompile(`<span class="pl">出版社:</span>[\s\S]*?<a.*?>([^<]+)</a>`)
	pagesRegexp     = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br>`)
	priceRegexp     = regexp.MustCompile(`<span class="pl">定价:</span> ([^<]+)<br>`)
	scoreRegexp     = regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> ([^<]+)</strong>`)
	introRegexp     = regexp.MustCompile(`<div class="intro">[\s\S]*?<p>([^<]+)</p></div>`)
)

func ParseBookDetail(content []byte) {
	//match := bookListRegexp.FindAllSubmatch(content, -1)
	//result := engine.ParseResult{}
	//for _, m := range match {
	//	result.Items = append(result.Items, m[2])
	//	result.Requests = append(result.Requests, engine.Request{
	//		Url:      "https://book.douban.com" + string(m[1]),
	//		ParseFun: engine.NilParse,
	//	})
	//}
	//return result
}
