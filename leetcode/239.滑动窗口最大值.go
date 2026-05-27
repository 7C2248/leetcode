/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
type node struct {
	val, index int
}
type heap struct {
	data []*node
}

func (h *heap) up(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if h.data[p].val >= h.data[i].val {
			break
		}
		h.data[p], h.data[i] = h.data[i], h.data[p]
		i = p
	}
}
func (h *heap) down(i int) {
	n := len(h.data)
	for {
		left, right := i*2+1, i*2+2
		smallest := i
		if left < n && h.data[left].val > h.data[smallest].val {
			smallest = left
		}
		if right < n && h.data[right].val > h.data[smallest].val {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}
func (h *heap) push(n *node) {
	h.data = append(h.data, n)
	h.up(len(h.data) - 1)
}
func (h *heap) pop() *node {
	if len(h.data) == 0 {
		return nil
	}
	result := h.data[0]
	last := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	if len(h.data) > 0 {
		h.data[0] = last
		h.down(0)
	}
	return result
}

func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	n := len(nums)
	h := &heap{}
	for i := 0; i < k; i++ {
		r := &node{val: nums[i], index: i}
		h.push(r)
	}
	for j := k - 1; j < n; j++ {
		if j != k-1 {
			r := &node{val: nums[j], index: j}
			h.push(r)
		}
		cur := h.data[0]
		for cur.index <= j-k {
			h.pop()
			cur = h.data[0]
		}
		result = append(result, cur.val)
	}
	return result
}

// @lc code=end

