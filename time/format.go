package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// 参考： https://www.liwenzhou.com/posts/Go/json_tricks_in_go/

type Data struct {
	ExpireTime ftime `json:"expire_time"`
}

type ftime time.Time

func (f ftime) MarshalJSON() ([]byte, error) { // struct -> json string
	var stamp = fmt.Sprintf("\"%s\"", time.Time(f).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (f *ftime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	// 按照东八区格式解析时间
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}

	t, _ := time.ParseInLocation("2006-01-02 15:04:05", strings.Trim(string(data), `"`), local)
	*f = ftime(t)
	return err
}

func main() {
	d1 := Data{
		ExpireTime: ftime(time.Now()),
	}

	b, err := json.Marshal(&d1)
	if err != nil {
		fmt.Printf("json.Marshal o1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)

	//
	//
	jsonStr := `{"expire_time":"2021-09-07 15:35:26"}`
	var d2 Data
	if err := json.Unmarshal([]byte(jsonStr), &d2); err != nil {
		fmt.Printf("json.Unmarshal jsonStr1 failed, err:%v\n", err)
		return
	}

	fmt.Printf("c1:%#v \n", time.Time(d2.ExpireTime).Format("2006-01-02 15:04:05"))
}
