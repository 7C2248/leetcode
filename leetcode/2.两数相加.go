/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	i, j := l1, l2
	var list3 ListNode
	k := &list3
	positive := 0
	for i != nil && j != nil {
		sum := i.Val + j.Val + positive
		positive = sum / 10
		var temp ListNode
		k.Next = &temp
		k = k.Next
		k.Val = sum % 10
		i = i.Next
		j = j.Next
	}

	if i == nil && j != nil {
		k.Next = j
		for j.Next != nil {
			j.Val, positive = (j.Val+positive)%10, (j.Val+positive)/10
			j = j.Next
		}
		j.Val, positive = (j.Val+positive)%10, (j.Val+positive)/10
		k = j
	}
	if j == nil && i != nil {
		k.Next = i
		for i.Next != nil {
			i.Val, positive = (i.Val+positive)%10, (i.Val+positive)/10
			i = i.Next
		}
		i.Val, positive = (i.Val+positive)%10, (i.Val+positive)/10
		k = i
	}
	if positive > 0 {
		var temp ListNode
		k.Next = &temp
		k = k.Next
		k.Val = positive
	}
	return list3.Next
}

// @lc code=end

