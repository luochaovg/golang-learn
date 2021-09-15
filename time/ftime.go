package fo

import (
	"fmt"
	"time"
)

package models

import (
"fmt"
"time"
)


type ztime time.Time

func (c ztime) MarshalJSON() ([]byte, error) {

	var stamp = fmt.Sprintf("\"%s\"", time.Time(c).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// 获取time.Time类型，方便拓展方法
func (c ztime) Time() time.Time {
	return time.Time(c)
}

// 格式化
func (c ztime) Format() string {
	return time.Time(c).Format("2006-01-02 15:04:05")
}

// 简单格式化
func (c ztime) FormatSimple() string {
	return time.Time(c).Format("2006-01-02")
}