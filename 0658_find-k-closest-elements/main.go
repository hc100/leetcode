package main

import (
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	startIndex := 0
	if x <= arr[0] {
		return arr[0:k]
	} else if x >= arr[len(arr)-1] {
		return arr[len(arr)-k : len(arr)]
	} else {
		lastDiff := x - arr[0]
		for i, v := range arr {
			diff := abs(x - v)
			if diff < lastDiff {
				lastDiff = diff
				startIndex = i
			}
		}

		stack := []int{}
		stack = append(stack, arr[startIndex])
		leftOf := startIndex - 1
		rightOf := startIndex + 1

		for len(stack) < k {
			var leftValue int
			var rightValue int

			if leftOf < 0 {
				stack = append(stack, arr[rightOf])
				rightOf = rightOf + 1
			} else if rightOf > len(arr) {
				stack = append(stack, arr[leftOf])
				leftOf = leftOf - 1
			} else {
				if rightOf >= len(arr) {
					leftValue = arr[leftOf]
					stack = append(stack, leftValue)
					leftOf = leftOf - 1
				} else if leftOf < 0 {
					rightValue = arr[rightOf]
					stack = append(stack, rightValue)
					rightOf = rightOf + 1
				} else {
					leftValue = arr[leftOf]
					rightValue = arr[rightOf]
					if dist(x, leftValue) == dist(x, rightValue) {
						stack = append(stack, leftValue)
						leftOf = leftOf - 1
					} else if dist(x, leftValue) < dist(x, rightValue) {
						stack = append(stack, leftValue)
						leftOf = leftOf - 1
					} else if dist(x, leftValue) > dist(x, rightValue) {
						stack = append(stack, rightValue)
						rightOf = rightOf + 1
					}
				}

			}
		}

		sort.Ints(stack)
		return stack
	}
}

func dist(x int, from int) int {
	return abs(from - x)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
