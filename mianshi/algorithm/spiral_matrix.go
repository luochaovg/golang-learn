package main

import "fmt"

// LeetCode 54 螺旋矩阵
// https://leileiluoluo.com/posts/leetcode-spiral-matrix.html

func spiralOrder(matrix [][]int) []int {
	// special
	if 0 == len(matrix) {
		return []int{}
	}

	var elements []int

	leftBoundary, rightBoundary := 0, len(matrix[0])-1
	topBoundary, bottomBoundary := 0, len(matrix)-1

	for leftBoundary < rightBoundary &&
		topBoundary < bottomBoundary {

		// top
		for j := leftBoundary; j < rightBoundary; j++ {
			elements = append(elements, matrix[topBoundary][j])
		}

		// right
		for i := topBoundary; i < bottomBoundary; i++ {
			elements = append(elements, matrix[i][rightBoundary])
		}

		// bottom
		for j := rightBoundary; j > leftBoundary; j-- {
			elements = append(elements, matrix[bottomBoundary][j])
		}

		// left
		for i := bottomBoundary; i > topBoundary; i-- {
			elements = append(elements, matrix[i][leftBoundary])
		}

		leftBoundary++
		rightBoundary--
		topBoundary++
		bottomBoundary--
	}

	if leftBoundary == rightBoundary && topBoundary <= bottomBoundary {
		for i := topBoundary; i <= bottomBoundary; i++ {
			elements = append(elements, matrix[i][leftBoundary])
		}

		return elements
	}

	if topBoundary == bottomBoundary && leftBoundary <= rightBoundary {
		for j := leftBoundary; j <= rightBoundary; j++ {
			elements = append(elements, matrix[topBoundary][j])
		}
	}

	return elements
}

func main() {
	//matrix := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}

	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println(spiralOrder(matrix))

}
