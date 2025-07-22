/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	i, j := list1, list2
	var list3 ListNode
	k := &list3
	for i != nil && j != nil {
		if i.Val >= j.Val {
			k.Next = j
			k = k.Next
			j = j.Next
		} else {
			k.Next = i
			k = k.Next
			i = i.Next
		}
	}
	if i == nil && j != nil {
		for j.Next != nil {
			k.Next = j
			k = k.Next
			j = j.Next
		}
		k.Next = j
	}
	if j == nil && i != nil {
		for i.Next != nil {
			k.Next = i
			k = k.Next
			i = i.Next
		}
		k.Next = i
	}
	return list3.Next
}

// @lc code=end

