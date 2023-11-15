package fetcher

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

// Fetch 模拟浏览器访问
func Fetch(url string) ([]byte, error) {
	// 1.设置一个请求
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")

	// 2.新建一个客户端
	client := &http.Client{}

	// 3.使用客户端去发起请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d", resp.StatusCode)
		return nil, fmt.Errorf("error status code:%d", resp.StatusCode)
	}

	// 读取内容到一个 bufio 的 reader 中
	bodyReader := bufio.NewReader(resp.Body)

	// 探测出字符集编码
	e := DetermineEncoding(bodyReader)

	// 转换编码为 utf-8
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	// 读取内容
	return io.ReadAll(utf8Reader)
}

// DetermineEncoding 探测出字符集编码
func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetch error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
