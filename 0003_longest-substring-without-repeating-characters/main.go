package main

func lengthOfLongestSubstring(s string) int {
	var i, longest int = 0, 0
	mymap := make(map[byte]int)

	for j := 0; j < len(s); j++ {
		mymap[s[j]]++
		for mymap[s[j]] == 2 && i < j {
			mymap[s[i]]--
			i++
		}

		if longest < j-i+1 {
			longest = j - i + 1
		}
	}

	return longest
}
