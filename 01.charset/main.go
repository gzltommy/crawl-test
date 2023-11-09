package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.chinanews.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d", resp.StatusCode)
		return
	}

	// 读取内容到一个 bufio 的 reader 中
	bodyReader := bufio.NewReader(resp.Body)

	// 探测出字符集编码
	e := determineEncoding(bodyReader)

	// 转换编码为 utf-8
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	result, err := io.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", result)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetch error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
