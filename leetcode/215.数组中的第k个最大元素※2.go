/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
type heap struct {
	val []int
}

func build(nums []int) *heap {
	h := heap{}
	for i := range nums {
		h.push(nums[i])
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
		if r < n && this.val[r] > this.val[l] {
			largest = r
		}
		if this.val[largest] > this.val[index] {
			this.val[index], this.val[largest] = this.val[largest], this.val[index]
			index = largest
		} else {
			break
		}
	}
}

func (this *heap) down(index int) {

	for index > 0 {

		if this.val[(index-1)>>1] >= this.val[index] {
			break
		}
		this.val[(index-1)>>1], this.val[index] = this.val[index], this.val[(index-1)>>1]
		index = (index - 1) >> 1
	}
}

func (this *heap) pop() int {
	result := this.val[0]

	n := len(this.val) - 1
	this.val[0] = this.val[n]
	this.val = this.val[:n]
	if len(this.val) > 0 {
		this.up(0)
	}
	return result
}

func (this *heap) push(v int) {

	this.val = append(this.val, v)
	index := len(this.val) - 1
	this.down(index)
}

func findKthLargest(nums []int, k int) int {
	h := build(nums)
	result := 0
	for i := 0; i < k; i++ {
		result = h.pop()
	}
	return result
}

/*
//期望O(n)
func findKthLargest(nums []int, k int) int {
    n := len(nums)
    return quickselect(nums, 0, n - 1, n - k)
}

func quickselect(nums []int, l, r, k int) int{
    if (l == r){
        return nums[k]
    }
    partition := nums[l]
    i := l - 1
    j := r + 1
    for (i < j) {
        for i++;nums[i]<partition;i++{}
        for j--;nums[j]>partition;j--{}
        if (i < j) {
            nums[i],nums[j]=nums[j],nums[i]
        }
    }
    if (k <= j){
        return quickselect(nums, l, j, k)
    }else{
        return quickselect(nums, j + 1, r, k)
    }
}
*/

// @lc code=end

