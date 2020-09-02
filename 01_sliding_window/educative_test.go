package sliding_window

import (
	"fmt"
	"testing"
)

// FindAveragesOfSubArrays
func FindAveragesOfSubArrays(K int, arr []int) []float64 {
	results := []float64{}
	windowsSum, windowsStart := 0, 0
	for windowsEnd := range arr {
		windowsSum += arr[windowsEnd]
		if windowsEnd >= K-1 {
			result := float64(windowsSum) / float64(K)
			results = append(results, result)
			windowsSum -= arr[windowsStart]
			windowsStart += 1
		}
	}
	return results
}

func TestFindAveragesOfSubArrays(t *testing.T) {
	// 时间复杂度为O(n), 因为每一个元素都只被加减1次
	results := FindAveragesOfSubArrays(5, []int{1, 3, 2, 6, -1, 4, 1, 8, 2})
	t.Log("Averages of subarrays of size K: ", results)
}

// FindMaxSubArrayOfSizeK
func FindMaxSubArrayOfSizeK(K int, arr []int) int {
	max := 0
	windowsSum, windowsStart := 0, 0
	for windowsEnd := range arr {
		windowsSum += arr[windowsEnd]
		if windowsEnd >= K-1 {
			if windowsSum > max {
				max = windowsSum
			}
			windowsSum -= arr[windowsStart]
			windowsStart += 1
		}
	}
	return max
}

func TestFindMaxSubArrayOfSizeK(t *testing.T) {
	t.Log(FindMaxSubArrayOfSizeK(2, []int{2, 3, 4, 1, 5})) // 预期结果为7
}

// SmallestSubarrayWithGivenSum
func SmallestSubarrayWithGivenSum(arr []int, s int) int {
	sumShortest := len(arr) + 1
	sumLen := 0
	windowsSum, windowsStart := 0, 0

	for windowsEnd := range arr {
		windowsSum += arr[windowsEnd]
		sumLen += 1
		for windowsSum >= s {
			if sumLen < sumShortest {
				sumShortest = sumLen
			}
			sumLen -= 1
			windowsSum -= arr[windowsStart]
			windowsStart += 1
		}
	}

	if sumShortest == len(arr)+1 {
		return 0
	}

	return sumShortest
}

func TestSmallestSubarrayWithGivenSum(t *testing.T) {
	t.Log(SmallestSubarrayWithGivenSum([]int{2, 1, 5, 2, 3, 2}, 7)) // 预期：2
	t.Log(SmallestSubarrayWithGivenSum([]int{2, 1, 5, 2, 8}, 7))    // 预期：1
	t.Log(SmallestSubarrayWithGivenSum([]int{3, 4, 1, 1, 6}, 8))    // 预期：3
}

// LongestSubstringKDistinct
func LongestSubstringKDistinct(s string, k int) int {
	longest := 0
	windowsStart := 0
	windows := make(map[byte]int)
	for windowsEnd := range s {
		windows[s[windowsEnd]] += 1
		if len(windows) <= k && windowsEnd-windowsStart+1 > longest {
			longest = windowsEnd - windowsStart + 1
		}
		for len(windows) > k && windowsStart <= windowsEnd {
			if v, ok := windows[s[windowsStart]]; ok && v == 1 {
				delete(windows, s[windowsStart])
			} else {
				windows[s[windowsStart]] -= 1
			}
			windowsStart += 1
		}
	}
	return longest
}

func TestLongestSubstringKDistinct(t *testing.T) {
	t.Log(LongestSubstringKDistinct("araaci", 2)) // 预期：4
	t.Log(LongestSubstringKDistinct("raa", 1))    // 预期：2
	t.Log(LongestSubstringKDistinct("cbbebi", 3)) // 预期：5
}

// Fruits into Baskets
func FruitsIntoBaskets(fruits []byte) int {
	longest := 0
	windowsMap := make(map[byte]int)
	windowsStart := 0
	for windowsEnd, f := range fruits {
		windowsMap[f] += 1
		if len(windowsMap) <= 2 && longest < windowsEnd-windowsStart+1 {
			longest = windowsEnd - windowsStart + 1
		}
		for len(windowsMap) > 2 {
			if v, ok := windowsMap[fruits[windowsStart]]; ok && v == 1 {
				delete(windowsMap, fruits[windowsStart])
			} else {
				windowsMap[fruits[windowsStart]] -= 1
			}
			windowsStart += 1
		}
	}
	return longest
}

