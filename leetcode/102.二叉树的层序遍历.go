/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)

	type queueItem struct {
		node  *TreeNode
		level int
	}

	queue := make([]queueItem, 0)
	queue = append(queue, queueItem{root, 1})

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]
		node, level := current.node, current.level

		if level > len(result) {
			result = append(result, make([]int, 0))
		}
		result[level-1] = append(result[level-1], node.Val)

		if node.Left != nil {
			queue = append(queue, queueItem{node.Left, level + 1})
		}
		if node.Right != nil {
			queue = append(queue, queueItem{node.Right, level + 1})
		}
	}
	return result
}

// @lc code=end

