package main

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := listNode2Slice(l1)
	s2 := listNode2Slice(l2)

	max := len(s1)
	longer := l1
	shorter := l2
	if len(s2) > len(s1) {
		max = len(s2)
		longer = l2
		shorter = l1
	}

	var result *ListNode
	inc := 0
	for i := 0; i < max; i++ {
		if shorter != nil {
			v := longer.Val + shorter.Val + inc
			if v >= 10 {
				inc = 1
				v = v - 10
			} else {
				inc = 0
			}
			result = &ListNode{v, result}
			longer = longer.Next
			shorter = shorter.Next
		} else {
			v := longer.Val + inc
			if v >= 10 {
				inc = 1
				v = v - 10
			} else {
				inc = 0
			}
			result = &ListNode{v, result}
			longer = longer.Next
		}
	}

	if inc == 1 {
		result = &ListNode{1, result}
	}

	result = reverseLn(result)
	return result
}

func reverseLn(ln *ListNode) *ListNode {
	var l *ListNode
	for ln != nil {
		l = &ListNode{ln.Val, l}
		ln = ln.Next
	}
	return l
}

func listNode2Slice(ln *ListNode) []int {
	r := []int{}
	for ln != nil {
		r = append(r, ln.Val)
		ln = ln.Next
	}
	return r
}
