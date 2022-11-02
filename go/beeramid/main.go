package beeramid

import "math"

func Beeramid(bonus int, price float64) int {
	var level int
	if bonus < 0 || price <= 0 {
		return level
	}
	bottles := float64(bonus) / price

	for i := 1; bottles > 0; i++ {
		bottles -= math.Pow(float64(i), 2)
		if bottles >= 0 {
			level = i
		}
	}

	return level
}
