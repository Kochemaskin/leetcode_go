package main

import (
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	n := len(nums)
	result := make([][]int, 0)

	sort.Ints(nums)

	for i := 0; i < n; i += 3 {
		if i+2 < n && nums[i+2]-nums[i] <= k {
			result = append(result, []int{nums[i], nums[i+1], nums[i+2]})
		} else {
			return [][]int{}
		}
	}

	return result
}

/*
TC O(n log n)
SC O(n)
*/
