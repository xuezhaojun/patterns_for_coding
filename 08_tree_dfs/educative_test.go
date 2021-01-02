// 使用recursion进行遍历，空间复杂度最多为O(H),H是树的深度
package tree_dfs

import "math"

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
// 空间使用有待优化，毕竟每次传递[][]int的数据是不小的开销
// 使用同时传递 current path 和 all path的思路
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

// Sum of Path Numbers
// https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/?utm_source=LCUS&utm_medium=ip_redirect&utm_campaign=transfer2china
// 其实本题就是上面题的变式，首先还是需要求出所有的path，然后将所有的path变成数字加起来即可
// 本题学习的是dfs从上到下的思路，即可以从低而上，也可以从顶而下【重要】
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, preSum int) int {
	if root == nil {
		return 0
	}

	sum := preSum*10 + root.Val

	// if node is leaf
	if root.Left == nil && root.Right == nil {
		return sum
	}

	return dfs(root.Left, sum) + dfs(root.Right, sum)
}

// Path With Given Sequence
func pathWithSequence(root *TreeNode, seq []int) bool {
	if root == nil {
		return false
	}

	if len(seq) < 1 {
		return false
	}

	if root.Val != seq[0] {
		return false
	}

	if len(seq) == 1 && root.Left == nil && root.Right == nil && root.Val == seq[0] {
		return true
	}

	return pathWithSequence(root.Left, seq[1:]) || pathWithSequence(root.Right, seq[1:])
}

// Count Paths for a Sum
// https://leetcode-cn.com/problems/path-sum-iii/
// 本题的巧妙在于，每次统计当前path下满足的数量时候，都是必须加上当前节点，确保了不重复序列的【不重复】
func pathSum3(root *TreeNode, sum int) int {
	return dfs_3(root, []int{}, sum)
}

func dfs_3(root *TreeNode, path []int, sum int) int {
	if root == nil {
		return 0
	}

	same := 0
	path = append(path, root.Val)
	temp := 0
	for i := len(path) - 1; i >= 0; i-- {
		temp += path[i]
		if temp == sum {
			same += 1
		}
	}

	return same + dfs_3(root.Left, path, sum) + dfs_3(root.Right, path, sum)
}

// Tree Diameter
// https://leetcode.com/problems/diameter-of-binary-tree/
func diameterOfBinaryTree(root *TreeNode) int {
	var ans = 1       // 本题中，定义了ans，一个全局变量
	depth(root, &ans) // 此处传递的了ans的地址，而不是ans本身
	return ans - 1
}

func depth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	left := depth(root.Left, ans)
	right := depth(root.Right, ans)

	// depth中需要更新两个概念：
	// length是我们最终需要返回的答案，是两个leaf之间的距离
	// 而depth则是当前node的深度
	// 通过这两个概念可以简化左右节点同时记录的麻烦

	newLength := left + right + 1
	if *ans < newLength {
		*ans = newLength
	}

	var newDepth int
	if left > right {
		newDepth = left
	} else {
		newDepth = right
	}

	return newDepth + 1
}

// Path with Maximum Sum
// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
// 本题为上一个题的延续，但是同时存在难点，比如：
// 路径的定义改变，可以从任意一个节点出发到任意一个节点（并不一定是leaf开始和结束）原来是统计depth
// 原来是计算
func maxPathSum(root *TreeNode) int {
	// 本题中的一个有趣的点，如何使用匿名函数进行递归
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 只有子节点贡献值大于0的时候，才会选择子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		priceNewPath := node.Val + leftGain + rightGain

		maxSum = max(maxSum, priceNewPath) // update maxSum

		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
