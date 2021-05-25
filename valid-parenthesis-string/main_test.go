package main

import (
	"testing"
)

func TestExample1Success(t *testing.T) {
	result := checkValidString("()")
	if !result {
		t.Fatal("failed test")
	}
}

func TestExample2Success(t *testing.T) {
	result := checkValidString("(*)")
	if !result {
		t.Fatal("failed test")
	}
}

func TestExample3Success(t *testing.T) {
	result := checkValidString("(*))")
	if !result {
		t.Fatal("failed test")
	}
}

func TestExample4Failure(t *testing.T) {
	result := checkValidString("*))")
	if result {
		t.Fatal("failed test")
	}
}
