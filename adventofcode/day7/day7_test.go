package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertLineToBagWhenItContainsOtherBags(t *testing.T) {
	//given
	testCases := []struct {
		line     string
		expected bag
	}{
		{
			line:     "light red bags contain 1 bright white bag, 2 muted yellow bags.",
			expected: bag{color: "light red bags", bags: map[int]bag{1: {color: "bright white"}, 2: {color: "muted yellow"}}},
		},
		{
			line:     "bright white bags contain 1 shiny gold bag.",
			expected: bag{color: "bright white", bags: map[int]bag{1: {color: "shiny gold"}}},
		},
		{
			line:     "faded blue bags contain no other bags.",
			expected: bag{color: "faded blue"},
		},
	}

	for _, tc := range testCases {
		//when
		got := convertLineToBag(tc.line)

		//then
		assert.Equal(t, tc.expected, got)
	}
}
