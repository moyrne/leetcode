package no_704_binary_search

// https://leetcode-cn.com/problems/binary-search/
/*
704. 二分查找
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

示例 1:
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
*/

func search(nums []int, target int) int {
	n := len(nums)
	i, j := 0, n-1
	for i <= j {
		idx := (i + j) / 2
		number := nums[idx]
		switch {
		case number == target:
			return idx
		case number < target:
			i = idx + 1
		default:
			j = idx - 1
		}
	}
	return -1
}
