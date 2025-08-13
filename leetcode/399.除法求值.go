/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 */

// @lc code=start

import (
	"math/big"
)

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	n := 0
	Vmap := make(map[string]int)
	for i := 0; i < len(equations); i++ {
		for j := 0; j < 2; j++ {
			if _, ok := Vmap[equations[i][j]]; !ok {
				Vmap[equations[i][j]] = n
				n++
			}
		}
	}
	roadMap := make([][]*big.Float, n)

	for i := 0; i < n; i++ {
		roadMap[i] = make([]*big.Float, n)
		for j := 0; j < n; j++ {
			roadMap[i][j] = big.NewFloat(-1.0).SetPrec(100)
			if j == i {
				roadMap[i][j] = big.NewFloat(1.0).SetPrec(100)
			}
		}
	}

	for i := 0; i < len(equations); i++ {
		roadMap[Vmap[equations[i][0]]][Vmap[equations[i][1]]] = big.NewFloat(values[i]).SetPrec(100)
		roadMap[Vmap[equations[i][1]]][Vmap[equations[i][0]]] = new(big.Float).Quo(big.NewFloat(1), roadMap[Vmap[equations[i][0]]][Vmap[equations[i][1]]]).SetPrec(100)
	}

	for k := range roadMap {
		for i := range roadMap {
			for j := range roadMap {
				if roadMap[i][k].Cmp(big.NewFloat(0)) == 1 && roadMap[k][j].Cmp(big.NewFloat(0)) == 1 {
					roadMap[i][j] = new(big.Float).Mul(roadMap[i][k], roadMap[k][j]).SetPrec(100)
				}
			}
		}
	}

	result := []float64{}
	for i := 0; i < len(queries); i++ {
		_, ok0 := Vmap[queries[i][0]]
		_, ok1 := Vmap[queries[i][1]]
		if ok0 && ok1 {
			r, _ := roadMap[Vmap[queries[i][0]]][Vmap[queries[i][1]]].Float64()
			result = append(result, r)
		} else {
			result = append(result, -1.0)
		}
	}

	return result
}

// @lc code=end

