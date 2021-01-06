package subsets

import "sort"

// subsets
// https://leetcode-cn.com/problems/subsets/
func subsets(nums []int) [][]int {
	result := [][]int{{}}

	for _, num := range nums {
		newResults := [][]int{}
		for i := range result {
			newResults = append(newResults, result[i])
		}
		for i := range result {
			newR := append([]int{num}, result[i]...)
			newResults = append(newResults, newR)
		}
		result = newResults
	}

	return result
}

// subsets with duplication
// https://leetcode-cn.com/problems/subsets-ii/
func subsetsWithDup(nums []int) [][]int {
	result := [][]int{{}}
	// 对nums进行排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	lastExpand := 0
	for i, num := range nums {
		newResults := [][]int{}
		for j := range result {
			newResults = append(newResults, result[j])
		}
		if i > 0 && num == nums[i-1] {
			start := len(result) - lastExpand
			result = result[start:]
		}
		lastExpand = 0
		for j := range result {
			newR := append(result[j], num)
			newResults = append(newResults, newR)
			lastExpand += 1
		}
		result = newResults
	}

	return result
}

// Permutations
// https://leetcode-cn.com/problems/permutations/
// 标答里面有用到queue，可以节约很大一部分空间
// 使用交换的概念可以提高效率
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	result := [][]int{{
		nums[0],
	}}
	for index, num := range nums {
		if index == 0 {
			continue
		}
		newResult := [][]int{}
		for i := range result {
			r := result[i] // [1,2]
			l := len(r)
			for j := 0; j < l; j++ {
				// 两种情况：
				// 第一个，插两边
				// 其他，插右边位置
				switch j {
				case 0:
					// [1,2] 要查到0位置
					newR := append([]int{num}, r...)
					newResult = append(newResult, newR)
					newR2 := append([]int{r[0], num})
					newR2 = append(newR2, r[1:]...)
					newResult = append(newResult, newR2)
				default:
					// [1,2] 要插入1的位置
					newR := append(r[:j+1], num)
					newR = append(newR, r[j+1:]...)
					newResult = append(newResult, newR)
				}

			}
		}
		result = newResult
	}
	return result
}

// 以下是DFS的方式
func permute_DFS(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	var result [][]int

	l := len(nums)
	var swap func(i int, nums []int)
	swap = func(i int, ns []int) {
		if i == l {
			result = append(result, ns)
		}
		for j := i; j < l; j++ {
			newNums := make([]int, l) // deep copy
			copy(newNums, ns)
			newNums[i], newNums[j] = newNums[j], newNums[i]
			swap(i+1, newNums)
		}
	}
	swap(0, nums)

	return result
}

// String Permutations by changing case
// https://leetcode-cn.com/problems/letter-case-permutation/
func letterCasePermutation(S string) []string {
	// 关键是golang中如何指定将某一个字节变成大写
	cIndex := []int{}
	for index, c := range S {
		if c > '9' || c < '0' {
			cIndex = append(cIndex, index)
		}
	}

	result := []string{S}
	for _, ci := range cIndex {
		newResult := []string{}
		for _, r := range result {
			newR := change(r, ci)
			newResult = append(newResult, newR)
		}
		result = append(result, newResult...)
	}

	return result
}

func change(s string, index int) string {
	bs := []byte(s)
	var newbs []byte
	for i, b := range bs {
		if i == index {
			if b <= 'Z' && b >= 'A' {
				newbs = append(newbs, b+32)
			} else {
				newbs = append(newbs, b-32)
			}
		} else {
			newbs = append(newbs, b)
		}
	}
	return string(newbs)
}

// Structurally Unique Binary Search Tree
// https://leetcode-cn.com/problems/unique-binary-search-trees/
// 这个解题思路和我的思路并不一样，是按照分治递归的方式解决的，而不是BFS的方式
func numTrees(n int) int {
	if n <= 1 {
		return 1
	}
	count := 0
	for i := 1; i <= n; i++ {
		cL := numTrees(i - 1)
		cR := numTrees(n - i)
		count += (cL * cR)
	}
	return count
}
