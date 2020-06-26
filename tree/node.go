package tree

import "algorithm/my_stack"

// 树节点
type TreeNode struct {
	val    int64
	lChild *TreeNode
	rChild *TreeNode
}

// 按层次生成二叉树
func NewBinaryTree(arr []int64) *TreeNode {
	var head *TreeNode{}
	my_stack.N

	for idx := range arr {
		tmpTreeNode := &TreeNode{
			val: arr[idx],
		}
		if idx == 0 {
			head = tmpTreeNode
		}
		treeStack.
	}
}
