// 使用recursion进行遍历，空间复杂度最多为O(H),H是树的深度
package tree_dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

// Binary Tree Path Sum
// https://leetcode-cn.com/problems/path-sum/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	// 如何判断一个节点是否是叶子节点？
	// 如果一个节点的left和right都是null，则说明这个阶段是叶子节点
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

// All Paths for a Sum
// https://leetcode-cn.com/problems/path-sum-ii/
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	if root.Left == nil && root.Right == nil && root.Val == sum {
		return [][]int{
			{root.Val},
		}
	}

	var result [][]int
	leftResult := pathSum(root.Left, sum-root.Val)
	rightResult := pathSum(root.Right, sum-root.Val)

	if len(leftResult) > 0 {
		for _, r := range leftResult {
			result = append(result, append([]int{root.Val}, r...))
		}
	}

	if len(rightResult) > 0 {
		for _, r := range rightResult {
			result = append(result, append([]int{root.Val}, r...))
		}
	}

	return result
}
