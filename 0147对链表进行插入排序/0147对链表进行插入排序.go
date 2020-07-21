// @Author:  limuye@sgepri.sgcc.com.cn
// @Desc:    排序 链表
// @Version: V1.0.0
// @Create:  2020/7/20 18:51

package LeetCode

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	res := &ListNode{0, nil}
	pre := res
	var tmpPre *ListNode

	for cur := head; cur != nil; cur = cur.Next {
		if pre.Next == nil {
			pre.Next = head
			pre = head
			continue
		}

		if cur.Val >= pre.Val {
			pre = cur
			continue
		}

		tmpPre = res
		for tmp := res.Next; tmp != pre.Next; tmp = tmp.Next {
			if cur.Val <= tmp.Val {
				pre.Next = cur.Next
				tmpPre.Next = cur
				cur.Next = tmp
				cur = pre
				break
			}
			tmpPre = tmp
		}
	}

	return res.Next
}
