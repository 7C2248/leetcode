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

/*
// 自底向上递归，空间O(1)
// cut 从 head 开始切出前 n 个节点（不足则切到链表尾）
// 返回剩余链表头 next
func cut(head *ListNode, n int) (next *ListNode) {
    if head == nil {
        return nil
    }
    cur := head
    for i := 1; i < n && cur.Next != nil; i++ {
        cur = cur.Next
    }

    next = cur.Next
    cur.Next = nil // 切断
    return
}

// merge 合并两个已排序链表，返回新头和新尾
func merge(l, r *ListNode) (head, tail *ListNode) {
    var dummy ListNode          // 栈上分配，无堆内存
    cur := &dummy

    for l != nil && r != nil {
        if l.Val <= r.Val {
            cur.Next = l
            l = l.Next
        } else {
            cur.Next = r
            r = r.Next
        }
        cur = cur.Next
    }
    // 连接剩余部分
    if l != nil {
        cur.Next = l
    } else {
        cur.Next = r
    }
    // 找尾
    for cur.Next != nil {
        cur = cur.Next
    }
    return dummy.Next, cur
}

func sortList(head *ListNode) *ListNode {
    // 1. 计算长度
    length := 0
    for n := head; n != nil; n = n.Next {
        length++
    }

    start := &ListNode{Next: head}

    // 2. 自底向上归并
    for step := 1; step < length; step <<= 1 {
        prev := start       // 已合并部分的前驱
        cur := start.Next   // 当前待处理段头

        for cur != nil {
            // 切第一段
            left := cur
            mid := cut(left, step)

            if mid == nil {
                // 只剩一段，无需合并，直接连回
                prev.Next = left
                break
            }

            // 切第二段
            right := mid
            next := cut(right, step)

            // 合并两段并接回链表
            mergedHead, mergedTail := merge(left, right)
            prev.Next = mergedHead
            mergedTail.Next = next

            // 移动指针
            prev = mergedTail
            cur = next
        }
    }

    return start.Next
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

