/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] O(1) 时间插入、删除和获取随机元素
 */

// @lc code=start
import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	len  int
	dict map[int]int
	list []int
	rng  *rand.Rand
}

func Constructor() RandomizedSet {
	r := RandomizedSet{
		len:  0,
		dict: make(map[int]int),
		list: make([]int, 0),
		rng:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	return r
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.dict[val]
	if ok {
		return false
	}
	this.list = append(this.list, val)
	this.dict[val] = this.len
	this.len += 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.dict[val]
	if ok {

		this.len -= 1
		this.list[index] = this.list[this.len]
		this.dict[this.list[index]] = index
		this.list = this.list[0:this.len]
		delete(this.dict, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	i := this.rng.Intn(this.len)
	return this.list[i]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

