package main

import (
	"testing"
)

func TestExample1Success(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	result := getDecimalValue(head)
	if result != 5 {
		t.Fatal("failed test")
	}
}

func TestExample2Success(t *testing.T) {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	result := getDecimalValue(head)
	if result != 0 {
		t.Fatal("failed test")
	}
}

func TestExample3Success(t *testing.T) {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	result := getDecimalValue(head)
	if result != 1 {
		t.Fatal("failed test")
	}
}

func TestExample4Failure(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 1,
								Next: &ListNode{
									Val: 1,
									Next: &ListNode{
										Val: 1,
										Next: &ListNode{
											Val: 0,
											Next: &ListNode{
												Val: 0,
												Next: &ListNode{
													Val: 0,
													Next: &ListNode{
														Val: 0,
														Next: &ListNode{
															Val: 0,
															Next: &ListNode{
																Val:  0,
																Next: nil,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	result := getDecimalValue(head)
	if result != 18880 {
		t.Fatal("failed test")
	}
}
