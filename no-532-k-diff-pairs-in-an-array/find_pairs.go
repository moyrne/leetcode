package no_532_k_diff_pairs_in_an_array

// https://leetcode.cn/problems/k-diff-pairs-in-an-array/
/*
532. 数组中的 k-diff 数对
给你一个整数数组 nums 和一个整数 k，请你在数组中找出 不同的 k-diff 数对，并返回不同的 k-diff 数对 的数目。
k-diff 数对定义为一个整数对 (nums[i], nums[j]) ，并满足下述全部条件：
0 <= i, j < nums.length
i != j
nums[i] - nums[j] == k
注意，|val| 表示 val 的绝对值。

示例 1：
输入：nums = [3, 1, 4, 1, 5], k = 2
输出：2
解释：数组中有两个 2-diff 数对, (1, 3) 和 (3, 5)。
尽管数组中有两个 1 ，但我们只应返回不同的数对的数量。

示例 2：
输入：nums = [1, 2, 3, 4, 5], k = 1
输出：4
解释：数组中有四个 1-diff 数对, (1, 2), (2, 3), (3, 4) 和 (4, 5) 。

示例 3：
输入：nums = [1, 3, 1, 5, 4], k = 0
输出：1
解释：数组中只有一个 0-diff 数对，(1, 1) 。
*/

func findPairs(nums []int, k int) int {
	// [sub]index
	finder := map[int]int{}
	for i, num := range nums {
		finder[num] = i
	}

	duplicate := map[int]int{}
	var ans int
	for i, num := range nums {
		sub := num - k
		if sub < 0 {
			sub = k + num
		}

		existSub1, ok1 := duplicate[num]
		existSub2, ok2 := duplicate[sub]
		if (ok1 && existSub1 == sub) || (ok2 && existSub2 == num) {
			continue
		}

		idx, ok := finder[sub]
		if !ok || idx == i {
			continue
		}

		duplicate[num] = sub
		ans++
	}

	return ans
}
