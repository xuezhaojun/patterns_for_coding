package tree_bfs

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Binary Tree Level Order Traversal
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
// 层序判断的话，除了要记录节点，还要记录节点所在的层数
func levelOrder(root *TreeNode) [][]int {
	type Node struct {
		*TreeNode
		level int
	}

	levels := make(map[int][]int)
	l := list.New()
	l.PushBack(Node{
		TreeNode: root,
		level:    0,
	})

	for l.Len() != 0 {
		front := l.Front().Value.(Node)
		if front.TreeNode != nil {
			levels[front.level] = append(levels[front.level], front.Val)
			l.PushBack(Node{
				TreeNode: front.Left,
				level:    front.level + 1,
			})
			l.PushBack(Node{
				TreeNode: front.Right,
				level:    front.level + 1,
			})
		}
		l.Remove(l.Front())
	}

	i := 0
	result := [][]int{}
	for {
		level, ok := levels[i]
		if !ok {
			break
		}
		result = append(result, level)
		i++
	}

	return result
}

// educative上的给的写法，对于统计层数部分code有所简化； 是真的直观来看一层一层统计的
func levelOrder2(root *TreeNode) [][]int {
	l := list.New()
	if root == nil {
		return [][]int{}
	}
	l.PushBack(root)

	result := [][]int{}
	for l.Len() != 0 {
		levelSize := l.Len()
		curLevel := []int{}
		for levelSize > 0 {
			front := l.Front().Value.(*TreeNode)
			curLevel = append(curLevel, front.Val)
			if front.Left != nil {
				l.PushBack(front.Left)
			}
			if front.Right != nil {
				l.PushBack(front.Right)
			}
			l.Remove(l.Front())
			levelSize--
		}
		result = append(result, curLevel)
	}

	return result
}

// Reverse Level Order Traversal
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
// 这种反向可能一开始想到的就是栈，通过栈来做逆序
func levelOrderBottom(root *TreeNode) [][]int {
	l := list.New()
	if root == nil {
		return [][]int{}
	}
	l.PushBack(root)

	result := [][]int{}
	for l.Len() != 0 {
		levelSize := l.Len()
		curLevel := []int{}
		for levelSize > 0 {
			front := l.Front().Value.(*TreeNode)
			curLevel = append(curLevel, front.Val)
			if front.Left != nil {
				l.PushBack(front.Left)
			}
			if front.Right != nil {
				l.PushBack(front.Right)
			}
			l.Remove(l.Front())
			levelSize--
		}
		result = append([][]int{curLevel}, result...) // 仅此处修改
	}

	return result
}

// Zigzag Traversal
// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
// 从zigzag这个题可以看出，curlevel这个概念确实简化了逻辑
func zigzagLevelOrder(root *TreeNode) [][]int {
	l := list.New()
	if root == nil {
		return [][]int{}
	}
	l.PushBack(root)

	result := [][]int{}
	zigzag := true
	for l.Len() != 0 {
		levelSize := l.Len()
		curLevel := []int{}
		for levelSize > 0 {
			front := l.Front().Value.(*TreeNode)
			if zigzag {
				curLevel = append(curLevel, front.Val)
			} else {
				curLevel = append([]int{front.Val}, curLevel...) // 如果逆序，则直接逆序添加curLevel即可
			}
			if front.Left != nil {
				l.PushBack(front.Left)
			}
			if front.Right != nil {
				l.PushBack(front.Right)
			}
			l.Remove(l.Front())
			levelSize--
		}
		zigzag = !zigzag
		result = append(result, curLevel)
	}

	return result
}
