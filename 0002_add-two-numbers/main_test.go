package main

import (
	"reflect"
	"testing"
)

func TestExample1Success(t *testing.T) {
	p1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	p2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	output := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}
	result := addTwoNumbers(p1, p2)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample2Success(t *testing.T) {
	p1 := &ListNode{0, nil}
	p2 := &ListNode{0, nil}
	output := &ListNode{0, nil}
	result := addTwoNumbers(p1, p2)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample3Success(t *testing.T) {
	p1 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	p2 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	output := &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}
	result := addTwoNumbers(p1, p2)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample4Success(t *testing.T) {
	p1 := &ListNode{1, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}
	p2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	output := &ListNode{6, &ListNode{6, &ListNode{4, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}
	result := addTwoNumbers(p1, p2)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}
