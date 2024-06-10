// Done
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

package main

import (
// "fmt"
)

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	for i, elem := range nums {
		tmp := target - elem

		_, ok := myMap[elem]
		if ok {
			result := []int{myMap[elem], i}
			return result
		}
		myMap[tmp] = i
	}
	return nil
}

// func main() {
// 	nums := []int{1, 2, 7, 11, 5, 15}
// 	target := 9
// 	result := twoSum(nums, target)
// 	fmt.Println(result)
// }
