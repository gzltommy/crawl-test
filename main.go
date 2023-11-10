package main

import (
	"fmt"
	"regexp"
)

const str = `<span class="pl">页数:</span> 200<br>
`

var (
	//authorRegexp    = regexp.MustCompile(`<span class="pl"> 作者</span>:[\s\S]*?<a.*?>([^<]+)</a>`)
	//publisherRegexp = regexp.MustCompile(`<span class="pl">出版社:</span>[\s\S]*?<a.*?>([^<]+)</a>`)
	pagesRegexp = regexp.MustCompile(`<span class="pl">页数:</span>([^<]+)<br>`)
	//priceRegexp = regexp.MustCompile(`<span class="pl">定价:</span> ([^<]+)<br>`)
	//scoreRegexp = regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> ([^<]+)</strong>`)
	//introRegexp = regexp.MustCompile(`<div class="intro">[\s\S]*?<p>([^<]+)</p></div>`)
)

func main() {
	match := pagesRegexp.FindString(str)
	fmt.Println("==", match)

	allSubMatch := pagesRegexp.FindAllSubmatch([]byte(str), -1)
	for _, subMath := range allSubMatch {
		for i, matchGroup := range subMath {
			if i == 0 {
				// 属于整个表达式匹配的内容
				continue
			}
			fmt.Printf("\n group[%d]:%s", i, matchGroup)
		}
	}
}
