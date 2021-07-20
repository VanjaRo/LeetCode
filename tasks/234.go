package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	fmt.Print(isPalindrome(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	l := 1
	curr := head
	for curr.Next != nil {
		curr = curr.Next
		l++
	}

	curr = head
	for i := 0; i < l/2; i++ {
		curr = curr.Next
	}
	var prev *ListNode
	for curr.Next != nil {
		curr, curr.Next, prev = curr.Next, prev, curr
	}
	curr.Next = prev

	for i := 0; i < l/2; i++ {
		if head.Val != curr.Val {
			return false
		}
		curr, head = curr.Next, head.Next
	}
	return true
}
