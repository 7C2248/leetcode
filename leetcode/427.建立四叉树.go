/*
 * @lc app=leetcode.cn id=427 lang=golang
 *
 * [427] 建立四叉树
 */

// @lc code=start
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	sum := buildPrefixSum(grid)
	return conHelper(grid, sum, 0, 0, len(grid))
}

func merge(tl, tr, bl, br *Node) *Node {
	if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf && tl.Val == tr.Val && tr.Val == bl.Val && bl.Val == br.Val {
		return &Node{Val: tl.Val, IsLeaf: true, TopLeft: nil, TopRight: nil, BottomLeft: nil, BottomRight: nil}
	}
	return &Node{Val: false, IsLeaf: false, TopLeft: tl, TopRight: tr, BottomLeft: bl, BottomRight: br}
}

func conHelper(grid [][]int, sum [][]int, x, y, levelLength int) *Node {

	total := rectSum(sum, x, y, x+levelLength-1, y+levelLength-1)

	if total == 0 {
		return &Node{Val: false, IsLeaf: true}
	}
	if total == levelLength*levelLength {
		return &Node{Val: true, IsLeaf: true}
	}

	if levelLength == 1 {
		return &Node{Val: grid[x][y] == 1, IsLeaf: true, TopLeft: nil, TopRight: nil, BottomLeft: nil, BottomRight: nil}
	}

	//注意行列索引！！
	//tl, tr, bl, br := conHelper(grid, x, y, levelLength/2), conHelper(grid, x+levelLength/2, y, levelLength/2), conHelper(grid, x, y+levelLength/2, levelLength/2), conHelper(grid, x+levelLength/2, y+levelLength/2, levelLength/2)
	tl, tr, bl, br := conHelper(grid, sum, x, y, levelLength/2), conHelper(grid, sum, x, y+levelLength/2, levelLength/2), conHelper(grid, sum, x+levelLength/2, y, levelLength/2), conHelper(grid, sum, x+levelLength/2, y+levelLength/2, levelLength/2)

	return merge(tl, tr, bl, br)
}

func buildPrefixSum(grid [][]int) [][]int {
	n := len(grid)
	sum := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		sum[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = grid[i-1][j-1] +
				sum[i-1][j] +
				sum[i][j-1] -
				sum[i-1][j-1]
		}
	}
	return sum
}

func rectSum(sum [][]int, x1, y1, x2, y2 int) int {
	return sum[x2+1][y2+1] - sum[x1][y2+1] - sum[x2+1][y1] + sum[x1][y1]
}

// @lc code=end

