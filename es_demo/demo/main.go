package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Persion struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {

	client, e := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200"))
	if e != nil {
		panic(e)
	}

	fmt.Printf("connect to es success")
	p1 := Persion{
		Name:    "luochao",
		Age:     30,
		Married: false,
	}
	response, e := client.Index().Index("user").BodyJson(p1).Do(context.Background())
	if e != nil {
		panic(e)
	}
	fmt.Printf("Indexed user %s to index %s , type %s \n", response.Id, response.Index, response.Type)
}
