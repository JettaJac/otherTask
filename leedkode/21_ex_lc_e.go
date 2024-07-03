/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode { // done, but slowly. need to do it through recursion
	head := &ListNode{}
	result := head
	res := 0

	for list1 != nil || list2 != nil {
		// newNode := &ListNode{}
		fmt.Println("_____")
		if list1 == nil {
			res = list2.Val
			list2 = list2.Next
		} else if list2 == nil {
			res = list1.Val
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			fmt.Println("l1<")
			res = list1.Val
			// fmt.Println(list1.Val, res)
			list1 = list1.Next

		} else if list1.Val >= list2.Val {
			fmt.Println("l1>")
			res = list2.Val
			list2 = list2.Next
		}
		newNode := &ListNode{Val: res}
		result.Next = newNode
		result = newNode

	}
	// fmt.Println("uuuu", head.Next.Val)

	return head.Next
}

func main() {
	l1 := &ListNode{Val: 1}
	// l1.Next = &ListNode{Val: 2}
	// l1.Next.Next = &ListNode{Val: 4}
	// l1.Next.Next.Next = &ListNode{Val: 9}
	// l1.Next.Next.Next.Next = &ListNode{Val: 9}
	// l1.Next.Next.Next.Next.Next = &ListNode{Val: 9}
	// l1.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}

	l2 := &ListNode{Val: 2}
	// l2.Next = &ListNode{Val: 3}
	// l2.Next.Next = &ListNode{Val: 4}
	// l2.Next.Next.Next = &ListNode{Val: 9}
	// l2.Next.Next.Next = &ListNode{Val: 9}

	res := mergeTwoLists(l1, l2)
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next
	}

}
