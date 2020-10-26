/*
 * Creation in 2020-10-15 by Chuan Liu (Jayce)
 */

package leetcode

// MergeKLists First idea
func MergeKLists(lists []*ListNode) *ListNode {
	d := &ListNode{}
	p := d
	lists = removeEmptyList(lists) // time: O(N)
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n > 1 {
		insertSortList(lists) // time: O(N*N) or O(N)
	}

	for {
		if n <= 1 {
			p.Next = lists[0]
			break
		}
		nextv := lists[1].Val
		for lists[0] != nil && lists[0].Val <= nextv { // time: O(k*N)
			p.Next = lists[0]
			p, lists[0] = p.Next, lists[0].Next
		}
		if lists[0] == nil { // remove to tail
			n--
			for i := 0; i < n-1; i++ { // time: O(N)
				lists[i], lists[i+1] = lists[i+1], lists[i]
			}
		} else {
			for i := 0; i < n-1; i++ { // sort time: O(N)
				if lists[i].Val <= lists[i+1].Val {
					break
				}
				lists[i], lists[i+1] = lists[i+1], lists[i]
			}
		}
	}

	return d.Next
}

func removeEmptyList(lists []*ListNode) []*ListNode {
	data := make([]*ListNode, 0)
	for i, v := range lists {
		if v != nil {
			data = append(data, v)
		}
	}
	return data
}

func insertSortList(lists []*ListNode) {
	for i := 1; i < len(lists); i++ {
		for j := i; j > 0; j-- {
			if lists[j].Val < lists[j-1].Val {
				lists[j], lists[j-1] = lists[j-1], lists[j]
			}
		}
	}
}

// MergeKListsII inspire by meregeTwoLists
func MergeKListsII(lists []*ListNode) *ListNode { // time: O(n+m) * O(nlogn)
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	if n == 2 {
		return mergeTwoLists(lists[0], lists[1]) // time: O(n+m)
	}
	m := n / 2
	return mergeTwoLists(MergeKListsII(lists[0:m]), MergeKListsII(lists[m:]))
}

// MergeKListsIII non-recursive
func MergeKListsIII(lists []*ListNode) *ListNode { // time: O(k*N)
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	dummy := &ListNode{}
	for i := 0; i < len(lists); i++ { // time: O(k)
		dummy.Next = mergeTwoLists(dummy.Next, lists[i]) // time: O(N)
	}
	return dummy.Next
}

func mergeTwoLists(A, B *ListNode) *ListNode {
	dummy := &ListNode{}
	mergeTwoListsII(dummy, A, B)
	return dummy.Next
}

func mergeTwoListsII(dummy, A, B *ListNode) *ListNode {
	if A == nil {
		dummy.Next = B
		return A
	}
	if B == nil {
		dummy.Next = A
		return B
	}
	if A.Val > B.Val {
		dummy.Next = B
		return mergeTwoListsII(dummy.Next, A, B.Next)
	}
	dummy.Next = A
	return mergeTwoListsII(dummy.Next, A.Next, B)
}
