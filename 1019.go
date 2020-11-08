/*
 * Creation in 2020-11-08 by Chuan Liu (Jayce)
 */

package leetcode

// nextLargerNodes First idea
func nextLargerNodes(head *ListNode) []int {
	m := make([]int, 0, 100)

	for p := head; p != nil; p = p.Next { // time: O(N*N)
		n := p.Next
		for ; n != nil; n = n.Next {
			if n.Val > p.Val {
				m = append(m, n.Val)
				break
			}
		}

		if n == nil {
			m = append(m, 0)
		}
	}
	return m
}

// [1, 7, 5, 1, 9, 2, 5, 1]
//  0  1                     => stack[0]       1 -> [0] = [1]
//     1  2  3               => stack[1, 2, 3] 4 -> [1, 2, 3] = [4]
//              4  5         => stack[4, 5]    6 -> [5] = [6]
//                    6  7   => stack[4, 6, 7]   -> [4, 6, 7] = 0
func nextLargerNodes(head *ListNode) []int {
	nodes := make([]*ListNode, 0)         // space: O(N)
	for p := head; p != nil; p = p.Next { // time: O(N)
		nodes = append(nodes, p)
	}

	results := make([]int, len(nodes))
	stack := make([]int, 0)           // space: O(M)
	for i := 0; i < len(nodes); i++ { // time: O(N) + O(M)
		for j := len(stack); j > 0 && nodes[stack[j-1]].Val < nodes[i].Val; j-- {
			results[stack[j-1]] = nodes[i].Val
			stack = stack[:j-1]
		}
		stack = append(stack, i)
	}
	return results
}
