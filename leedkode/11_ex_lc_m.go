// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.

package main

import "fmt"

// func maxArea(height []int) int { // slowly
// 	res := 0
// 	for i := 0; i < len(height); i++ {
// 		for j := i + 1; j < len(height); j++ {
// 			if height[i] < height[j] {
// 				res = max(res, (j-i)*height[i])
// 			} else {
// 				res = max(res, (j-i)*height[j])
// 			}
// 		}
// 	}

// 	return res
// }

func maxArea(height []int) int { // все равно медленно, переделать на по типу алгоритма скользяшего окна, только на уменьшение

	res, tmp, tmpJ := 0, 0, 0
	light := 10000

	for i := 0; i < len(height); i++ {
		if height[i] < tmp {
			continue
		}
		tmpJ = 0
		for j := len(height); j > 0 && j-1 > i; j-- {
			if j == len(height) || (j > 1 && height[j-1] > tmpJ) {

				light = min(height[i], height[j-1])
				tmp = max(tmp, height[i])
				tmpJ = max(tmpJ, height[j-1])
				res = max(res, (j-i-1)*light)

				if height[i] <= height[j-1] {
					break
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
}
