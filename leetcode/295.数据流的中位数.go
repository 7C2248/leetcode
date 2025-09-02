/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
type minheap struct {
	data []int
}

func buildmin(data []int) *minheap {
	h := minheap{}
	for i := range data {
		h.push(data[i])
	}
	return &h
}
func (this *minheap) Len() int {
	return len(this.data)
}
func (this *minheap) down(index int) {
	n := this.Len()
	for {
		l := (index << 1) + 1
		r := (index << 1) + 2
		if l >= n {
			break
		}
		smallest := l
		if r < n && this.data[r] < this.data[l] {
			smallest = r
		}
		if this.data[smallest] < this.data[index] {
			this.data[index], this.data[smallest] = this.data[smallest], this.data[index]
			index = smallest
		} else {
			break
		}
	}
}
func (this *minheap) up(index int) {
	for index > 0 {
		if this.data[(index-1)>>1] <= this.data[index] {
			break
		}
		this.data[(index-1)>>1], this.data[index] = this.data[index], this.data[(index-1)>>1]
		index = (index - 1) >> 1
	}
}
func (this *minheap) pop() int {
	result, n := this.data[0], this.Len()-1
	this.data[0] = this.data[n]
	this.data = this.data[0:n]
	if this.Len() > 0 {
		this.down(0)
	}
	return result
}
func (this *minheap) push(v int) {
	this.data = append(this.data, v)
	this.up(this.Len() - 1)
}

type maxheap struct {
	data []int
}

func buildmax(data []int) *maxheap {
	h := maxheap{}
	for i := range data {
		h.push(data[i])
	}
	return &h
}
func (this *maxheap) Len() int {
	return len(this.data)
}
func (this *maxheap) down(index int) {
	n := this.Len()
	for {
		l := (index << 1) + 1
		r := (index << 1) + 2
		if l >= n {
			break
		}
		largest := l
		if r < n && this.data[r] > this.data[l] {
			largest = r
		}
		if this.data[largest] > this.data[index] {
			this.data[index], this.data[largest] = this.data[largest], this.data[index]
			index = largest
		} else {
			break
		}
	}
}
func (this *maxheap) up(index int) {
	for index > 0 {
		if this.data[(index-1)>>1] >= this.data[index] {
			break
		}
		this.data[(index-1)>>1], this.data[index] = this.data[index], this.data[(index-1)>>1]
		index = (index - 1) >> 1
	}
}
func (this *maxheap) pop() int {
	result, n := this.data[0], this.Len()-1
	this.data[0] = this.data[n]
	this.data = this.data[0:n]
	if this.Len() > 0 {
		this.down(0)
	}
	return result
}
func (this *maxheap) push(v int) {
	this.data = append(this.data, v)
	this.up(this.Len() - 1)
}

type MedianFinder struct {
	minh *maxheap
	maxh *minheap
}

func Constructor() MedianFinder {
	return MedianFinder{buildmax([]int{}), buildmin([]int{})}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minh.Len() < 1 {
		this.minh.push(num)
		return
	}

	if this.minh.Len() > this.maxh.Len() {
		if num < this.minh.data[0] {
			rr := this.minh.pop()
			this.minh.push(num)
			this.maxh.push(rr)
		} else {
			this.maxh.push(num)
		}
	} else {
		if num < this.maxh.data[0] {
			this.minh.push(num)
		} else {
			rr := this.maxh.pop()
			this.minh.push(rr)
			this.maxh.push(num)
		}
	}

	return
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.minh.Len()+this.maxh.Len())%2 == 0 {
		return float64((this.minh.data[0] + this.maxh.data[0])) / 2.0
	}
	return float64(this.minh.data[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

