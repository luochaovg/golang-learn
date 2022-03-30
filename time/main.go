package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func time01() {
	now := time.Now()
	fmt.Println(now.Year(), now.Day())
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	fmt.Printf("%#v", now)
	fmt.Printf("%v", now)

	//ret := time.Unix(1606457598, 0)
	//fmt.Println(ret.Year())
	//fmt.Println(time.Second)

	//fmt.Println(now.Add(time.Hour * 24))

	d := now.Sub(now.Add(time.Hour * 24 * 1))
	fmt.Println(d.Hours())

	//fmt.Println(222222)
	//n := time.Duration(5 * time.Second) // 时间间隔
	//time.Sleep(time.Second * 5)
	//fmt.Println(33333)
	// 定时器
	//timer := time.Tick(time.Second)
	//for t := range timer {
	//	fmt.Println(t) // 1秒执行一次
	//}

	//  2006/1/2 15:04:05
	// Y     m  d  H  M  S
	// 2006  1  2  3  4  5

	// 格式化时间
	//fmt.Println(now.Format("2006-01-02"))
	//fmt.Println(now.Format("2006/01/02 15:04:05.000 PM"))

	// 把一个字符串的时间转换成时间戳
	//timeObj, err := time.Parse("2006-01-02", "1998-09-23")
	//if err != nil {
	//	fmt.Printf("err %v\n", err)
	//}
	//fmt.Println(timeObj.Unix())
}

func main() {
	//format := "2006-01-02 15:04:05"
	//format := "2006-01-02"
	////now := time.Now()
	//now, _ := time.Parse(format, time.Now().Format("2020-03-10 11:00:00"))
	////a, _ := time.Parse(format, "2020-03-10 11:00:00")
	//
	//fmt.Println("aaa \n",  now)
	//if now.After() {
	//	fmt.Println(1)
	//} else {
	//	fmt.Println(0)
	//}

	//st := "2021-12-31"
	//et := "2021-12-31"

	t := time.Now()
	fmt.Println(int(t.Weekday()))


	now := time.Now()

	year, month, _ := now.Date()

	endOfLastMonth := time.Date(year, month, 0, 0, 0, 0, 0, now.Location())
	fmt.Printf("End of the last month: %s\n", endOfLastMonth)
	firstDayOfThisMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	fmt.Printf("The fist day of the actual month: %s\n", firstDayOfThisMonth)
	endOfThisMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, now.Location())
	fmt.Printf("The end of this month: %s\n", endOfThisMonth)



	st := "2022-03-06"
	et := "2022-04-02"
	//
	//st := "2021-10-01"
	//et := "2022-03-31"

	dst, det,  a , b , err := CalcLinkTimeSE(st, et, "weekly")
	fmt.Println(dst, det, err, a, b )

	//theTimeEnd, err := time.Parse("2006-01-02", det)
	//
	//fmt.Println(theTimeEnd.Format("20060102"), err)
	//
	//b, _ := getDivider(st, et, "daily", "avg") // 多少个时间天
	//fmt.Println("B", b)
	//
	//rangtime := Get_Time(st, et)
	//fmt.Println("CC", rangtime)
	//
	//rst := addSliceByIndex([]string{"20211130", "20211210"}, []string{"202112011"}, 2)
	//
	//fmt.Println(rst)
	//return

	//parse, _ := time.Parse("20060102150405", "20141030133525")
	//t3 := parse.Format("2006-01-02 15:04:05")
	//fmt.Println(t3)
	//
	//now = time.Now() // 本地时间
	//fmt.Println(222222)
	//fmt.Println(now)
	//
	//明天的这个时间
	//timea, _ := time.Parse("20060102", "20201128")
	//
	//fmt.Println("aaaaaa", timea.Format("2006-01-02"))
	//
	//// 按照东八区格式解析时间
	//l, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil {
	//	fmt.Printf("11, %v, \n", err)
	//}
	//
	//timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-11-28 14:31:50", l)
	//if err != nil {
	//	fmt.Printf("22, %v, \n", err)
	//}
	//
	//fmt.Println(timeObj, 444)
	//
	//td := timeObj.Sub(now)
	//fmt.Println(td)

}

