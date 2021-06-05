package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	r := ""
	row := 1
	for row <= numRows {
		step := (numRows - 1) * 2
		if row == 1 {
			i := 0
			for len(s) > i {
				r = r + string(s[i])
				i = i + step
			}
		} else if row == numRows {
			i := numRows - 1
			for len(s) > i {
				r = r + string(s[i])
				i = i + step
			}
		} else {
			i := row - 1
			trips := (numRows - 1) * 2
			trip1 := trips - ((row - 1) * 2)
			for len(s) > i {
				r = r + string(s[i])
				if len(s) > (i + trip1) {
					r = r + string(s[i+trip1])
				}
				i = i + trips
			}
		}
		row = row + 1
	}

	return r
}
