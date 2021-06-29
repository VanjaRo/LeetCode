package main

import "fmt"

func main() {
	t := 10
	fmt.Print(t / 3)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{
		Val: 0,
	}
	left := 0
	retHead := ret
	for l1 != nil || l2 != nil {
		sum := left
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		tmp := &ListNode{
			Val: sum % 10,
		}
		ret.Next = tmp
		ret = ret.Next
		left = sum / 10
	}
	if left == 1 {
		ret.Next = &ListNode{
			Val: 1,
		}
	}

	return retHead.Next
}
