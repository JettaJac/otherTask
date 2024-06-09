// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

package main

import (
	"fmt"
)

type ListNode struct {
	value int
	Next  *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	memory := 0
	res := 0
	i := 0
	for l1 != nil || l2 != nil || memory != 0 {
		fmt.Println(i)
		i++

		res = l1.value + l2.value + memory
		if res >= 10 {
			memory = 1
			result.value = res % 10
		}

		result.Next = &ListNode{value: res}
		fmt.Println("___")
		l1 = l1.Next
		l2 = l2.Next
	}

	return result
}

// func main() {
// 	l1 := &ListNode{value: 9}
// 	l1.Next = &ListNode{value: 9}
// 	l1.Next.Next = &ListNode{value: 9}
// 	l1.Next.Next.Next = &ListNode{value: 9}
// 	l1.Next.Next.Next.Next = &ListNode{value: 9}
// 	l1.Next.Next.Next.Next.Next = &ListNode{value: 9}
// 	l1.Next.Next.Next.Next.Next.Next = &ListNode{value: 9}

// 	l2 := &ListNode{value: 9}
// 	l2.Next = &ListNode{value: 9}
// 	l2.Next.Next = &ListNode{value: 9}
// 	l2.Next.Next.Next = &ListNode{value: 9}
// 	l2.Next.Next.Next = &ListNode{value: 9}

// 	result := addTwoNumbers(l1, l2)
// 	fmt.Println(result)
// }