func addSliceByIndex(old []string, add []string, i int) []string {
	return append(old[:i], append(add, old[i:]...)...) //在第i个位置插入x
}

func Get_Time(start_time, stop_time string) (args []string) {
	tm1, _ := time.Parse("2006-01-02", start_time)
	tm2, _ := time.Parse("2006-01-02", stop_time)

	sInt := tm1.Unix()
	eInt := tm2.Unix()
	for {
		st := time.Unix(sInt, 0).Format("20060102")
		sInt += 86400
		args = append(args, st)
		if sInt > eInt {
			return
		}
	}
}

//确定时间间隔段
func getCountDay(startDate, endDate string, dateType string) int {
	t, _ := time.ParseInLocation("2006-01-02", startDate, time.Local)
	t1, _ := time.ParseInLocation("2006-01-02", endDate, time.Local)
	a := t1.Sub(t).Hours()
	var b float64
	if dateType == "day" {
		b = a/float64(24) + 1
	} else if dateType == "week" {
		b = a / float64(168)
	} else {
		b = a/float64(720) + 1
	}
	return int(math.Ceil(b))
}

// 获取当前时间和前30天
func getNowAndBeforTime(days int) (string, string) {
	linkEt := time.Now().Format("2006-01-02")
	theTime, _ := time.Parse("2006-01-02", linkEt)
	linkSt := theTime.AddDate(0, 0, days).Format("2006-01-02")
	return linkSt, linkEt
}

// 获取环比的开始时间和结束时间
func CalcLinkTime(st, et, dt string) (string, string, error) {
	theTimeStart, err := time.Parse("2006-01-02", st)
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("time parse error:%s ", st))
	}
	theTimeEnd, err := time.Parse("2006-01-02", et)
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("time parse error:%s ", et))
	}
	if dt == "daily" || dt == "weekly" || dt == "week" || dt == "day" {
		dur := theTimeEnd.Sub(theTimeStart).Hours()
		linkEt := theTimeStart.AddDate(0, 0, -1).Format("2006-01-02")
		linkSt := theTimeStart.AddDate(0, 0, -int(dur/24)-1).Format("2006-01-02")
		return linkSt, linkEt, nil
	} else if dt == "monthly" || dt == "month" {
		linkEt := theTimeStart.AddDate(0, -1, 0).Format("2006-01-02")
		linkSt := theTimeStart.AddDate(0, -(int(theTimeEnd.Month())-int(theTimeStart.Month()))-1, 0).Format("2006-01-02")
		return linkSt, linkEt, nil
	}
	return "", "", nil

}

func getDivider(st, et, dt, avgFlg string) (int, error) {
	//var divider float64
	divider := 1
	theTimeStart, err := time.Parse("2006-01-02", st)
	if err != nil {
		return divider, errors.New(fmt.Sprintf("time parse error:%s ", st))
	}
	theTimeEnd, err := time.Parse("2006-01-02", et)
	if err != nil {
		return divider, errors.New(fmt.Sprintf("time parse error:%s ", et))
	}
	subDays := SubDays(theTimeEnd, theTimeStart) + 1
	if avgFlg == "avg" {
		if dt == "daily" {
			divider = subDays
		} else if dt == "weekly" {
			divider = subDays / 7
		} else if dt == "monthly" {
			divider = int(math.Ceil(float64(subDays) / 31))
		}
	}
	return divider, nil
}

