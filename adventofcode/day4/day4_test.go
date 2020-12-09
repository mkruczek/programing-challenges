package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnTrueForValidDocument(t *testing.T) {
	//given
	randomMockNumber := 2000
	randomMockString := "abc"
	doc := document{
		byr: &randomMockNumber,
		iyr: &randomMockNumber,
		eyr: &randomMockNumber,
		hgt: &randomMockString,
		hcl: &randomMockString,
		ecl: &randomMockString,
		pid: &randomMockNumber,
		cid: &randomMockNumber,
	}
	//when
	got := doc.isValid()
	//then
	assert.True(t, got)
}

func TestShouldReturnTrueForValidDocumentWithOutCID(t *testing.T) {
	//given
	randomMockNumber := 2000
	randomMockString := "abc"
	doc := document{
		byr: &randomMockNumber,
		iyr: &randomMockNumber,
		eyr: &randomMockNumber,
		hgt: &randomMockString,
		hcl: &randomMockString,
		ecl: &randomMockString,
		pid: &randomMockNumber,
	}
	//when
	got := doc.isValid()
	//then
	assert.True(t, got)
}

func TestShouldReturnFalseForDocumentWithMissingOneField(t *testing.T) {
	//given
	randomMockNumber := 2000
	randomMockString := "abc"
	documents := []document{
		{iyr: &randomMockNumber, eyr: &randomMockNumber, hgt: &randomMockString, hcl: &randomMockString, ecl: &randomMockString, pid: &randomMockNumber},
		{byr: &randomMockNumber, eyr: &randomMockNumber, hgt: &randomMockString, hcl: &randomMockString, ecl: &randomMockString, pid: &randomMockNumber},
		{byr: &randomMockNumber, iyr: &randomMockNumber, hgt: &randomMockString, hcl: &randomMockString, ecl: &randomMockString, pid: &randomMockNumber},
		{byr: &randomMockNumber, iyr: &randomMockNumber, eyr: &randomMockNumber, hcl: &randomMockString, ecl: &randomMockString, pid: &randomMockNumber},
		{byr: &randomMockNumber, iyr: &randomMockNumber, eyr: &randomMockNumber, hgt: &randomMockString, ecl: &randomMockString, pid: &randomMockNumber},
		{byr: &randomMockNumber, iyr: &randomMockNumber, eyr: &randomMockNumber, hgt: &randomMockString, hcl: &randomMockString, pid: &randomMockNumber},
		{byr: &randomMockNumber, iyr: &randomMockNumber, eyr: &randomMockNumber, hgt: &randomMockString, hcl: &randomMockString, ecl: &randomMockString},
	}

	for _, d := range documents {
		//when
		got := d.isValid()
		//then
		assert.False(t, got)
	}
}
