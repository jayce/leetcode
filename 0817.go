/*
 * Creation in 2020-11-01 by Chuan Liu (Jayce)
 */

package leetcode

func linkListComponent(G []int, head *ListNode) int {
	gm := make(map[int]bool, 0)
	for _, v := range G {
		gm[v] = true
	}

	component := 0
	for p := head; p != nil; p = p.Next {
		if gm[p.Val] && (p.Next == nil || !gm[p.Next.Val]) {
			component++
		}
	}
	return component
}
