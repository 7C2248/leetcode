/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
/*
 // dfs
 func rightSideView(root *TreeNode) []int {

    result:=make([]int,0)
    var dfs func(level int, node *TreeNode)
    dfs = func(level int, node *TreeNode){
        if node == nil{
            return
        }
        if len(result) < level{
            result = append(result,node.Val)
        }
        dfs(level+1,node.Right)
        dfs(level+1,node.Left)
        return
    }
    dfs(1,root)
    return result
}
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)

	queue0 := make([]*TreeNode, 0)
	queue0 = append(queue0, root)
	level := 1
	for len(queue0) > 0 {

		queue1 := make([]*TreeNode, 0)
		for i := 0; i < len(queue0); i++ {
			node := queue0[i]

			if len(result) < level {
				result = append(result, node.Val)
			}

			if node.Right != nil {
				queue1 = append(queue1, node.Right)
			}
			if node.Left != nil {
				queue1 = append(queue1, node.Left)
			}

		}
		queue0 = queue1
		level++

	}
	return result
}

// @lc code=end

