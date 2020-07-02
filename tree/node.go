package tree

import (
	myStack "algorithm/my_stack"
	"fmt"
)

// 树节点
type TreeNode struct {
	val    int64
	lChild *TreeNode
	rChild *TreeNode
}

// 按层次生成二叉树
func NewBinaryTree(arr []int64) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return &TreeNode{
			val: arr[0],
		}
	}

	head := &TreeNode{
		val: arr[0],
	}

	count := len(arr)

	// 定义一个栈
	stack := myStack.NewStack()
	// 第一个节点进栈
	stack.Push(head)
	idx := 1
	for {
		// 出栈
		popNode := stack.Pop()
		// 从数组中取出2个元素，分别作为刚出栈节点的左右孩子
		popNode.lChild = &TreeNode{
			val: arr[idx],
		}
		idx += 1
		if idx >= count {
			break
		}
		popNode.rChild = &TreeNode{
			val: arr[idx],
		}
		idx += 1
		if idx >= count {
			break
		}

		// 左右孩子入栈
		stack.Push(popNode.lChild)
		stack.Push(popNode.rChild)
	}

	// 返回二叉树根节点
	return head
}

// 递归输出二叉树的前序遍历
func PrintBinaryTree(head *TreeNode) {
	if head == nil {
		return
	}
	fmt.Print(head.val, "\n")
	PrintBinaryTree(head.lChild)
	PrintBinaryTree(headl.rChild)
}

// 非递归输出二叉树的前序遍历
