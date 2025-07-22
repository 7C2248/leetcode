/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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
func generateTree(Val []int) []*TreeNode {

	result := []*TreeNode{}
	if len(Val) == 0 {
		return append(result, nil)
	}
	if len(Val) == 1 {
		return append(result, &TreeNode{Val: Val[0]})
	}

	for i := 0; i < len(Val); i++ {
		small := append([]int{}, Val[:i]...)
		big := append([]int{}, Val[i+1:]...)
		val := Val[i]
		lefttrees := generateTree(small)
		righttrees := generateTree(big)
		if (len(lefttrees) > 0) && (len(righttrees) > 0) {
			for j := range lefttrees {
				for k := range righttrees {
					node := &TreeNode{Val: val, Left: lefttrees[j], Right: righttrees[k]}
					result = append(result, node)
				}
			}
		} else {
			if len(lefttrees) > 0 {
				for j := range lefttrees {
					node := &TreeNode{Val: val, Left: lefttrees[j]}
					result = append(result, node)
				}
			} else {
				for j := range righttrees {
					node := &TreeNode{Val: val, Right: righttrees[j]}
					result = append(result, node)
				}
			}
		}
	}
	return result
}

func generateTrees(n int) []*TreeNode {

	a := make([]int, n)
	for i := 1; i <= n; i++ {
		a[i-1] = i
	}
	return generateTree(a)
}

// @lc code=end

