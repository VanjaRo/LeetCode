package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return areSymmetricNodes(root.Left, root.Right)
}

func areSymmetricNodes(lSubTree *TreeNode, rSubTree *TreeNode) bool {
	if lSubTree == nil && rSubTree == nil {
		return true
	}

	if lSubTree == nil || rSubTree == nil || lSubTree.Val != rSubTree.Val {
		return false
	}

	return areSymmetricNodes(lSubTree.Left, rSubTree.Right) && areSymmetricNodes(lSubTree.Right, rSubTree.Left)

}
