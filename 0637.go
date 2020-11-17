/*
 * Creation in 2020-11-17 by Chuan Liu (Jayce)
 */

package leetcode

// averageOfLevels First idea
func averageOfLevels(root *TreeNode) []float64 { // time: O(N)
	if root == nil {
		return nil
	}
	averages := []float64{}
	aqueue := []*TreeNode{root}
	bqueue := []*TreeNode{}
	for len(aqueue) > 0 {
		sum := 0
		bqueue = bqueue[:0]
		for _, node := range aqueue {
			sum += node.Val
			if node.Left != nil {
				bqueue = append(bqueue, node.Left)
			}
			if node.Right != nil {
				bqueue = append(bqueue, node.Right)
			}
		}
		averages = append(averages, float64(sum)/float64(len(aqueue)))
		aqueue, bqueue = bqueue, aqueue
	}
	return averages
}
