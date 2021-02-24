package main

import (
	"errors"
	"fmt"
)

// 两数之和
//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//你可以按任意顺序返回答案。

func twoSum(nums []int, target int) ([]int, error) {
	if len(nums) <= 1 {
		return nil, errors.New("数组切片不能小于一个")
	}

	var ret []int

	// 1，从第一个循环，
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if k1 == k2 {
				break
			}

			if v1+v2 == target {
				ret = append(ret, k2, k1)
			}
		}
	}

	return ret, nil
}

func toSum2(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))

	for index, num := range nums {
		if value, ok := hashmap[num]; ok {
			nums = nums[:2]
			nums[0] = value
			nums[1] = index
			return nums
		}
		hashmap[target-num] = index
	}
	return nil
}
func main() {
	nums := []int{12, 34, 65, 2, 7, 11, 15}
	target := 9

	//ret, _ := twoSum(nums, target)
	ret := toSum2(nums, target)

	fmt.Println(ret)
}
