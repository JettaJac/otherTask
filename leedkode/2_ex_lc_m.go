// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

package main

import (
// "fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{}
	result := head
	memory := 0

	for l1 != nil || l2 != nil || memory != 0 {
		sum := memory
		memory = 0
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		memory = sum / 10

		newNode := &ListNode{Val: sum % 10}
		result.Next = newNode
		result = newNode

		sum = 0
	}

	return head.Next
}

// func main() {
// 	l1 := &ListNode{Val: 9}
// 	l1.Next = &ListNode{Val: 9}
// 	l1.Next.Next = &ListNode{Val: 9}
// 	l1.Next.Next.Next = &ListNode{Val: 9}
// 	l1.Next.Next.Next.Next = &ListNode{Val: 9}
// 	l1.Next.Next.Next.Next.Next = &ListNode{Val: 9}
// 	l1.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}

// 	l2 := &ListNode{Val: 9}
// 	l2.Next = &ListNode{Val: 9}
// 	l2.Next.Next = &ListNode{Val: 9}
// 	l2.Next.Next.Next = &ListNode{Val: 9}
// 	l2.Next.Next.Next = &ListNode{Val: 9}

// 	result := addTwoNumbers(l1, l2)
// 	fmt.Println("___3: ", result.Val)
// 	l2 = result
// 	for l2 != nil {
// 		fmt.Print(l2.Val)
// 		l2 = l2.Next
// 	}
// }
