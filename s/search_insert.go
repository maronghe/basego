/*
 * Copyright (c) 2020 RongHe Ma.
 * Recording from daily work flow.
 */

package s

func SearchInsert(nums []int, target int) int {
	low, high := 0 ,len(nums) - 1
	for low <= high {
		mid := low + (high - low) / 2
		if nums[mid] == target { return mid } else
		if nums[mid] > target { high = mid - 1 } else { low = mid + 1 }
	}
	return low
}
