/*
 * Creation in 2020-11-05 by Chuan Liu (Jayce)
 */

package leetcode

// IsSubPath First idea
func IsSubPath(head *ListNode, root *TreeNode) bool { // time: O(N)
	if root == nil || head == nil {
		return head == nil
	}

	return dfs(head, root) ||
		IsSubPath(head, root.Left) ||
		IsSubPath(head, root.Right)
}

func dfs(head *ListNode, root *TreeNode) bool { // time: O(N)
	if root == nil || head == nil {
		return head == nil
	}

	return head.Val == root.Val &&
		(dfs(head.Next, root.Left) || dfs(head.Next, root.Right))
}

// Second idea
func isSubPath(head *ListNode, root *TreeNode) bool {
	aqueue := make([]*TreeNode, 0, 10) // space: O(W)
	bqueue := make([]*TreeNode, 0, 10) // space: O(W)
	aqueue = append(aqueue, root)
	for len(aqueue) > 0 {
		bqueue = bqueue[:0]
		for _, node := range aqueue { // time: O(N)
			if dfs(head, node) { // time: O(N)
				return true
			}
			if node.Left != nil {
				bqueue = append(bqueue, node.Left)
			}
			if node.Right != nil {
				bqueue = append(bqueue, node.Right)
			}
		}
		aqueue, bqueue = bqueue, aqueue
	}
	return false
}
