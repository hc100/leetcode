package main

import (
	"fmt"
	"strconv"
)

// ListNode list node
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	next := true
	s := ""
	for next {
		s = s + strconv.Itoa(head.Val)
		if head.Next == nil {
			next = false
		} else {
			head = head.Next
		}
	}
	i, _ := strconv.ParseInt(s, 2, 32)

	return int(i)
}
