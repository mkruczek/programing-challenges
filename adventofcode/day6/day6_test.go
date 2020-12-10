package day6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inProgress(t *testing.T) {
	//given
	g := group{
		answers: []string{"abc", "abd", "abe", "x"},
	}

	//when
	//then
	assert.Equal(t, 6, g.calculateYes())
}
