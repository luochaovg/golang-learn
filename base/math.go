package main

import (
	"fmt"
	"math"
)

func main() {

	//a := 7
	//a++
	//
	//fmt.Println(a)

	// & 按位与
	// 5 二进制 101
	// 3 二进制 011
	//fmt.Println(5 & 3)
	//fmt.Println(5 | 3)
	//fmt.Println(5 << 1) // 1010  = 2*2^3 + 2*2^1
	//fmt.Println(5 >> 1) // 10 = 2
	//
	//var a uint = 1
	//var b uint = 2
	//
	//c := a - b
	//fmt.Println(c)

	arr := []int{631, 240, 3, 2, 1, 1, 1}
	//arr := []int{3,3,3}
	precision := 2

	for i := range arr {
		f := GetPercentWithPrecision(arr, i, precision)
		fmt.Printf("index:%d ,  Value:%d , Per:%.2f \n", i, arr[i], f)
	}

}

//GetPercentWithPrecision eleme 最大饥饿算法
/**
 * @param {Array.<number>} valueList a list of all data 一列数据
 * @param {number} idx index of the data to be processed in valueList 索引值（数组下标）
 * @param {number} precision integer number showing digits of precision 精度值
 * @return {number} percent ranging from 0 to 100 返回百分比从0到100
 * eg. GetPercentWithPrecision([]int{3, 3, 3}, 0, 2)   // 33.34
 */
func GetPercentWithPrecision(valueList []int, idx int, precision int) (res float64) {
	if idx >= len(valueList) {
		res = 0
		return
	}
	var digits = math.Pow10(precision)
	var total = 0
	for _, value := range valueList {
		total += value
	}
	if total == 0 {
		return
	}
	var votesPerQuota = make([]float64, len(valueList))
	for i, value := range valueList {
		votesPerQuota[i] = float64(value) / float64(total) * digits * 100
	}

	var targetSeats = digits * 100
	var currentSeatsSum int
	var seats = make([]int, len(valueList))
	var remainder = make([]float64, len(valueList))
	for i, votes := range votesPerQuota {
		seats[i] = int(votes)
		currentSeatsSum += int(votes)
		remainder[i] = votes - float64(seats[i])
	}

	for currentSeatsSum < int(targetSeats) {
		var max = float64(0)
		var maxID = 0
		for i, remainVal := range remainder {
			if remainVal > max {
				max = remainVal
				maxID = i
			}
		}
		seats[maxID]++
		remainder[maxID] = 0
		currentSeatsSum++
	}

	res = float64(seats[idx]) / digits
	return
}
