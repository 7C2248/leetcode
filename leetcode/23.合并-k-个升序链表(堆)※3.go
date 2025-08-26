/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并 K 个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Definition for singly-linked list.
type MinHeap struct {
	data []*ListNode
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

func (h *MinHeap) Push(node *ListNode) {
	h.data = append(h.data, node)
	h.up(h.Len() - 1)
}

func (h *MinHeap) Pop() *ListNode {
	if h.Len() == 0 {
		return nil
	}
	root := h.data[0]
	last := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	if h.Len() > 0 {
		h.data[0] = last
		h.down(0)
	}
	return root
}

func (h *MinHeap) up(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if h.data[p].Val <= h.data[i].Val {
			break
		}
		h.data[p], h.data[i] = h.data[i], h.data[p]
		i = p
	}
}

func (h *MinHeap) down(i int) {
	n := h.Len()
	for {
		left := i*2 + 1
		right := i*2 + 2
		smallest := i

		if left < n && h.data[left].Val < h.data[smallest].Val {
			smallest = left
		}
		if right < n && h.data[right].Val < h.data[smallest].Val {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	heap := &MinHeap{}

	for _, l := range lists {
		if l != nil {
			heap.Push(l)
		}
	}

	dummy := &ListNode{}
	cur := dummy

	for heap.Len() > 0 {
		node := heap.Pop()
		cur.Next = node
		cur = cur.Next
		if node.Next != nil {
			heap.Push(node.Next)
		}
	}

	return dummy.Next
}

// @lc code=end

