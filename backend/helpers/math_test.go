package helpers

import (
	"testing"
)

func TestPopulartiyFactorHighPlaycountLowRank(t *testing.T) {
	rank := uint64(1_000)
	plays := uint64(100_000)
	expected := 0.9

	res := GetPopularityFactor(rank, plays)

	t.Logf("Testing if %d plays and %d rank gives %f", plays, rank, expected)

	if !almostEqual(res, expected) {
		t.Errorf("Expected %f, got %f", expected, res)
	}
}

func TestPopulartiyFactorHighPlaycountHighRank(t *testing.T) {
	rank := uint64(10_000)
	plays := uint64(100_000)
	expected := 0.65

	res := GetPopularityFactor(rank, plays)

	t.Logf("Testing if %d plays and %d rank gives %f", plays, rank, expected)

	if !almostEqual(res, expected) {
		t.Errorf("Expected %f, got %f", expected, res)
	}
}

func TestPopulartiyFactorLowPlaycountLowRank(t *testing.T) {
	rank := uint64(100)
	plays := uint64(1_000)
	expected := 1.0

	res := GetPopularityFactor(rank, plays)

	t.Logf("Testing if %d plays and %d rank gives %f", plays, rank, expected)

	if !almostEqual(res, expected) {
		t.Errorf("Expected %f, got %f", expected, res)
	}
}

func TestPopulartiyFactorLowPlaycountHighRank(t *testing.T) {
	rank := uint64(800)
	plays := uint64(1_000)
	expected := 1.0

	res := GetPopularityFactor(rank, plays)

	t.Logf("Testing if %d plays and %d rank gives %f", plays, rank, expected)

	if !almostEqual(res, expected) {
		t.Errorf("Expected %f, got %f", expected, res)
	}
}

func TestPopulartiyFactorLowPlaycountMediumRank(t *testing.T) {
	rank := uint64(1_000)
	plays := uint64(6_000)
	expected := 0.8166

	res := GetPopularityFactor(rank, plays)

	t.Logf("Testing if %d plays and %d rank gives %f", plays, rank, expected)

	if !almostEqual(res, expected) {
		t.Errorf("Expected %f, got %f", expected, res)
	}
}
