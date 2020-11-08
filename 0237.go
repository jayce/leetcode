/*
 * Creation in 2020-11-08 by Chuan Liu (Jayce)
 */

package leetcode

func deleteNode(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
}
