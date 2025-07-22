/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
type LRUCache struct {
	capacity, used int
	hNode, tNode   *linklist
	hmap           map[int]*linklist
}

type linklist struct {
	key, value   int
	fNode, sNode *linklist
}

func Constructor(capacity int) LRUCache {
	var r, rr linklist

	r = linklist{
		key:   -1,
		value: -1,
		fNode: nil,
		sNode: &rr,
	}

	rr = linklist{
		key:   -1,
		value: -1,
		fNode: &r,
		sNode: nil,
	}

	return LRUCache{
		capacity: capacity,
		used:     0,
		hNode:    &r,
		tNode:    &rr,
		hmap:     make(map[int]*linklist),
	}
}

func (this *LRUCache) Get(key int) int {

	hNode, ok := this.hmap[key]
	if !ok {
		return -1
	}

	hNode.fNode.sNode = hNode.sNode
	hNode.sNode.fNode = hNode.fNode
	this.hNode.sNode.fNode = hNode
	hNode.sNode = this.hNode.sNode
	this.hNode.sNode, hNode.fNode = hNode, this.hNode

	return hNode.value
}

func (this *LRUCache) Put(key int, value int) {

	hNode, ok := this.hmap[key]

	if ok {
		hNode.value = value

		hNode.fNode.sNode = hNode.sNode

		hNode.sNode.fNode = hNode.fNode

		this.hNode.sNode.fNode = hNode

		hNode.sNode = this.hNode.sNode
		this.hNode.sNode, hNode.fNode = hNode, this.hNode
	} else {
		if this.used == this.capacity {
			delete(this.hmap, this.tNode.fNode.key)
			this.tNode.fNode, this.tNode.fNode.fNode.sNode = this.tNode.fNode.fNode, this.tNode
		} else {
			this.used++
		}

		r := linklist{
			key:   key,
			value: value,
			fNode: this.hNode,
			sNode: this.hNode.sNode,
		}
		this.hmap[key] = &r
		this.hNode.sNode.fNode = &r
		this.hNode.sNode = &r
	}

	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

