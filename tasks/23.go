package main

import "fmt"

func main() {

}

//  * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	step, lenLists := 1, len(lists)
	if lenLists == 0 {
		return nil
	}
	for step < lenLists {
		for i := 0; i+step < lenLists; i += 2 * step {
			lists[i] = mergeTwoLists(lists[i], lists[i+step])
		}
		step *= 2
	}
	return lists[0]
}

func (l *ListNode) Print() {
	tmp := l
	for tmp != nil {
		fmt.Print(tmp.Val)
		tmp = tmp.Next
	}
}

func mergeTwoLists(fList *ListNode, sList *ListNode) *ListNode {
	// var min *ListNode
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