func TestFruitsIntoBaskets(t *testing.T) {
	t.Log(FruitsIntoBaskets([]byte{'A', 'B', 'C', 'A', 'C'}))      // 预期：3
	t.Log(FruitsIntoBaskets([]byte{'A', 'B', 'C', 'B', 'B', 'C'})) // 预期：5
}

// NoRepeatSubstring
func NoRepeatSubstring(s string) int {
	windowsStart := 0
	windowsMap := make(map[byte]int) // key是字符，int是字符在字符串中的位置
	maxLength := 0
	for windowsEnd := range s {
		rightChar := s[windowsEnd] // 也就是滑窗右边的值
		// 如果这个rightChar已经再windows中存在了，那么就执行shrink操作
		if index, ok := windowsMap[rightChar]; ok {
			// 此处的shrink也是一个tricky，和之前的问题不同，此处的windowsStart可以直接的shrink到字符串中，上一个rightChar出现的位置的下一个位置
			// 此处还有一个tricky，我们没有实际的在map讲shrink过的字符删除，判断一个字符是否被删除，只要判断它的index是否大于windowsStart即可
			if windowsStart <= index {
				windowsStart = index + 1
			}
		}
		windowsMap[rightChar] = windowsEnd
		if windowsEnd-windowsStart+1 > maxLength {
			maxLength = windowsEnd - windowsStart + 1
		}
	}
	return maxLength
}

func TestNoRepeatSubstring(t *testing.T) {
	t.Log(NoRepeatSubstring("aabccbb")) // 预期：3
	t.Log(NoRepeatSubstring("abbbb"))   // 预期：2
	t.Log(NoRepeatSubstring("abccde"))  // 预期：3
}

// CharacterReplacement
// leetcode的一道类似的题 https://leetcode-cn.com/problems/longest-repeating-character-replacement/solution/tong-guo-ci-ti-liao-jie-yi-xia-shi-yao-shi-hua-don/
// 本题难点在于： maxRepeatLetterNumber 这个变量，一开始容易理解为 当前窗口中的最大重复字符数； 但实际这个变量的含义为，所有满足条件的窗口中，最大重复字符数
// 在本题中，因为我们只对“最长，有效的子字符串”感兴趣，所以窗口其实并没有shrink（收缩），严格来讲除了expand（扩展），就是做了shift（平移，即整体向右边移动的一格）
// 而平移过程中，当前窗口可能会覆盖到“无效”的子字符串（即不满足 windowsLength - maxRepeatChar <= k）
// 按理来收，每次平移之后，需要重新计算当前windows中的maxRepeatedChar，但是maxrepeatChar准确来说是历史上最大的重复值
// shrink 仅会发生在 maxRepeatChar 没有更新（即变得更大的时候), 且当前窗口长度-maxRepeatChar > k 的时候
func CharacterReplacement(s string, k int) int {
	windowsStart := 0
	longest := 0
	maxRepeatLetterNumber := 0
	windowsMap := make(map[byte]int) // key: 字符串中的字符 value: 字符在window中出现的次数
	sbyte := []byte(s)
	for windowEnd, b := range sbyte {
		windowsMap[b] += 1
		if windowsMap[b] >= maxRepeatLetterNumber {
			maxRepeatLetterNumber = windowsMap[b]
		}

		// 否则则有机会执行expand
		if windowEnd-windowsStart+1-maxRepeatLetterNumber > k {
			windowsMap[sbyte[windowsStart]] -= 1
			windowsStart += 1
			continue
		}

		// 仅当expand的时候，longest可能会更新
		if longest < windowEnd-windowsStart+1 {
			longest = windowEnd - windowsStart + 1
		}
	}
	return longest
}

func TestCharacterReplacement(t *testing.T) {
	t.Log(CharacterReplacement("aabccbb", 2)) // 预期：5
	t.Log(CharacterReplacement("abbcb", 1))   // 预期：4
	t.Log(CharacterReplacement("abccde", 1))  // 预期：3
	t.Log(CharacterReplacement("baaab", 2))   // 预期：5
}

