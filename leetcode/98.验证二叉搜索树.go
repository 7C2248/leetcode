/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
import (
	"math"
)

func isValidBST(root *TreeNode) bool {

	//_, _, result := getBSTpost(root)
	//result := getBSTpre(root, math.MinInt, math.MaxInt)
	result := getBSTin(root)
	return result
}

//后序遍历
func getBSTpost(root *TreeNode) (int, int, bool) {

	if root == nil {
		return math.MaxInt, math.MinInt, true
	}

	v := root.Val

	lmin, lmax, lb := getBSTpost(root.Left)
	rmin, rmax, rb := getBSTpost(root.Right)

	if !lb || !rb || lmax >= v || rmin <= v {
		return math.MinInt, math.MaxInt, false
	}

	return min(lmin, v), max(rmax, v), true
}

//前序遍历
func getBSTpre(root *TreeNode, min int, max int) bool {

	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return getBSTpre(root.Left, min, root.Val) && getBSTpre(root.Right, root.Val, max)

}

//中序遍历
func getBSTin(root *TreeNode) bool {

	pre := math.MinInt
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if !dfs(node.Left) { // 左
			return false
		}
		if node.Val <= pre { // 中
			return false
		}
		pre = node.Val
		return dfs(node.Right) // 右
	}
	return dfs(root)

}

// @lc code=end

