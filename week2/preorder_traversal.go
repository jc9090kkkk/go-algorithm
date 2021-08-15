// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
// 输入：root = [1,null,2,3]
// 输出：[1,2,3]
//
// 输入：root = []
// 输出：[]
//
// 输入：root = [1]
// 输出：[1]
// 解题思路:按照访问根节点——左子树——右子树的方式遍历这棵树，而在访问左子树或者右子树的时候，我们按照同样的方式遍历，直到遍历完整棵树。
// 所以树的遍历具有递归的性质，可以直接用递归函数来模拟这一过程。
// 时间复杂度:O(N)
// 空间复杂度:O(N)

package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {
	var preorder func(*TreeNode)
	vals := []int{}
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return vals
}

