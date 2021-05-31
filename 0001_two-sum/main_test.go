package main

import (
	"reflect"
	"testing"
)

func TestExample1Success(t *testing.T) {
	input := []int{2, 7, 11, 15}
	output := []int{0, 1}
	result := twoSum(input, 9)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample2Success(t *testing.T) {
	input := []int{3, 2, 4}
	output := []int{1, 2}
	result := twoSum(input, 6)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample3Success(t *testing.T) {
	input := []int{3, 3}
	output := []int{0, 1}
	result := twoSum(input, 6)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}
