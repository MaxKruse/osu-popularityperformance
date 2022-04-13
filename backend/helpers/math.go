package helpers

import "math"

const (
	MIN_FACTOR float64 = 0.01
	MAX_FACTOR float64 = 1.01
)

func GetPopularityFactor(rank uint64, plays uint64) float64 {
	if plays == 0 {
		return MIN_FACTOR
	}
	factor := float64(rank) / float64(plays)
	factorRatio := math.Max(MIN_FACTOR, math.Min(factor, MAX_FACTOR))

	res := formula(factorRatio)
	return res
}

func formula(normalizedFactor float64) float64 {
	return 1.0 / ((2.0 - MAX_FACTOR) + normalizedFactor)
}
