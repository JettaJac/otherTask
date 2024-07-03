// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.

package main

import (
	"fmt"
	"time"
)

func maxArea(height []int) int { // slowly
	start := time.Now()
	res := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			if height[i] < height[j] {
				res = max(res, (j-i)*height[i])
			} else {
				res = max(res, (j-i)*height[j])
			}
		}
	}
	finish := time.Now()
	fmt.Println("maxArea", finish.Sub(start))
	return res
}

func maxArea2(height []int) int { // все равно медленно, переделать на по типу алгоритма скользяшего окна, только на уменьшение
	start := time.Now()
	res, tmpI, tmpJ := 0, 0, 0
	light := 10000

	for i := 0; i < len(height); i++ {
		if height[i] < tmpI {
			continue
		}
		tmpJ = 0
		for j := len(height); j > 0 && j-1 > i; j-- {
			if j == len(height) || (j > 1 && height[j-1] > tmpJ) {

				light = min(height[i], height[j-1])
				tmpI = max(tmpI, height[i])
				tmpJ = max(tmpJ, height[j-1])
				res = max(res, (j-i-1)*light)

				if height[i] <= height[j-1] {
					break
				}
			}
		}
	}
	finish := time.Now()
	fmt.Println("maxArea2", finish.Sub(start))

	return res
}

func maxArea3(height []int) int {
	res, tmpI, tmpJ := 0, 0, 0
	light := 10000

	for i, j := 0, len(height)-1; i < j; /* i < len(height) && j >= 0*/ {
		if height[i] < tmpI { // возможно есть смысл ставить цикл
			i++
			continue
		}
		if height[j] < tmpJ { // возможно есть смысл ставить цикл
			j--
			continue
		}

		tmpI = max(tmpI, height[i])
		tmpJ = max(tmpJ, height[j])

		light = min(height[i], height[j])
		res = max(res, (j-i)*light)

		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}

	return res
}

// func main() {
// 	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
// 	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
// 	fmt.Println(maxArea3([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
// }
