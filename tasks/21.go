package main

func main() {

}

//  * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(fList *ListNode, sList *ListNode) *ListNode {
	var head *ListNode
	var current *ListNode
	for fList != nil || sList != nil {
		min, isFirst := minValNode(fList, sList)
		if isFirst {
			fList = fList.Next
		} else {
			sList = sList.Next
		}
		if head == nil {
			head = min
			current = head
			continue
		}
		current.Next = min
		current = current.Next

	}
	return head
}

func minValNode(fNode, sNode *ListNode) (*ListNode, bool) {
	if sNode == nil {
		return fNode, true
	}
	if fNode == nil {
		return sNode, false
	}
	if sNode.Val < fNode.Val {
		return sNode, false
	}
	return fNode, true
}