// LengthOfLongestSubstring
func LengthOfLongestSubstring(arr []int, k int) int {
	longest := 0
	windowStart := 0
	frequecyOfZero := 0
	for windowEnd := range arr {
		if arr[windowEnd] == 0 {
			frequecyOfZero += 1
		}
		for frequecyOfZero > k {
			if arr[windowStart] == 0 {
				frequecyOfZero -= 1
			}
			windowStart += 1
		}
		if windowEnd-windowStart+1 > k {
			longest = windowEnd - windowStart + 1
		}
	}
	return longest
}

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(LengthOfLongestSubstring([]int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}, 2))
	t.Log(LengthOfLongestSubstring([]int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}, 3))
}

// Problem Challenge 1 : FindPermutation
func FindPermutation(s1, s2 string) bool {
	// init
	exist := make(map[byte]bool)
	frequency := make(map[byte]int)
	completed := make(map[byte]struct{})
	windowStart := 0
	sbyte1 := []byte(s1)
	sbyte2 := []byte(s2)

	for _, b := range sbyte2 {
		exist[b] = true
		frequency[b] += 1
		completed[b] = struct{}{}
	}

	for windowEnd := range sbyte1 {
		if exist[sbyte1[windowEnd]] {
			frequency[sbyte1[windowEnd]] -= 1
			if frequency[sbyte1[windowEnd]] == 0 {
				delete(completed, sbyte1[windowEnd])
			}
		}

		if windowEnd < len(sbyte2) {
			continue
		}

		if exist[sbyte1[windowStart]] {
			frequency[sbyte1[windowStart]] += 1
			if frequency[sbyte1[windowStart]] == 1 {
				completed[sbyte1[windowStart]] = struct{}{}
			}
		}
		windowStart += 1

		if len(completed) == 0 {
			return true
		}
	}

	if len(completed) == 0 {
		return true
	}

	return false
}

// 官方做法基本和我的做法相同，重点在于最后的统计
// 我是通过一个奇怪的三重map的组成，而其实只需要一个map外加一个match计数
// 用match的大小和子串的大小做比较，判断是否全匹配
func FindPermutation2(s1, s2 string) bool {
	// init
	frequency := make(map[byte]int)
	windowStart, match := 0, 0
	sbyte1 := []byte(s1)
	sbyte2 := []byte(s2)

	for _, b := range sbyte2 {
		frequency[b] += 1
	}

	for windowEnd := range sbyte1 {
		rightChar := sbyte1[windowEnd]
		if _, ok := frequency[rightChar]; ok {
			frequency[rightChar] -= 1
			if frequency[rightChar] == 0 {
				match += 1
			}
		}

		// 我喜欢这里的技巧，把对match的比较提前
		// 在routine中，match总能移动到一个适当的位置，以免被多写几次
		// expand - shrink - compare 也可以改写为 expand - compare - shrink 结果是不变的，但是逻辑看起来清楚一些
		if match == len(frequency) { // 这里是和frequency比较，而不是len of s2
			return true
		}

		if windowEnd >= len(sbyte2)-1 {
			leftChar := sbyte1[windowStart]
			if _, ok := frequency[leftChar]; ok {
				frequency[leftChar] += 1
				if frequency[leftChar] == 1 {
					match -= 1
				}
			}
			windowStart += 1
		}
	}

	return false
}

func TestFindPermutation(t *testing.T) {
	fmt.Println(FindPermutation2("oidbcaf", "abc"))         // except : true
	fmt.Println(FindPermutation2("odicf", "dc"))            // except : false
	fmt.Println(FindPermutation2("abc", "abc"))             // except: true
	fmt.Println(FindPermutation2("bcdxabcdy", "bcdxabcdy")) // except: true
	fmt.Println(FindPermutation2("aaacb", "abc"))           // except : true
}

