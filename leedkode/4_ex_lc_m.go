package main

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).

// Example 1:
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

import (
	"fmt"
	"math"
	"sort"
)

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 { //Done
// 	size := len(nums1) + len(nums2)
// 	pointMedian := int((size)/2 + 1)
// 	resultSlice := make([]int, pointMedian)
// 	z := 0
// 	for i, j := 0, 0; z < pointMedian; z++ {

// 		for ; i < len(nums1) && j < len(nums2) && z < pointMedian; z++ {
// 			if nums1[i] <= nums2[j] {
// 				resultSlice[z] = nums1[i]
// 				i++
// 			} else {
// 				resultSlice[z] = nums2[j]
// 				j++
// 			}
// 		}
// 		if i < len(nums1) && z < pointMedian {
// 			resultSlice[z] = nums1[i]
// 			i++
// 		}
// 		if j < len(nums2) && z < pointMedian {
// 			resultSlice[z] = nums2[j]
// 			j++
// 		}
// 	}

//		if size%2 == 0 {
//			return float64(resultSlice[pointMedian-1]+resultSlice[pointMedian-2]) / 2
//		}
//		return float64(resultSlice[pointMedian-1])
//	}

func changeResult2(nums, result []int) (a, b int) {
	memory := math.MinInt64 // !!!
	preMemory := math.MinInt64
	tmp := result[0]
	fmt.Println(result)
	for i, j := cap(result)-1, 0; i > 0; {
		for len(result) != cap(result) {
			result = append(result, nums[j])
			memory = max(result[i], nums[j])
			i--
			j++
		}
		if j >= len(nums) || i <= 0 || result[i] <= nums[j] {
			if memory == math.MinInt64 {
				memory = result[i-1]
			}
			tmp = memory
			if i > 0 && j < len(nums) && result[i] == nums[j] {
				tmp = max(result[i-1], memory)
			}
			return tmp, max(result[i], preMemory)
		}

		if result[i] > nums[j] {
			result[i] = nums[j]
			preMemory = memory
			memory = nums[j] // а вообще предыдущее значением
			i--
			j++
		}
		tmp = max(result[i], preMemory)
	}
	return memory, tmp
}

func changeResult(nums, result []int) (a, b int) { //медленнее чем, певый(())
	memory := math.MinInt64
	preMemory := math.MinInt64
	tmp := result[0]
	for i, j := cap(result)-1, 0; i > 0; {
		for len(result) != cap(result) {
			result = append(result, nums[j])
			i--
			j++
		}
		if j >= len(nums) || i <= 0 || result[i] <= nums[j] {

			sort.Ints(result)
			size := len(result) - 1
			return result[size], result[size-1]
		}

		if result[i] > nums[j] {
			result[i] = nums[j]
			preMemory = memory
			memory = nums[j] // а вообще предыдущее значением
			i--
			j++
		}
		tmp = max(result[i], preMemory)
	}
	sort.Ints(result)
	size := len(result) - 1
	if size > 0 {
		tmp = result[size-1]
	} else {
		tmp = result[size]
	}
	return result[size], tmp
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size := len(nums1) + len(nums2)
	var last, preLast int
	pointMedian := int((size)/2 + 1)

	if len(nums1) > len(nums2) {
		resultSlice := make([]int, min(pointMedian, len(nums1)), pointMedian)
		copy(resultSlice, nums1[0:])
		preLast, last = changeResult(nums2, resultSlice)

	} else {

		resultSlice := make([]int, min(pointMedian, len(nums2)), pointMedian)
		copy(resultSlice, nums2[0:])
		preLast, last = changeResult(nums1, resultSlice)

	}
	if size%2 == 0 {
		return float64(preLast+last) / 2
	}
	return float64(max(preLast, last))
}

// func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
// 	size := len(nums1) + len(nums2)
// 	pointMedian := int((size)/2 + 1)

// 	tmp := make([]int, pointMedian)
// 	copy(tmp, nums1[0:])

// 	return 0.0
// }

// func main() {
// 	nums1 := []int{}
// 	nums2 := []int{}

// 	nums1 = []int{}
// 	nums2 = []int{2}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 2)

// 	nums1 = []int{}
// 	nums2 = []int{2, 3}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 2.5)

// 	nums1 = []int{1, 2}
// 	nums2 = []int{-1, 3}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 1.5)

// 	nums2 = []int{1, 2}
// 	nums1 = []int{-1, 3}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 1.5)

// 	nums1 = []int{}
// 	nums2 = []int{2, 5, 7, 8}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 6)

// 	nums1 = []int{0, 0}
// 	nums2 = []int{0, 0}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 0)

// 	nums1 = []int{1, 3}
// 	nums2 = []int{1, 3}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 2)

// 	nums1 = []int{1, 2}
// 	nums2 = []int{3, 4}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 2.5)

// 	nums1 = []int{1, 3}
// 	nums2 = []int{2}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 2)

// 	nums2 = []int{1, 2, 7, 11, 12}
// 	nums1 = []int{1, 11}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 7)

// 	nums1 = []int{4}
// 	nums2 = []int{2}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 3)

// 	nums1 = []int{1, 2, 7, 11, 12, 15, 16}
// 	nums2 = []int{5, 8}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 8)

// 	nums1 = []int{11, 12, 13, 14, 15}
// 	nums2 = []int{5, 8}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 12)

// 	nums1 = []int{1, 2, 5, 11, 12}
// 	nums2 = []int{7, 8}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 7)

// 	nums2 = []int{1, 2, 7, 11, 12}
// 	nums1 = []int{5, 8}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 7)

// 	nums2 = []int{1, 2, 7, 11, 12, 16}
// 	nums1 = []int{5, 9}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 8)

// 	nums1 = []int{1, 2, 7, 11, 12, 15}
// 	nums2 = []int{5}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 7)

// 	nums1 = []int{1, 2, 7, 11, 12}
// 	nums2 = []int{}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 7)

// 	nums1 = []int{1, 2, 7, 8, 12}
// 	nums2 = []int{15, 18}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 8)

// 	nums1 = []int{1, 2, 3}
// 	nums2 = []int{1, 2, 2}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 2)

// 	nums1 = []int{1, 2, 2}
// 	nums2 = []int{1, 2, 3}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 2)

// 	nums1 = []int{2, 2, 4, 4}
// 	nums2 = []int{2, 2, 4, 4}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 3)

// 	nums1 = []int{3, 4}
// 	nums2 = []int{1, 2, 5, 6}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 3.5)

// 	nums1 = []int{3}
// 	nums2 = []int{-2, -1}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), -1)

// 	nums1 = []int{1}
// 	nums2 = []int{2, 3, 4}
// 	fmt.Println("-------- ", findMedianSortedArrays(nums1, nums2), 2.5)

// }
