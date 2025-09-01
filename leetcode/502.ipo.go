/*
 * @lc app=leetcode.cn id=502 lang=golang
 *
 * [502] IPO
 */

// @lc code=start
import (
	"sort"
)

type pc struct {
	profit  int
	capital int
}
type heap struct {
	val []*pc
}

func build(nums []pc) *heap {
	h := heap{}
	for i := range nums {
		h.push(&nums[i])
	}
	return &h
}

func (this *heap) up(index int) {

	n := len(this.val)
	for {
		l, r := index*2+1, index*2+2
		if l >= n {
			break
		}
		largest := l
		if r < n && this.val[r].profit > this.val[l].profit {
			largest = r
		}
		if this.val[largest].profit > this.val[index].profit {
			this.val[index], this.val[largest] = this.val[largest], this.val[index]
			index = largest
		} else {
			break
		}
	}
}

func (this *heap) down(index int) {
	for index > 0 {
		if this.val[(index-1)>>1].profit >= this.val[index].profit {
			break
		}
		this.val[(index-1)>>1], this.val[index] = this.val[index], this.val[(index-1)>>1]
		index = (index - 1) >> 1
	}
}

func (this *heap) pop() *pc {
	n := len(this.val) - 1
	if n < 0 {
		return nil
	}
	result := this.val[0]
	this.val[0] = this.val[n]
	this.val = this.val[:n]
	if len(this.val) > 0 {
		this.up(0)
	}
	return result
}

func (this *heap) push(v *pc) {

	this.val = append(this.val, v)
	index := len(this.val) - 1
	this.down(index)
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {

	l := []pc{}
	n := len(capital)
	for i := 0; i < n; i++ {
		rr := pc{profits[i], capital[i]}
		l = append(l, rr)
	}
	sort.Slice(l, func(i, j int) bool { return l[i].capital < l[j].capital })
	//sort(l, 0, n-1)

	i := 0
	for ; i < n && l[i].capital <= w; i++ {
	}

	h := build(l[:i])

	for j := 0; j < k; j++ {
		rr := h.pop()
		if rr == nil {
			break
		}
		w += rr.profit
		for ; i < n && l[i].capital <= w; i++ {
			h.push(&l[i])
		}
	}
	return w
}

/*

//性能不够
func sort(input []pc, l, r int) {
	if l >= r {
		return
	}

	pivot := input[l].capital
	i, j := l-1, r+1
	for {
		for {
			i++
			if input[i].capital >= pivot {
				break
			}
		}
		for {
			j--
			if input[j].capital <= pivot {
				break
			}
		}
		if i >= j {
			break
		}
		input[i], input[j] = input[j], input[i]
	}
	sort(input, l, j)
	sort(input, j+1, r)
}
*/
// @lc code=end

