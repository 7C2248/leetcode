/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	if n == 1 {
		rr := TreeNode{
			Val:   postorder[n-1],
			Left:  nil,
			Right: nil,
		}
		return &rr
	}

	i := 0
	for i < n && inorder[i] != postorder[n-1] {
		i++
	}
	l := buildTree(inorder[:i], postorder[:i])
	r := buildTree(inorder[i+1:], postorder[i:n-1])

	rr := TreeNode{
		Val:   postorder[n-1],
		Left:  l,
		Right: r,
	}
	return &rr
}

// @lc code=end

