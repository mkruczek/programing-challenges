package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturn0ForFFFFFFF(t *testing.T) {
	//explanation
	//[0 - 127] F -> [0 - 63] F -> [0 - 31] F -> [0 - 15] F -> [0 - 7] F -> [0 - 3] F -> [0 - 1] F = 0

	//given
	s := seat{
		row: "FFFFFFF",
	}
	//when
	//then
	assert.Equal(t, 0, s.calculateRow())
}

func TestShouldReturn127ForBBBBBBB(t *testing.T) {
	//explanation TODO check the table
	//[0 - 127] B -> [64 - 127] F -> [94 - 127] F -> [110 - 127] F -> [119 - 127] F -> [123 - 127] F -> [126 - 127] F = 127

	//given
	s := seat{
		row: "BBBBBBB",
	}
	//when
	//then
	assert.Equal(t, 127, s.calculateRow())
}

func TestShouldReturn44ForFBFBBFF(t *testing.T) {
	//explanation
	//[0 - 127] B -> [0 - 63] F -> [32 - 63] F -> [32 - 47] F -> [40 - 47] F -> [44 - 47] F -> [44 - 45] F = 44

	//given
	s := seat{
		row: "FBFBBFF",
	}
	//when
	//then
	assert.Equal(t, 44, s.calculateRow())
}

func TestShouldReturn0ForLLL(t *testing.T) {
	//explanation
	//[0 - 7] L -> [0 - 3] L -> [0 - 1] L = 0

	//given
	s := seat{
		column: "LLL",
	}
	//when
	//then
	assert.Equal(t, 0, s.calculateColumn())
}

func TestShouldReturn7ForRRR(t *testing.T) {
	//explanation
	//[0 - 7] R -> [4 - 7] R -> [6 - 7] R = 7

	//given
	s := seat{
		column: "RRR",
	}
	//when
	//then
	assert.Equal(t, 7, s.calculateColumn())
}

func TestShouldReturn5ForRLR(t *testing.T) {
	//explanation
	//[0 - 7] R -> [4 - 7] L -> [4 - 5] R = 5

	//given
	s := seat{
		column: "RLR",
	}
	//when
	//then
	assert.Equal(t, 5, s.calculateColumn())
}

func TestCalculateSeatID(t *testing.T) {
	type given struct {
		row         string
		column      string
		expectation int
	}
	givens := []given{
		{row: "BFFFBBF", column: "RRR", expectation: 567},
		{row: "FFFBBBF", column: "RRR", expectation: 119},
		{row: "BBFFBBF", column: "RLL", expectation: 820},
	}

	for _, g := range givens {
		got := newSeat(g.row, g.column)
		assert.Equal(t, g.expectation, got.seatID)
	}
}
