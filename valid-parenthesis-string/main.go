package main

import ()

func checkValidString(s string) bool {
	lo, hi := 0, 0
	lo, hi = possibleRange(lo, hi, s)

	return lo == 0
}

func possibleRange(lo int, hi int, s string) (int, int) {
	for _, r := range s {
		c := string(r)
		if c == "(" {
			lo = lo + 1
		} else {
			lo = lo - 1
		}

		if c != ")" {
			hi = hi + 1
		} else {
			hi = hi - 1
		}

		if hi < 0 {
			return lo, hi
		}

		if lo < 0 {
			lo = 0
		}
	}

	return lo, hi
}
