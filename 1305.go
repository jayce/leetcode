/*
 * Creation in 2020-12-14 by Chuan Liu (Jayce)
 */

package leetcode

// getAllElements First idea
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	return merge(dfs(root1), dfs(root2)) // time: O(3N)
}

func dfs(root *TreeNode) []int { // time: O(N)
	if root == nil {
		return []int{}
	}
	// malloc: O(N*N)
	return append(dfs(root.Left), append([]int{root.Val}, dfs(root.Right)...)...)
}

func merge(a, b []int) []int { // time: O(N)
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	if a[0] < b[0] {
		return append([]int{a[0]}, merge(a[1:], b)...)
	}
	return append([]int{b[0]}, merge(a, b[1:])...)
}

// Second idea for optimize for malloc, time: O(N) space: O(N)
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	data := make([]int, 0, 4) // space: O(N)
	var dfs func(root *TreeNode) []int
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		data = append(data, root.Val)
		dfs(root.Right)
	}

	dfs(root1) // time: O(N)
	an := len(data)
	dfs(root2) // time: O(N)

	c := make([]int, len(data)) // space: O(N)
	merge(data[0:an], data[an:], c)
	return c
}

func merge(a, b, c []int) { // time: O(N)
	if len(a) == 0 {
		copy(c, b)
		return
	}
	if len(b) == 0 {
		copy(c, a)
		return
	}
	if a[0] > b[0] {
		c[0] = b[0]
		b = b[1:]
	} else {
		c[0] = a[0]
		a = a[1:]
	}
	merge(a, b, c[1:])
}

// Third idea traversal them at the same time
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	data := make([]int, 0, 4)
	anext, bnext := inorderFunc(root1), inorderFunc(root2)
	aroot, broot := anext(), bnext()   // space: O(H)
	croot := (*TreeNode)(nil)          // just init
	for aroot != nil || broot != nil { // time: O(N)
		switch {
		case broot == nil || (aroot != nil && aroot.Val <= broot.Val):
			croot = aroot
			aroot = adfs()
		case aroot == nil || aroot.Val > broot.Val:
			croot = broot
			broot = bnext()
		}

		data = append(data, croot.Val)
	}
	return data
}

func inorderFunc(root *TreeNode) func() *TreeNode {
	stack := make([]*TreeNode, 0, 4)
	return func() *TreeNode {
		for len(stack) > 0 || root != nil {
			for root != nil {
				stack = append(stack, root)
				root = root.Left
			}

			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			value := root
			root = root.Right
			return value
		}
		return nil
	}
}
