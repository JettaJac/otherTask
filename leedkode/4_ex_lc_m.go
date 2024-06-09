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
	tmp := make([]int, size)
	result := 0.0
	// if size%2 == 0 {
	// 	size = size/2 + 1
	// } else {

	// }

	for i := 0; i < size/2+1; i++ {
		if nums1[i] < nums2[i] {
			tmp = append(tmp, nums1[i])
		}

	}
	return result
}

func main() {
	nums1 := []int{1, 2, 7, 11, 12, 15}
	nums2 := []int{5}
	fmt.Println(findMedianSortedArrays(nums1, nums2), 5)

}
