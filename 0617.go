/*
 * Creation in 2020-11-15 by Chuan Liu (Jayce)
 */

package leetcode

// mergeTrees First idea
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode { // time: O(min(N,M))
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

// Second idea time: O(max(N,M))
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		t1 = &TreeNode{}
	}
	if t2 == nil {
		t2 = &TreeNode{}
	}
	return &TreeNode{
		Val:   t1.Val + t2.Val,
		Left:  mergeTrees(t1.Left, t2.Left),
		Right: mergeTrees(t1.Right, t2.Right),
	}
}

// Third idea
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	t3 := &TreeNode{}
	root := t3
	stack := []*TreeNode{}                         // space: O(3*max(t1,t2))
	for len(stack) > 0 || t1 != nil || t2 != nil { // time: O(max(t1,t2))
		for t1 != nil || t2 != nil {
			if t1 == nil {
				t1 = &TreeNode{}
			}
			if t2 == nil {
				t2 = &TreeNode{}
			}

			stack = append(stack, t1, t2, t3)
			t1, t2 = t1.Left, t2.Left

			if t1 != nil || t2 != nil {
				t3.Left = &TreeNode{}
				t3 = t3.Left
			}
		}

		t1 = stack[len(stack)-3]
		t2 = stack[len(stack)-2]
		t3 = stack[len(stack)-1]
		t3.Val = t1.Val + t2.Val
		stack = stack[:len(stack)-3]

		t1, t2 = t1.Right, t2.Right
		if t1 != nil || t2 != nil {
			t3.Right = &TreeNode{}
			t3 = t3.Right
		}
	}

	return root
}
