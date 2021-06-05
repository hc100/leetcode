package main

import (
	"testing"
)

func Test1Success(t *testing.T) {
	result := convert("PAYPALISHIRING", 3)
	if result != "PAHNAPLSIIGYIR" {
		t.Fatal("failed test")
	}
}

func Test2Success(t *testing.T) {
	result := convert("PAYPALISHIRING", 4)
	if result != "PINALSIGYAHRPI" {
		t.Fatal("failed test")
	}
}

func Test3Success(t *testing.T) {
	result := convert("A", 1)
	if result != "A" {
		t.Fatal("failed test")
	}
}
