// 给定一个二叉树的根节点 root ，返回它的 中序 遍历。
// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
//
// 输入：root = [1]
// 输出：[1]
//
// 输入：root = [1,null,2]
// 输出：[1,2]
//
// 解题思路:按照访问左子树——根节点——右子树的方式遍历这棵树，而在访问左子树或者右子树的时候我们按照同样的方式遍历，直到遍历完整棵树。
// 因此可以直接用递归函数来模拟。
// 时间复杂度: O(N)
// 空间复杂度: O(N)
package main

func InorderTraversal(root *TreeNode) []int {
	var inorder func(node *TreeNode)
	vals := []int{}
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		vals = append(vals, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return vals
}