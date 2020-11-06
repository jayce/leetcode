/*
 * Creation in 2020-11-06 by Chuan Liu (Jayce)
 */
package leetcode

func getDecimalValue(head *ListNode) int {
	n := 0
	for ; head != nil; head = head.Next { // time: O(N)
		n = n<<1 + head.Val
	}
	return n
}
