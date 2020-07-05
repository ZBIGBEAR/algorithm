package tree

import (
	myStack "algorithm/my_stack"
	"fmt"
)

// 处理二叉树

// 1.遍历二叉树
// 1.1 前序遍历
// 1.1.1 递归
func PreOrderTraverseBinaryTree(head *TreeNode) {
	if head == nil {
		return
	}

	// 先遍历根节点
	fmt.Printf("%d\t", head.val)
	// 递归遍历左子树
	if head.lChild != nil {
		PreOrderTraverseBinaryTree(head.lChild)
	}
	// 递归遍历右子树
	if head.rChild != nil {
		PreOrderTraverseBinaryTree(head.rChild)
	}
}

// 1.1.2 循环 方法一
func PreOrderTraverseBinaryTreeWithLoop1(head *TreeNode) {
	if head == nil {
		return
	}
	// 定义一个栈
	stack := myStack.NewStack()
	// 根节点入栈
	stack.Push(head)
	var node *TreeNode
	var ok bool
	for {
		// 出栈
		node, ok = stack.Pop().(*TreeNode)
		if !ok || node == nil {
			break
		}

		// 输出栈顶元素
		fmt.Printf("%d\t", node.val)
		// 如果有右孩子，则进栈
		if node.rChild != nil {
			stack.Push(node.rChild)
		}
		// 如果有左孩子，则进栈
		if node.lChild != nil {
			stack.Push(node.lChild)
		}
	}
	return
}

// 1.1.3 循环 方法二
func PreOrderTraverseBinaryTreeWithLoop2(head *TreeNode) {
	if head == nil {
		return
	}
	// 定义一个栈
	stack := myStack.NewStack()
	lChild := head
	for {
		if lChild != nil {
			fmt.Printf("%d\t", lChild.val)
			stack.Push(lChild)
			lChild = lChild.lChild
		} else {
			break
		}
	}
	var node *TreeNode
	var ok bool
	for {
		// 出栈
		node, ok = stack.Pop().(*TreeNode)
		if !ok || node == nil {
			break
		}
		// 访问右子树，如果存在，则输出，进栈
		rChild := node.rChild
		// 右子树存在，则输出右子树，并进栈，且它的左节点进栈
		lChild = rChild
		for {
			if lChild != nil {
				fmt.Printf("%d\t", lChild.val)
				stack.Push(lChild)
				lChild = lChild.lChild
			} else {
				break
			}
		}
	}
	return
}

// 1.2 中序遍历
// 1.2.1 递归
func MidOrderTraverseBinaryTree(head *TreeNode) {
	if head == nil {
		return
	}
	// 输出左子树
	if head.lChild != nil {
		MidOrderTraverseBinaryTree(head.lChild)
	}
	// 输出根节点
	fmt.Printf("%d\t", head.val)
	// 输出右子树
	if head.rChild != nil {
		MidOrderTraverseBinaryTree(head.rChild)
	}
}

// 1.2.2 循环
func MidOrderTraverseBinaryTreeWithLoop(head *TreeNode) {
	if head == nil {
		return
	}
	// 定义一个栈
	stack := myStack.NewStack()
	lChild := head
	for {
		if lChild != nil {
			stack.Push(lChild)
			lChild = lChild.lChild
		} else {
			break
		}
	}
	var node *TreeNode
	var ok bool
	for {
		// 出栈
		node, ok = stack.Pop().(*TreeNode)
		if !ok || node == nil {
			break
		}
		fmt.Printf("%d\t", node.val)
		// 访问右子树，如果存在，则输出，进栈
		rChild := node.rChild
		// 右子树存在，则输出右子树，并进栈，且它的左节点进栈
		lChild = rChild
		for {
			if lChild != nil {
				stack.Push(lChild)
				lChild = lChild.lChild
			} else {
				break
			}
		}
	}
	return
}

// 1.3 后续遍历
// 1.3.1 递归
func PostOrderTraverseBinaryTree(head *TreeNode) {
	if head == nil {
		return
	}
	// 输出左子树
	if head.lChild != nil {
		PostOrderTraverseBinaryTree(head.lChild)
	}
	// 输出右子树
	if head.rChild != nil {
		PostOrderTraverseBinaryTree(head.rChild)
	}
	// 输出根节点
	fmt.Printf("%d\t", head.val)
}

// 1.3.2 循环
func PostOrderTraverseBinaryTreeWithLoop1(head *TreeNode) {
	// 1.如果当前节点是叶子节点，则直接输出
	// 2.如果有孩子，且孩子没有访问过则按照右孩子、左孩子依次入栈
	// 3.如果有孩子，且都访问过则访问P
	if head == nil {
		return
	}
	p := head
	cur := head
	stack := myStack.NewStack()
	stack.Push(p)
	for {
		node, ok := stack.Pop().(*TreeNode)
		if !ok {
			break
		}
		visit := false
		if node.lChild == nil && node.rChild == nil {
			// 当前节点是叶子节点，直接访问
			visit = true
		} else if cur == node.lChild && node.rChild == nil {
			// 当前节点有左孩子，且刚刚被访问过，且没有右孩子，则直接访问
			visit = true
		} else if cur == node.rChild {
			// 当前节点有右孩子，且刚刚被访问过，则直接访问
			visit = true
		}
		if visit {
			fmt.Printf("%d\t", node.val)
			cur = node
		} else {
			stack.Push(node)
			if node.rChild != nil {
				stack.Push(node.rChild)
			}
			if node.lChild != nil {
				stack.Push(node.lChild)
			}
		}
	}

}

// 1.3.3 后序遍历二叉树循环方法二
// 每个节点保存2遍到栈中，先进栈的那个节点用于访问当前节点，后面那个节点用于把它的左右子树保存进来。
// 出栈的时候，判断已出栈的节点与栈顶节点是否相同，如果相同则说明该节点的左右孩子还未处理；否则说明它
// 的左右孩子已经处理，直接访问当前节点即可
func PostOrderTraverseBinaryTreeWithLoop2(head *TreeNode) {
	if head == nil {
		return
	}
	stack := myStack.NewStack()
	stack.Push(head)
	stack.Push(head)
	for {
		if stack.IsEmpty() {
			break
		}
		node, ok := stack.Pop().(*TreeNode)
		if !ok {
			break
		}
		if node == stack.Top() {
			if node.rChild != nil {
				stack.Push(node.rChild)
				stack.Push(node.rChild)
			}
			if node.lChild != nil {
				stack.Push(node.lChild)
				stack.Push(node.lChild)
			}
		} else {
			fmt.Printf("%d\t", node.val)
		}
	}
}

// 1.4 层次遍历

// 2.将一颗普通的二叉树变成排序二叉树

// 3.从有序二叉树中查找某个元素
