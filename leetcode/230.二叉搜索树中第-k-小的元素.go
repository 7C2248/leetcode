/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第 K 小的元素
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
func kthSmallest(root *TreeNode, k int) int {

	stack := make([]*TreeNode, 0)

	for {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) <= 0 {
			break
		}

		stack, root = stack[:len(stack)-1], stack[len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right

	}
	return -1

	//return dfs_kth(root, k)
}

/*
//递归开销过大
func dfs_kth(root *TreeNode, k int) int {

	var dfs func(root *TreeNode) bool
	var result int
	i := 0
	dfs = func(root *TreeNode) bool {

		if root == nil {
			return true
		}

		if dfs(root.Left) {
			i++

			if i == k {
				result = root.Val
				return false
			}
			return dfs(root.Right)
		}
		return false
	}
	if !dfs(root) {
		return result
	} else {
		return -1
	}
}
*/
// @lc code=end

