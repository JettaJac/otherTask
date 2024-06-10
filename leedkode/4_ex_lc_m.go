// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).

// Example 1:
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size := len(nums1) + len(nums2)
	pointMedian := int((size)/2 + 1)
	resultSlice := make([]int, pointMedian)
	z := 0
	// for z := 0; z < pointMedian+1; z++ {
	// for i := 0; i < len(nums1); {
	for i, j := 0, 0; j < len(nums2) && i < len(nums1) && z < pointMedian; {

		fmt.Println(resultSlice, " ", nums1[i], nums2[j])
		if nums1[i] <= nums2[j] {
			resultSlice[z] = nums1[i]
			i++
		} else {
			resultSlice[z] = nums2[j]
			j++
		}
		z++

		fmt.Println("______")
	}
	// }

	// }
	if size%2 == 0 {
		return float64(resultSlice[pointMedian-1]+resultSlice[pointMedian-2]) / 2
	}
	return float64(resultSlice[pointMedian-1])
}

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	size := len(nums1) + len(nums2)
// 	pointMedian := int((size)/2 + 1)
// 	// tmp := make([]int, size/2+1)
// 	// tmp := nums1[0:pointMedian]
// 	tmp := make([]int, pointMedian)
// 	copy(tmp, nums1[0:])
// 	fmt.Println(len(tmp), tmp)
// 	_ = tmp
// 	// result := 0.0
// 	// i := 0
// 	// fmt.Println(pointMedian)
// 	// memory := 0

// 	// for j := pointMedian - 1; j > 0; j-- {
// 	// 	fmt.Println("W___")
// 	// 	if tmp[j] > nums2[i] {
// 	// 		fmt.Println("+++++")
// 	// 		tmp[j] = nums2[i]
// 	// 		memory = nums2[i]
// 	// 		fmt.Println("+++++/")

// 	// 	}
// 	// 	fmt.Println("____", len(nums2), nums2[i+1])
// 	// 	if i+1 > len(nums2) || tmp[j-1] <= nums2[i+1] {
// 	// 		fmt.Println("____")
// 	// 		// fmt.Println(tmp, " ", tmp[j], nums2[i])

// 	// 		fmt.Println(tmp, " ", tmp[j], nums2[i], "memory", memory)
// 	// 		if size%2 == 0 {
// 	// 			result = float64(max(tmp[j], memory)+max(tmp[j-1], memory)) / 2

// 	// 		} else {

// 	// 			result = float64(max(tmp[j], memory))
// 	// 		}
// 	// 		return result
// 	// 	}
// 	// 	i++

// 	// }

// 	// if pointMedian%2 == 0 {
// 	// 	pointMedian = size/2 + 1
// 	// } else {

// 	// }

// 	// for i := 0; i < size/2+1; i++ {
// 	// 	if nums1[i] < nums2[i] {
// 	// 		tmp = append(tmp, nums1[i])
// 	// 	}

// 	// }

// 	return 0.0
// }

func main() {
	nums1 := []int{}
	nums2 := []int{}

	// nums1 = []int{1, 2, 7, 11, 12, 15, 16}
	// nums2 = []int{5, 8}
	// fmt.Println(findMedianSortedArrays(nums1, nums2), 8)

	// nums1 = []int{1, 2, 5, 11, 12}
	// nums2 = []int{7, 8}
	// fmt.Println(findMedianSortedArrays(nums1, nums2), 7)

	// nums2 = []int{1, 2, 7, 11, 12}
	// nums1 = []int{5, 8}
	// fmt.Println(findMedianSortedArrays(nums1, nums2), 7)

	// nums2 = []int{1, 2, 7, 11, 12, 16}
	// nums1 = []int{5, 9}
	// fmt.Println(findMedianSortedArrays(nums1, nums2), 8)

	nums1 = []int{1, 2, 7, 11, 12, 15}
	nums2 = []int{5}
	fmt.Println(findMedianSortedArrays(nums1, nums2), 7)

	// nums1 = []int{1, 2, 7, 11, 12}
	// nums2 = []int{}
	// fmt.Println(findMedianSortedArrays(nums1, nums2), 7)

	// nums1 = []int{1, 2, 7, 8, 12}
	// nums2 = []int{15, 18}
	// fmt.Println(findMedianSortedArrays(nums1, nums2), 8)

}
