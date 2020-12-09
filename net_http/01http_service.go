package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

func f1(w http.ResponseWriter, r *http.Request) {
	//str := `<h1 style="color:red">hello world</h1>` // 字符串本质就是字节类型的切片
	//w.Write([]byte(str))

	b, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}

	w.Write(b)
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))

	// 对于get请求，参数都放在url，请求体是没有数据的
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)

	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/hello", f1)
	http.HandleFunc("/get", get)

	http.ListenAndServe("127.0.0.1:8080", nil)
}