// Problem Chanlenge : StringAnagrams
// 本题基本就是上一个题的延申，只要把结果从返回一个改为返回多个即可
func StringAnagrams(str, pattern string) []int {
	// init
	frequency := make(map[byte]int)
	windowStart, match := 0, 0
	sbyte1 := []byte(str)
	sbyte2 := []byte(pattern)
	result := []int{}

	for _, b := range sbyte2 {
		frequency[b] += 1
	}

	for windowEnd := range sbyte1 {
		rightChar := sbyte1[windowEnd]
		if _, ok := frequency[rightChar]; ok {
			frequency[rightChar] -= 1
			if frequency[rightChar] == 0 {
				match += 1
			}
		}

		// 我喜欢这里的技巧，把对match的比较提前
		// 在routine中，match总能移动到一个适当的位置，以免被多写几次
		// expand - shrink - compare 也可以改写为 expand - compare - shrink 结果是不变的，但是逻辑看起来清楚一些
		if match == len(frequency) { // 这里是和frequency比较，而不是len of s2
			result = append(result, windowStart)
		}

		if windowEnd >= len(sbyte2)-1 {
			leftChar := sbyte1[windowStart]
			if _, ok := frequency[leftChar]; ok {
				frequency[leftChar] += 1
				if frequency[leftChar] == 1 {
					match -= 1
				}
			}
			windowStart += 1
		}
	}

	return result
}

func TestStringAnagrams(t *testing.T) {
	fmt.Println(StringAnagrams("ppqp", "pq"))     // [1,2]
	fmt.Println(StringAnagrams("abbcabc", "abc")) // [2,3,4]
}

// MinimumWindowSubstring
func MinimumWindowSubstring(str, pattern string) string {
	frequency := make(map[byte]int)
	match := 0
	smallest := len(str) + 1
	result := ""

	// init from pattern
	patternBytes := []byte(pattern)
	for _, b := range patternBytes {
		frequency[b] += 1
	}

	winStart := 0
	strBytes := []byte(str)
	for winEnd := range strBytes {
		if _, ok := frequency[strBytes[winEnd]]; ok {
			frequency[strBytes[winEnd]] -= 1
			if frequency[strBytes[winEnd]] == 0 {
				match += 1
			}
		}

		// 直到match，才能进行shrink
		if match < len(frequency) {
			continue
		}

		for {
			if match < len(frequency) {
				break
			}
			if smallest > winEnd-winStart+1 {
				smallest = winEnd - winStart + 1
				result = str[winStart : winEnd+1]
			}
			if _, ok := frequency[strBytes[winStart]]; ok {
				if frequency[strBytes[winStart]] == 0 {
					match -= 1
				}
				frequency[strBytes[winStart]] += 1
			}
			winStart += 1
		}
	}

	return result
}

// 本题和 leetcode https://leetcode-cn.com/problems/minimum-window-substring/ 相同，以上算法以通过leetcode的测例
func TestMinimumWindowSubstring(t *testing.T) {
	fmt.Println(MinimumWindowSubstring("aabdec", "abc"))  // abdec
	fmt.Println(MinimumWindowSubstring("abdabca", "abc")) // abc
	fmt.Println(MinimumWindowSubstring("adcad", "abc"))   // 空字符串
}

// WordsConcatenation
// 这个题即使是官方给的解法，时间也是O(N*M*Len)
// 此题种 words 可能含有重复的
func WordsConcatenation(str string, words []string) []int {
	// 首先判断边际情况
	if len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}

	// str的长度本身不够
	wordsCount := len(words)
	wordLen := len(words[0])
	if len(str) < wordsCount*wordLen {
		return []int{}
	}

	// 初始化map
	result := []int{}
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word] += 1
	}

	for i := 0; i <= len(str)-wordLen*wordsCount; i++ {
		match := make(map[string]int)
		for j := 0; j < wordsCount; j++ {
			winStart := i + j*wordLen
			winEnd := winStart + wordLen

			// 找到当前的子串
			subStr := str[winStart:winEnd]

			if _, ok := wordsMap[subStr]; !ok {
				// 如果当前字串不存在，则直接返回
				break
			}

			match[subStr] += 1
			if match[subStr] > wordsMap[subStr] {
				break
			}

			if j == wordsCount-1 {
				result = append(result, i)
			}
		}
	}

	return result
}

func TestWordsConcatenation(t *testing.T) {
	fmt.Println(WordsConcatenation("catfoxcat", []string{"cat", "fox"}))    // [0,3]
	fmt.Println(WordsConcatenation("catcatfoxfox", []string{"cat", "fox"})) // [3]
}
