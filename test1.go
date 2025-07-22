package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*test := []int{2, 1, 3, 2, 1, 3, 3, 6, 9, 8, 5, 4, 1, 0, 2}
	sort(test, 0, len(test))
	fmt.Println(test)

	r := Constructor()
	r.Insert(0)
	r.Insert(1)
	r.Remove(0)
	r.Insert(2)
	r.Remove(1)
	fmt.Println(r.GetRandom())
	*/
	s := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(s))

}

func threeSum(nums []int) [][]int {

	result := make([][]int, 0)
	n := len(nums)
	sort(nums, 0, n)
	for i := 0; i <= n-3; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}

		for j, k := i+1, n-1; j < k; j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for nums[i]+nums[j]+nums[k] > 0 && k > j {
				k--
			}

			if nums[i]+nums[j]+nums[k] == 0 {
				tt := make([]int, 0)
				tt = append(tt, nums[i], nums[j], nums[k])
				result = append(result, tt)
			}
		}
	}
	return result
}

func sort(aa []int, left int, right int) {

	if right-left <= 1 {
		return
	}

	i, j := left, right-1
	for i < j {
		for aa[j] > aa[left] && i < j {
			j--
		}
		for aa[i] <= aa[left] && i < j {
			i++
		}

		aa[i], aa[j] = aa[j], aa[i]
	}
	aa[left], aa[i] = aa[i], aa[left]
	sort(aa, left, i)
	sort(aa, i+1, right)
	return
}

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
