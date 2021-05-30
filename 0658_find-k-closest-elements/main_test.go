package main

import (
	"reflect"
	"testing"
)

func TestExample1Success(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	output := []int{1, 2, 3, 4}
	result := findClosestElements(input, 4, 3)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample2Success(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	output := []int{1, 2, 3, 4}
	result := findClosestElements(input, 4, -1)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample3Success(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	output := []int{2, 3, 4, 5}
	result := findClosestElements(input, 4, 100)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample4Success(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	output := []int{1}
	result := findClosestElements(input, 1, 1)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample5Success(t *testing.T) {
	input := []int{1, 2, 3, 5, 6}
	output := []int{3}
	result := findClosestElements(input, 1, 4)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample6Success(t *testing.T) {
	input := []int{1, 2, 3, 5, 6}
	output := []int{3, 5}
	result := findClosestElements(input, 2, 4)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}

func TestExample7Success(t *testing.T) {
	input := []int{-2, -1, 1, 2, 3, 4, 5}
	output := []int{-2, -1, 1, 2, 3, 4, 5}
	result := findClosestElements(input, 7, 3)
	if !reflect.DeepEqual(result, output) {
		t.Fatal("failed test")
	}
}
