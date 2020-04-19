/*
 * Copyright (c) 2020 RongHe Ma.
 * Recording from daily work flow.
 */

package s


// time complexity O(logN)
func BinarySearch(nums []int, target int) int {
	i, j := 0 , len(nums) - 1

	for i <= j {
		mid := i + (j - i) / 2
		if target == nums[mid] { return mid } else
		if target > nums[mid] { i = mid + 1 } else { j = mid - 1 }
	}

	return -1
}



// time complexity O(n)
func OSearch(nums []int, target int) int {
	// 1 check params
	if len(nums) == 0 {
		return -1
	}

	// 2 loop
	index := -1
	for k := range nums{
		if nums[k] == target{
			index = k
			break
		}
	}

	return index
}