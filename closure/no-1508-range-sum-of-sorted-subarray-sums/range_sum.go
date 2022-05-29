package no_1508_range_sum_of_sorted_subarray_sums

import (
	"sort"
)

var mod int = 1e9 + 7

func rangeSum(nums []int, n int, left int, right int) int {
	newNums := make([]int, 0, n*(n+1)/2)
	for start := 0; start < len(nums); start++ {
		var sum int
		for end := start; end < len(nums); end++ {
			sum += nums[end]
			sum %= mod
			newNums = append(newNums, sum)
		}
	}

	sort.Ints(newNums)

	var ans int
	for i := left - 1; i < right; i++ {
		ans += newNums[i]
		ans %= mod
	}

	return ans
}
