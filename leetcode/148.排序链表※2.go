/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*

//较慢,最坏情况下O(N^2)
func sortList(head *ListNode) *ListNode {
	mapping := make(map[int][]*ListNode)
	valArray := []int{}

	for i := head; i != nil; i = i.Next {
		if _, exists := mapping[i.Val]; !exists {
			valArray = append(valArray, i.Val)
		}
		mapping[i.Val] = append(mapping[i.Val], i)
	}

	n := len(valArray)
	sort.Ints(valArray)
	th := &ListNode{}
	now := th
	for i := 0; i < n; i++ {
		for j := range mapping[valArray[i]] {
			now.Next = mapping[valArray[i]][j]
			now = now.Next
		}
	}
	now.Next = nil

	return th.Next
}
*/

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 找到中点
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 分割链表
	mid := slow.Next
	slow.Next = nil

	// 递归排序
	left := sortList(head)
	right := sortList(mid)

	// 合并
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

// @lc code=end

