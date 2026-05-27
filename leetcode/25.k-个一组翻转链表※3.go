/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := &ListNode{Next: head}
	temp := pre
	for pre != nil {
		var start *ListNode = pre.Next
		if start == nil {
			break
		}

		kth := pre
		for i := 0; i < k; i++ {
			kth = kth.Next
			if kth == nil {
				return temp.Next
			}
		}
		tail := kth.Next
		kth.Next = nil

		prev := tail
		cur := start
		for cur != nil {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}
		pre.Next = prev

		pre = start
	}
	return temp.Next
}

// @lc code=end

