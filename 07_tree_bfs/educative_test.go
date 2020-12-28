package tree_bfs

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
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

// Level Averages in a Binary Tree
// https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
func averageOfLevels(root *TreeNode) []float64 {
	result := []float64{}

	l := list.New()
	if root == nil {
		return result
	}
	l.PushBack(root)

	for l.Len() != 0 {
		len := l.Len()
		count := len
		sum := 0
		for count > 0 {
			count--
			front := l.Front()
			l.Remove(front)
			treeNode := front.Value.(*TreeNode)
			sum += treeNode.Val
			if treeNode.Left != nil {
				l.PushBack(treeNode.Left)
			}
			if treeNode.Right != nil {
				l.PushBack(treeNode.Right)
			}
		}
		result = append(result, float64(sum)/float64(len))
	}

	return result
}

// Minimum Depth of a Binary Tree
// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
func minDepth(root *TreeNode) int {
	level := 0

	l := list.New()
	if root == nil {
		return 0
	}
	l.PushBack(root)

	for l.Len() != 0 {
		level++
		len := l.Len()
		for len > 0 {
			len--
			node := l.Front()
			l.Remove(node)
			treeNode := node.Value.(*TreeNode)
			if treeNode.Left == nil && treeNode.Right == nil {
				return level
			}
			if treeNode.Left != nil {
				l.PushBack(treeNode.Left)
			}
			if treeNode.Right != nil {
				l.PushBack(treeNode.Right)
			}
		}
	}

	return level
}

// Level Order Successor
// 给定一个二分树和一个node，找到这个点在这个树上level order successor
// 感觉上和zigzag的问题相似
// 官方给的解决方案中，直接在key相等的时候，break掉了，然后返回queue中的0位置的entry大小
func levelOrderSuccessor(root *TreeNode, key int) (los int) {
	// recall之前的知识： 使用quene进行广度遍历
	level := 0

	l := list.New()
	if root == nil {
		return 0
	}
	l.PushBack(root)

	next := false

	for l.Len() != 0 {
		level++
		len := l.Len()
		for len > 0 {
			len--
			node := l.Front()
			l.Remove(node)
			treeNode := node.Value.(*TreeNode)
			treeValue := treeNode.Val
			if next {
				return treeValue
			}
			if treeValue == key {
				next = true
			}
			if treeNode.Left != nil {
				l.PushBack(treeNode.Left)
			}
			if treeNode.Right != nil {
				l.PushBack(treeNode.Right)

			}
		}
	}

	return
}

// Connect Level Order Siblings
// leetcode https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/?utm_source=LCUS&utm_medium=ip_redirect&utm_campaign=transfer2china
// 优化方案（目前未实现）：
// 第N+1层的节点，其实可以通过N层节点来协助建立
// 本题还存在变式：即并不是连接右边，而是connect to next node
func connect(root *Node) *Node {
	// recall之前的知识： 使用quene进行广度遍历
	level := 0
	l := list.New()
	if root == nil {
		return root
	}
	l.PushBack(root)

	for l.Len() != 0 {
		level++
		len := l.Len()
		for len > 0 {
			len--
			node := l.Front()
			l.Remove(node)
			treeNode := node.Value.(*Node)

			if l.Len() != 0 && len != 0 {
				nextNode := l.Front()
				nextTreeNode := nextNode.Value.(*Node)
				treeNode.Next = nextTreeNode
			}

			if treeNode.Left != nil {
				l.PushBack(treeNode.Left)
			}
			if treeNode.Right != nil {
				l.PushBack(treeNode.Right)
			}
		}
	}
	return root
}

// Right View of a Binary Tree
// https://leetcode-cn.com/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) (results []int) {
	level := 0

	l := list.New()
	if root == nil {
		return
	}
	l.PushBack(root)

	for l.Len() != 0 {
		level++
		len := l.Len()
		for len > 0 {
			len--
			node := l.Front()
			l.Remove(node)
			treeNode := node.Value.(*TreeNode)
			treeValue := treeNode.Val
			if treeNode.Left != nil {
				l.PushBack(treeNode.Left)
			}
			if treeNode.Right != nil {
				l.PushBack(treeNode.Right)

			}
			if len == 0 {
				results = append(results, treeValue)
			}
		}
	}

	return
}
