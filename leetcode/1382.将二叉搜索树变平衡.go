/*
 * @lc app=leetcode.cn id=1382 lang=golang
 *
 * [1382] 将二叉搜索树变平衡
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
func balanceBST(root *TreeNode) *TreeNode {

	//return balanceByRotate(root)
	return balanceByRebuild(root)
}

func balanceByRebuild(root *TreeNode) *TreeNode {

	inorder := getInorder(root)
	return buildBST(inorder)
}

func buildBST(inorder []int) *TreeNode {

	n := len(inorder)
	if n == 0 {
		return nil
	}

	v := &TreeNode{Val: inorder[n/2]}
	v.Left = buildBST(inorder[:n/2])
	v.Right = buildBST(inorder[n/2+1:])
	return v
}

func getInorder(root *TreeNode) []int {
	inorder := make([]int, 0)
	var getInorder0 func(root *TreeNode)
	getInorder0 = func(root *TreeNode) {

		if root == nil {
			return
		}
		getInorder0(root.Left)
		inorder = append(inorder, root.Val)
		getInorder0(root.Right)
	}

	getInorder0(root)

	return inorder
}

func balanceByRotate(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	r := true
	for r {
		r = false
		root.Left, root.Right = balanceBST(root.Left), balanceBST(root.Right)

		lh, lb := getHighAndBlance(root.Left)
		rh, rb := getHighAndBlance(root.Right)

		if lh-rh > 1 {
			if lb >= 0 {
				root = rightRotate(root)
			} else {
				root.Left = leftRotate(root.Left)
				root = rightRotate(root)
			}
			r = true
		}

		if lh-rh < -1 {
			if rb <= 0 {
				root = leftRotate(root)
			} else {
				root.Right = rightRotate(root.Right)
				root = leftRotate(root)
			}
			r = true
		}

	}

	return root
}

func leftRotate(root *TreeNode) *TreeNode {

	r := root.Right
	root.Right, r.Left = r.Left, root

	return r
}

func rightRotate(root *TreeNode) *TreeNode {

	l := root.Left
	root.Left, l.Right = l.Right, root

	return l
}

func getHighAndBlance(root *TreeNode) (int, int) {

	if root == nil {
		return 0, 0
	}
	l, _ := getHighAndBlance(root.Left)
	r, _ := getHighAndBlance(root.Right)
	return max(l, r) + 1, l - r
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

