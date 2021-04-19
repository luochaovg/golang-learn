package main

import "fmt"

// 经典排序算法总结与Go实现
// https://www.jianshu.com/p/06b6424042d5
// https://blog.csdn.net/books1958/article/details/42267301

// 冒泡
// 思路:正如“冒泡”二字，我的理解是重复依次比较相邻的两个数，大的数放在后面，小的数放在前面，一直重复到没有任何一对数字需要交换位置为止。就像冒泡一样，大的数不断浮上来。
// 插入排序和冒泡排序在平均和最坏情况下的时间复杂度都是 O(n^2)，最好情况下都是 O(n)，空间复杂度是 O(1)。
func Bubble(nums []int) []int {
	//冒泡排序（排序10000个随机整数，用时约145ms）
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}

	return nums
}

// 选择
// 选择排序的原理是，对给定的数组进行多次遍历，每次均找出最小的一个值的索引。
func SelectSort(nums []int) []int {
	length := len(nums)
	//选择排序（排序10000个随机整数，用时约45ms）
	for i := 0; i < length; i++ {
		minIndex := 0
		for j := 1; j < length-i; j++ {
			if nums[j] > nums[minIndex] {
				minIndex = j
			}
		}

		nums[length-i-1], nums[minIndex] = nums[minIndex], nums[length-i-1]
	}

	return nums
}

// 插入
// 插入排序的原理是，从第二个数开始向右侧遍历，每次均把该位置的元素移动至左侧，放在放在一个正确的位置（比左侧大，比右侧小）。
func InsertSort(nums []int) []int {
	//插入排序（排序10000个整数，用时约30ms）
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			j := i - 1
			temp := nums[i]

			for j >= 0 && nums[j] > temp {
				nums[j+1] = nums[j]
				j--
			}

			nums[j+1] = temp
		}

	}

	return nums
}

// 希尔

// 堆排序（没有理解）
// 思路和算法：堆排序的思想就是先将待排序的序列建成大根堆，使得每个父节点的元素大于等于它的子节点。此时整个序列最大值即为堆顶元素，
// 我们将其与末尾元素交换，使末尾元素为最大值，然后再调整堆顶元素使得剩下的 n-1n−1 个元素仍为大根堆，再重复执行以上操作我们即能得到一个有序的序列。
// 时间复杂度 O(nlogn)  , 空间复杂度 O(1)
// 参考： https://gist.github.com/xianlubird/cfe7e861ea22042e11b0
func HeapSort() {

}

// 归并
//思路:归并排序利用了分治的思想来对序列进行排序。对一个长为 n 的待排序的序列，我们将其分解成两个长度为 n/2 的子序列。每次先递归调用函数使两个子序列有序，然后我们再线性合并两个有序的子序列使整个序列有序。
// 时间复杂度：O(nlogn)  空间复杂度：O(n)
// golang https://learnku.com/articles/45755
func MergeSort(nums []int) []int {
	length := len(nums)

	if length == 1 {
		return nums
	}

	m := length / 2

	leftNums := MergeSort(nums[:m])
	rightNums := MergeSort(nums[m:])

	return merge2(leftNums, rightNums)
}

func merge2(l, r []int) []int {
	lLen := len(l)
	rLen := len(r)

	res := make([]int, 0)

	lIndex, rIndex := 0, 0
	for lIndex < lLen && rIndex < rLen {
		if l[lIndex] > r[rIndex] {
			res = append(res, r[rIndex])
			rIndex++
		} else {
			res = append(res, l[lIndex])
			lIndex++
		}

	}

	if lIndex < lLen { //左边的还有剩余元素
		res = append(res, l[lIndex:]...)
	}
	if rIndex < rLen {
		res = append(res, r[rIndex:]...)
	}

	return res
}

// 快排序
// 快速排序通过一个切分元素将数组分为两个子数组，左子数组小于等于切分元素，右子数组大于等于切分元素，将这两个子数组排序也就将整个数组排序了。
// 均情况下快速排序的时间复杂度是Θ(𝑛lg𝑛)，最坏情况是𝑛2，但通过随机算法可以避免最坏情况。由于递归调用，快排的空间复杂度是Θ(lg𝑛)。
func QuickSort(nums []int) []int {
	//插入排序（排序10000个整数，用时约30ms）

	if len(nums) <= 1 {
		return nums
	}

	base := nums[0]

	var (
		left  []int
		right []int
	)

	for i := 1; i < len(nums); i++ {
		if nums[i] > base {
			right = append(right, nums[i])
		} else {
			left = append(left, nums[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	temp := append(left, base)
	return append(temp, right...)
}

func main() {
	demoNums := []int{30, 4, 3, 5, 12, 1, 32, 23, 2}
	//fmt.Println(Bubble(demoNums))
	//fmt.Println(SelectSort(demoNums))
	// fmt.Println(InsertSort(demoNums))
	//fmt.Println(QuickSort(demoNums))
	fmt.Println(MergeSort(demoNums))
}
