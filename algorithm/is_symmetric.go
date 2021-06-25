package algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
	return symmetricRecur(root.Left, root.Right)
}

func symmetricRecur(lSide, rSide *TreeNode) bool {
	var out, in bool
	if lSide.Val != rSide.Val {
		return false
	}
	if lSide.Left != nil && rSide.Right != nil {
		out = symmetricRecur(lSide.Left, rSide.Right)
	} else if lSide.Left == nil && rSide.Right == nil {
		out = true
	}
	if lSide.Right != nil && rSide.Left != nil {
		in = symmetricRecur(lSide.Right, rSide.Left)
	} else if lSide.Right == nil && rSide.Left == nil {
		in = true
	}
	return out && in
}
