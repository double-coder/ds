package main

func generate(numRows int) [][]int {
	/*
		PASCAL TRIANGLE
		1
		11
		121
		1331
	*/
	res := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
	}
	// rrow, rcol := 0, 0
	for r := 0; r < len(res); r++ {
		for j := 0; j < len(res[r]); j++ {
			if j == 0 || j == len(res[r])-1 {
				res[r][j] = 1
			} else {
				res[r][j] = res[r-1][j-1] + res[r-1][j]
			}
		}
	}
	return res
}
