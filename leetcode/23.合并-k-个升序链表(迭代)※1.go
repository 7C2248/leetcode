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
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	interval := 1
	for interval < n {
		for i := 0; i+interval < n; i += interval * 2 {
			lists[i] = merge(lists[i], lists[i+interval])
		}
		interval *= 2
	}
	return lists[0]
}

func merge(a, b *ListNode) *ListNode {
	head := &ListNode{}
	now := head
	for a != nil && b != nil {
		if a.Val > b.Val {
			now.Next = b
			b = b.Next
		} else {
			now.Next = a
			a = a.Next
		}
		now = now.Next
	}
	if a != nil {
		now.Next = a
	}
	if b != nil {
		now.Next = b
	}
	return head.Next
}

// @lc code=end

