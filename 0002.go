/*
 * Creation in 2020-11-01 by Chuan Liu (Jayce)
 */

package leetcode

// AddTwoNumbers First idea
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	l3 := dummy
	t, v1, v2, v3 := 0, 0, 0, 0
	for l1 != nil || l2 != nil || t > 0 { // time: O(N+M)
		l1, v1 = popValueAndNext(l1)
		l2, v2 = popValueAndNext(l2)

		v3 = v1 + v2 + t
		t = v3 / 10

		l3.Next = &ListNode{Val: v3 % 10}
		l3 = l3.Next
	}

	return dummy.Next
}

func popValueAndNext(root *ListNode) (*ListNode, int) {
	if root == nil {
		return nil, 0
	}

	return root.Next, root.Val
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersHelp(l1, l2, 0)
}

func addTwoNumbersHelp(l1, l2 *ListNode, remain int) *ListNode { // time: O(min(N,M))
	if l1 == nil && l2 == nil {
		if remain > 0 {
			return &ListNode{Val: remain}
		}
		return nil
	}

	if l1 == nil {
		if remain == 0 {
			return l2
		}
		sum := l2.Val + remain
		l2.Val = sum % 10
		l2.Next = addTwoNumbersHelp(l1, l2.Next, sum/10)
		return l2
	}

	if l2 == nil {
		if remain == 0 {
			return l1
		}
		sum := l1.Val + remain
		l1.Val = sum % 10
		l1.Next = addTwoNumbersHelp(l1.Next, l2, sum/10)
		return l1
	}

	sum := l1.Val + l2.Val + remain
	l1.Val = sum % 10
	l1.Next = addTwoNumbersHelp(l1.Next, l2.Next, sum/10)
	return l1
}
