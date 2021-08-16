package main

import "math"

func ListSquared(m, n int) [][]int {
	var tmp int
	var res = [][]int{}
	for i := m; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				tmp += j * j
			}
		}
		if int(math.Sqrt(float64(tmp)))*int(math.Sqrt(float64(tmp))) == tmp {
			res = append(res, []int{i, tmp})
		}
		tmp = 0
	}

	return res
}
