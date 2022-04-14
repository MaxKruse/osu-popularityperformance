package helpers

import "math"

const (
	UPPER_BOUND = 100_000.0
	LOWER_BOUND = 1_000.0
)

const float64EqualityThreshold = 1e-4

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func GetPopularityFactor(rank uint64, plays uint64) float64 {
	if plays == 0 || rank == 0 {
		return 1.0
	}

	fRank := float64(rank)
	// Squish fPlays between UPPWER_BOUND and LOWER_BOUND
	fPlays := math.Max(LOWER_BOUND, math.Min(float64(plays), UPPER_BOUND))

	// In what percentage of the rankings are we?
	ratio := (fRank / fPlays)

	res := formula(ratio, fPlays)
	return res
}

func formula(ratio float64, plays float64) float64 {
	// if we are in the upper percentages, dont nerf if plays are low
	playsLog := math.Log10(plays)
	// by design, 3.0 < playsLogs < 5.0

	// so few plays, just give them a 1.0
	if playsLog <= 3.25 {
		return 1.0
	}

	if playsLog <= 4.0 {
		// <10_000 plays on this map. Give good factors
		// Top 10%
		if ratio <= 0.1 {
			return 1.0
		}

		// Top 25%
		if ratio >= 0.25 {
			return 0.75
		}

		// at this point, the ratio is between 0.1 and 0.25
		return 0.75 + (ratio - 0.1)
	}

	// >10_000 plays on this map. Give bad factors
	if ratio <= 0.01 {
		return 0.9
	}

	if ratio >= 0.025 {
		return 0.65
	}

	// at this point, ratio is between 0.01 and 0.025.
	// Normalize ratio
	ratio = (ratio - 0.01) / 0.024

	return ratio
}
