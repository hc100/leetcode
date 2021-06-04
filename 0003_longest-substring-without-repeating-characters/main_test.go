package main

import (
	"testing"
)

func Test1Success(t *testing.T) {
	result := lengthOfLongestSubstring("abcabcbb")
	if result != 3 {
		t.Fatal("failed test")
	}
}

func Test2Success(t *testing.T) {
	result := lengthOfLongestSubstring("bbbbb")
	if result != 1 {
		t.Fatal("failed test")
	}
}

func Test3Success(t *testing.T) {
	result := lengthOfLongestSubstring("pwwkew")
	if result != 3 {
		t.Fatal("failed test")
	}
}

func Test4Success(t *testing.T) {
	result := lengthOfLongestSubstring("")
	if result != 0 {
		t.Fatal("failed test")
	}
}

func Test5Success(t *testing.T) {
	result := lengthOfLongestSubstring("abababababa")
	if result != 2 {
		t.Fatal("failed test")
	}
}

func Test6Success(t *testing.T) {
	result := lengthOfLongestSubstring("abcdefghijklmnopqrstuvw")
	if result != 23 {
		t.Fatal("failed test")
	}
}

func Test7Success(t *testing.T) {
	result := lengthOfLongestSubstring("a")
	if result != 1 {
		t.Fatal("failed test")
	}
}

func Test8Success(t *testing.T) {
	result := lengthOfLongestSubstring("dvdf")
	if result != 3 {
		t.Fatal("failed test")
	}
}

func Test9Success(t *testing.T) {
	result := lengthOfLongestSubstring("ckilbkd")
	if result != 5 {
		t.Fatal("failed test")
	}
}

func Test10Success(t *testing.T) {
	result := lengthOfLongestSubstring("ckilbikd")
	if result != 5 {
		t.Fatal("failed test")
	}
}

func Test11Success(t *testing.T) {
	result := lengthOfLongestSubstring("ccccccccccckilbikd")
	if result != 5 {
		t.Fatal("failed test")
	}
}

func Test12Success(t *testing.T) {
	result := lengthOfLongestSubstring("cckilbikkkkkkkd")
	if result != 5 {
		t.Fatal("failed test")
	}
}

func Test13Success(t *testing.T) {
	result := lengthOfLongestSubstring("tmmzuxt")
	if result != 5 {
		t.Fatal("failed test")
	}
}

func Test14Success(t *testing.T) {
	result := lengthOfLongestSubstring("wobgrovw")
	if result != 6 {
		t.Fatal("failed test")
	}
}
