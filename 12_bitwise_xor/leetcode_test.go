package bitwise_xor

// https://leetcode-cn.com/problems/single-number-iii/
func singleNumber(nums []int) []int {
	bitmask := onesingleNumber(nums)

	diff := bitmask & (-bitmask)

	group1 := []int{}
	group2 := []int{}

	for _, num := range nums {
		if num&diff == 0 {
			group1 = append(group1, num)
		} else {
			group2 = append(group2, num)
		}
	}
	return []int{onesingleNumber(group1), onesingleNumber(group2)}
}

func onesingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// https://leetcode-cn.com/problems/complement-of-base-10-integer/solution/
func bitwiseComplement(num int) int {
	if num == 0 {
		return 1
	}
	bitCount, n := 0, num
	for n > 0 {
		bitCount += 1
		n = n / 2
	}
	allBitSet := 1
	for i := 0; i < bitCount; i++ {
		allBitSet *= 2
	}
	allBitSet -= 1

	return num ^ allBitSet
}
