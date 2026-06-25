/*
 * @lc app=leetcode.cn id=2430 lang=golang
 *
 * [2430] 对字母串可执行的最大删除数
 */

// @lc code=start
func deleteString(s string) int {
    n := len(s)
    // 预处理 LCP（用 int16 节省空间）
    // 表示s[i:]和s[j:]的公共前缀长
    lcp := make([][]int16, n+1)
    for i := range lcp {
        lcp[i] = make([]int16, n+1)
    }
    // 利用LCP构建O(1)查询长度为2k的字符串是否前后部分相等
    for i := n - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if s[i] == s[j] {
                lcp[i][j] = 1 + lcp[i+1][j+1]
            }
        }
    }

    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = -1
    }
    dp[0] = 0

    for i := 0; i < n; i++ {
        if dp[i] == -1 {
            continue
        }
        // 操作 1：全删
        if dp[i]+1 > dp[n] {
            dp[n] = dp[i] + 1
        }
        // 操作 2：删满足条件的前缀 k
        maxK := (n - i) / 2
        for k := 1; k <= maxK; k++ {
            if lcp[i][i+k] >= int16(k) {
                if dp[i]+1 > dp[i+k] {
                    dp[i+k] = dp[i] + 1
                }
            }
        }
    }
    fmt.Println(dp)
    return dp[n]
}
// @lc code=end

aaabaab
[0,-1,-1,-1,-1,-1,-1,-1]
[0, 1, 1,-1,-1,-1,-1, 1]
[0, 1, 2,-1, 2,-1,-1, 2]
[0, 1, 2,-1, 2, 3,-1, 3]
[0, 1, 2,-1, 2, 3,-1, 4]