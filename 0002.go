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

// Second idea
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersHelp(l1, l2, 0)
}

// space: O(1), time: O(min(N, M))
func addTwoNumbersHelp(l1, l2 *ListNode, remain int) *ListNode {
	if l1 == nil && l2 == nil && remain == 0 {
		return nil
	}

	v1, v2 := l1, l2
	if l1 != nil {
		l1 = l1.Next
	}
	if l2 != nil {
		l2 = l2.Next
	}

	if v2 == nil {
		v2 = &ListNode{}
	}
	if v1 == nil {
		v1 = v2
	}

	v1.Val = v1.Val + v2.Val + remain
	v1.Next = addTwoNumbersHelp(l1, l2, v1.Val/10)
	v1.Val = v1.Val % 10
	return v1
}

// space: O(max(N, M)), time: O(max(N, M))
func addTwoNumbersHelpII(l1, l2 *ListNode, remain int) *ListNode {
	if l1 == nil && l2 == nil && remain == 0 {
		return nil
	}

	v1, v2 := 0, 0
	if l1 != nil {
		v1, l1 = l1.Val, l1.Next
	}

	if l2 != nil {
		v2, l2 = l2.Val, l2.Next
	}

	sum := v1 + v2 + remain
	return &ListNode{
		Val:  sum % 10,
		Next: addTwoNumbersHelp(l1, l2, sum/10),
	}
}
