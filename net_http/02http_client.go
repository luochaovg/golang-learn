package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// http/client
// https://www.liwenzhou.com/posts/Go/go_http/

// 后端请求比较频繁是，可共用一个client，（不然会造成没有关闭的连接够多，占用网络IO）
var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func getRequest() {
	urlPath := "http://127.0.0.1:8080/get?name=law&age=18"
	// 1.第一种
	//resp, err := http.Get(urlPath)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}

	// 2.第二种
	data := url.Values{}
	urlObj, _ := url.Parse(urlPath)
	data.Set("name", "罗超")
	data.Set("age", "12")
	queryStr := data.Encode() // URL encode 之后的URL
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)

	//resp, err := http.DefaultClient.Do(req) // 发请求
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}

	// 禁用keepAlive的client ， 请求不是特别频繁，用完就关闭该连接
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	defer resp.Body.Close() // 一定要记得关闭resp.Body

	// 从resp中把服务端返回的数据读出来
	// var data []byte
	// resp.Body.Read()
	// resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func postForm() {
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	// contentType := "application/x-www-form-urlencoded"
	// data := "name=小王子&age=18"
	// json数据
	contentType := "application/json"
	data := `{"name":"小王子","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close() // TODO 注意关闭
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

func main() {
	getRequest()
}