// 获取环比的开始时间和结束时间
func CalcLinkTimeSE(st, et, dt string) (string, string, string, string, error) {
	theTimeStart, err := time.Parse("2006-01-02", st)
	if err != nil {
		return "", "", "", "", errors.New(fmt.Sprintf("time parse error:%s ", st))
	}
	theTimeEnd, err := time.Parse("2006-01-02", et)
	if err != nil {
		return "", "", "", "", errors.New(fmt.Sprintf("time parse error:%s ", et))
	}
	subDays := SubDays(theTimeEnd, theTimeStart)
	if dt == "daily" {
		dur := theTimeEnd.Sub(theTimeStart).Hours()
		linkEt := theTimeStart.AddDate(0, 0, -1).Format("2006-01-02")
		linkSt := theTimeStart.AddDate(0, 0, -int(dur/24)-1).Format("2006-01-02")
		return st, et, linkSt, linkEt, nil
	} else if dt == "weekly" {
		et := theTimeEnd.AddDate(0, 0, -6).Format("2006-01-02")
		st := theTimeStart.AddDate(0, 0, 0).Format("2006-01-02")
		linkEt := theTimeStart.AddDate(0, 0, -7).Format("2006-01-02")
		linkSt := theTimeStart.AddDate(0, 0, -subDays-1).Format("2006-01-02")
		return st, et, linkSt, linkEt, nil
	} else if dt == "monthly" {
		if theTimeStart.Month() != theTimeEnd.Month() {
			et := theTimeEnd.AddDate(0, 0, -(theTimeEnd.Day() - 1)).Format("2006-01-02")
			st := theTimeStart.AddDate(0, 0, 0).Format("2006-01-02")
			linkEt := theTimeStart.AddDate(0, -1, 0).Format("2006-01-02")
			linkSt := theTimeStart.AddDate(0, -(int(theTimeEnd.Month())-int(theTimeStart.Month()))-1, 0).Format("2006-01-02")
			return st, et, linkSt, linkEt, nil
		} else {
			// et := theTimeStart.AddDate(0, 0, 0).Format("2006-01-02")
			// st := theTimeStart.AddDate(0, 0, 0).Format("2006-01-02")
			linkEt := theTimeStart.AddDate(0, -1, 0).Format("2006-01-02")
			linkSt := theTimeStart.AddDate(0, -1, 0).Format("2006-01-02")
			return st, et, linkSt, linkEt, nil
		}
	}
	return "", "", "", "", nil
}

// 获取环比的开始时间和结束时间
func CalcLinkTimeSEReq(st, et, dt string) (string, string, error) {
	theTimeStart, err := time.Parse("2006-01-02", st)
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("time parse error:%s ", st))
	}
	theTimeEnd, err := time.Parse("2006-01-02", et)
	if err != nil {
		return "", "", errors.New(fmt.Sprintf("time parse error:%s ", et))
	}
	subDays := SubDays(theTimeEnd, theTimeStart)
	if dt == "daily" {
		dur := theTimeEnd.Sub(theTimeStart).Hours()
		linkEt := theTimeStart.AddDate(0, 0, -1).Format("2006-01-02")
		linkSt := theTimeStart.AddDate(0, 0, -int(dur/24)-1).Format("2006-01-02")
		return linkSt, linkEt, nil
	} else if dt == "weekly" {
		linkEt := theTimeStart.AddDate(0, 0, -1).Format("2006-01-02")
		linkSt := theTimeStart.AddDate(0, 0, -subDays-1).Format("2006-01-02")
		return linkSt, linkEt, nil
	} else if dt == "monthly" {
		linkEt := theTimeStart.AddDate(0, 0, -1).Format("2006-01-02")
		linkSt := theTimeStart.AddDate(0, -(int(theTimeEnd.Month())-int(theTimeStart.Month()))-1, 0).Format("2006-01-02")
		return linkSt, linkEt, nil
	}
	return "", "", nil
}

func SubDays(t1, t2 time.Time) (day int) {
	day = int(t1.Sub(t2).Hours() / 24)
	return
}

//取最小的时间
func minTime(list []string) string {
	if len(list) == 0 {
		return ""
	}
	min := list[0]
	for _, value := range list {
		if value < min && value != "" {
			min = value
		}
	}
	return min
}

//列表去重
func RemoveDuplicates(list []string) (ret []string) {
	listLen := len(list)
	for i := 0; i < listLen; i++ {
		if (i > 0 && list[i-1] == list[i]) || len(list[i]) == 0 {
			continue
		}
		ret = append(ret, list[i])
	}
	return
}

//根据结束日期确定开始日期
func getHomeStartDateByEndDate(endDate string, months, days int) string {
	stamp, _ := time.ParseInLocation("2006-01-02", endDate, time.Local)
	startDate := stamp.AddDate(0, -months, -days).Format("2006-01-02")
	return startDate
}
