package main

func twoSum(nums []int, target int) []int {
	// Brute Force
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if (nums[i] + nums[j]) == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }

	// One-pass Hash Table
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if val, ok := m[complement]; ok {
			return []int{val, i}
		}
		m[nums[i]] = i
	}

	panic("No two sum solution")
}
