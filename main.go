package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// MaxPathSum returns the largest sum over all paths (any node to any node,
// downward-connected, at least one node). Each node's best bending path is
// its value plus both children's non-negative gains; what flows upward is
// the value plus the better single gain.
func MaxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	best := root.Val
	var gain func(n *TreeNode) int
	gain = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l := max(gain(n.Left), 0) // negative subtrees are simply not taken
		r := max(gain(n.Right), 0)
		if n.Val+l+r > best {
			best = n.Val + l + r // path bending through n
		}
		return n.Val + max(l, r) // only one arm can continue upward
	}
	gain(root)
	return best
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(-10, tn(9, nil, nil), tn(20, tn(15, nil, nil), tn(7, nil, nil)))
	fmt.Println(MaxPathSum(root))
}
