package topological_sort

import (
	"container/list"
	"testing"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 首先建立图 和 inDegree计数
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for i := 0; i < numCourses; i++ {
		inDegree[i] = 0
		graph[i] = []int{}
	}

	for _, prerequisite := range prerequisites {
		child := prerequisite[0]
		parent := prerequisite[1]

		inDegree[child] += 1
		graph[parent] = append(graph[parent], child)
	}

	// 初始化queue
	queue := list.New()
	for num, degree := range inDegree {
		if degree == 0 {
			queue.PushBack(num)
		}
	}

	var count int
	for queue.Len() > 0 {
		e := queue.Front()
		num := e.Value.(int)
		count += 1

		for _, child := range graph[num] {
			inDegree[child] -= 1
			if inDegree[child] == 0 {
				queue.PushBack(child)
			}
		}

		queue.Remove(e)
	}

	return count == numCourses
}

func TestCanFinish(t *testing.T) {
	// canFinish(2, [][]int{{0, 1}, {1, 0}})
	canFinish(2, [][]int{{0, 1}})
}
