package no_448_find_all_numbers_disappeared_in_an_array

// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
/*
448. 找到所有数组中消失的数字
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

示例 1：
输入：nums = [4,3,2,7,8,2,3,1]
输出：[5,6]

示例 2：
输入：nums = [1,1]
输出：[2]

提示：
n == nums.length
1 <= n <= 105
1 <= nums[i] <= n
*/

func findDisappearedNumbers(nums []int) []int {
	dump := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dump[nums[i]-1] = true
	}

	var ans []int
	for i := 0; i < len(nums); i++ {
		if !dump[i] {
			ans = append(ans, i+1)
		}
	}
	return ans
}
