package main

import (
	"strconv"
	"strings"
)

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var ret []string
	var dfs func(head *TreeNode)
	dfs = func(head *TreeNode) {
		if head == nil {
			ret = append(ret, "x")
			return
		}
		ret = append(ret, strconv.Itoa(head.Val))
		dfs(head.Left)
		dfs(head.Right)
	}
	dfs(root)

	return strings.Join(ret, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	index := -1
	arr := strings.Split(data, ",")
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		index++
		if index >= len(data) || arr[index] == "x" {

			return nil
		}
		val, _ := strconv.Atoi(arr[index])

		head := &TreeNode{Val: val}
		head.Left = dfs()
		head.Right = dfs()
		return head
	}
	return dfs()
}
