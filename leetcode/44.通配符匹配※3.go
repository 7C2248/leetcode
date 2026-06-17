/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	ls,lp := len(s),len(p)

	dp := make([]bool,lp+1)
	dp[0] = true
	for i:=1;i<=lp && p[i-1]=='*';i++{
		dp[i]=true
	}
	for i:=1;i<=ls;i++{
		prev := dp[0]
		for j:=1;j<=lp;j++{
			temp := dp[j]
			
			if p[j-1] == '*' {
				// '*'号匹配时，可以不消耗这个符号，从而从上侧转移
				dp[j] = dp[j-1] || dp[j]
			}else{
				dp[j] = (s[i-1]==p[j-1] || p[j-1]=='?') && prev
			}
			prev = temp
		}
		dp[0] = false
	}

	return dp[lp] 
}
// @lc code=end

zacabz
*a?b*

[1,0,0,0,0,0] 
[1,1,0,0,0,0] z
[1,1,1,0,0,0] a
[1,1,0,1,0,0] c
[1,1,1,0,0,0] a
[1,1,0,1,0,0] b
[1,1,0,0,0,0] z

zacakbz

[1,0,0,0,0,0] 
[1,1,0,0,0,0] z
[1,1,1,0,0,0] a
[1,1,0,1,0,0] c
[1,1,1,0,0,0] a
[1,1,0,1,0,0] k
[1,1,0,0,1,1] b
[1,1,0,0,0,1] z


