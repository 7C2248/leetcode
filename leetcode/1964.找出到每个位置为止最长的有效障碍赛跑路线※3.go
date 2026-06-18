/*
 * @lc app=leetcode.cn id=1964 lang=golang
 *
 * [1964] 找出到每个位置为止最长的有效障碍赛跑路线
 */

// @lc code=start

// 由于长度n满足1 <= n <= 10^5，而值域v满足1 <= v <= 10^7
// 从而 二分搜索更优，二分搜索和树状数组都是O(nlogn)的时间复杂度。

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	// tails保存长度为i时结尾的最大值
	// 最长为n
	tails := make([]int, 0)
	ans := make([]int, len(obstacles))
	for i, x := range obstacles {
		// 二分查找，找到第一个 > x 的位置
		pos := sort.Search(len(tails), func(j int) bool { return tails[j] > x })
		if pos == len(tails) {
			tails = append(tails, x)
		} else {
			tails[pos] = x
		}
		ans[i] = pos + 1
	}
	return ans
}

// @lc code=end

/*
// 离散状态 树状数组(Fenwick Tree)优化
type Fenwick struct {
    tree []int
    n    int
}

func NewFenwick(n int) *Fenwick {
    return &Fenwick{tree: make([]int, n+1), n: n}
}

func (f *Fenwick) Update(i, val int) {
    i++ // 树状数组下标从 1 开始
    for i <= f.n {
        if val > f.tree[i] {
            f.tree[i] = val
        }
        i += i & -i
    }
}

func (f *Fenwick) Query(i int) int {
    i++ // 查询前缀 [0..i] 的最大值
    res := 0
    for i > 0 {
        if f.tree[i] > res {
            res = f.tree[i]
        }
        i -= i & -i
    }
    return res
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
    // 离散化
    set := map[int]struct{}{}
    for _, v := range obstacles {
        set[v] = struct{}{}
    }
    // 最长为max(obstacles[i])
    sorted := make([]int, 0, len(set))
    for k := range set {
        sorted = append(sorted, k)
    }
    sort.Ints(sorted)
    rank := make(map[int]int, len(sorted))
    for i, v := range sorted {
        rank[v] = i
    }

    ft := NewFenwick(len(sorted))
    res := make([]int, len(obstacles))
    for i, v := range obstacles {
        idx := rank[v]
        best := ft.Query(idx) + 1
        ft.Update(idx, best)
        res[i] = best
    }
    return res
}
*/