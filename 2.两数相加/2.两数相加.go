// @Author:  limuye@sgepri.sgcc.com.cn
// @Desc:    
// @Version: V1.0.0
// @Create:  2020/7/5 17:21

package LeetCode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	p := l3
	carry := 0
	var l1Val, l2Val int

	for l1 != nil || l2 != nil || carry != 0 {
		l1Val = 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}

		l2Val = 0
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}

		p.Val = l1Val + l2Val + carry
		carry = p.Val / 10
		p.Val = p.Val % 10
		if l1 != nil || l2 != nil || carry != 0 {
			p.Next = &ListNode{}
			p = p.Next
		} else {
			p.Next = nil
			break
		}
	}

	return l3
}
