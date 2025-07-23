/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {

	n := len(preorder)
	if n == 0 {
		return nil
	}
	root := TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	if n == 1 {
		return &root
	}
	i := 0
	for i < n && inorder[i] != preorder[0] {
		i++
	}
	root.Left, root.Right = buildTree(preorder[1:1+i], inorder[:i]), buildTree(preorder[i+1:], inorder[i+1:])
	return &root
}

// @lc code=end

