/*
 * @lc app=leetcode.cn id=1105 lang=golang
 *
 * [1105] 填充书架
 */

// @lc code=start
func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
	}

	for i := 0; i < n; i++ {
		curWidth := 0
		maxH := 0
		for j := i; j > -1; j-- {
			curWidth += books[j][0]
			maxH = max(maxH, books[j][1])
			if curWidth <= shelfWidth {
				dp[i+1] = min(dp[i+1], maxH+dp[j])
			} else {
				break
			}
		}
	}
	return dp[n]
}

// @lc code=end
