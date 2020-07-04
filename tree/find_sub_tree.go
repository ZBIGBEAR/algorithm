package tree

// 判断一棵树是否为另外一颗树的子树
func IsSubTree(dst *TreeNode, src *TreeNode) bool {
	// 空树是任意树的子树
	if src == nil {
		return true
	}

	if dst == nil {
		return false
	}

	if dst.val == src.val {
		if IsSubTreeWithHead(dst, src) {
			return true
		}
	}
	return IsSubTree(dst.lChild, src) || IsSubTree(dst.rChild, src)
}

// 判断从根节点开始，一棵树是否为另外一颗树的子树
func IsSubTreeWithHead(dst *TreeNode, src *TreeNode) bool {
	if src == nil {
		return true
	}
	if dst == nil {
		return false
	}

	if dst.val != src.val {
		return false
	}
	return IsSubTreeWithHead(dst.lChild, src.lChild) && 
			IsSubTreeWithHead(dst.rChild, src.rChild)
}