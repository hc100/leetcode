package main

func arrayNesting(nums []int) int {
	visited := make(map[int]bool, len(nums))
	max := 0

	for i, _ := range nums {
		if !visited[i] {
			start := i
			count := 0
			dup := false
			dups := make(map[int]bool, len(nums))

			for !dup {
				count = count + 1
				dups[start] = true
				visited[start] = true
				start = nextIndex(nums, start)
				if dups[start] {
					dup = true
				}
			}

			if count > max {
				max = count
			}
		}
	}

	return max
}

func nextIndex(nums []int, i int) int {
	return nums[i]
}
