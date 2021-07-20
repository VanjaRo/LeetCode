package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ret := recursiveBackCounterDelete(head, n)
	if ret == -1 {
		return head.Next
	}
	return head
}

// 0 – found
// -1 – delete next
func recursiveBackCounterDelete(head *ListNode, n int) int {
	if head == nil {
		return 1
	}
	ret := recursiveBackCounterDelete(head.Next, n)
	if ret == 0 {
		return 0
	}
	if ret == -1 {
		head.Next = head.Next.Next
		return 0
	}
	if ret == n {
		return -1
	}
	return ret++
}

