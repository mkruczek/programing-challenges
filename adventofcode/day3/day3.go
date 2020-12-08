package day3

import (
	"bufio"
	"github.com/mkruczek/programing-challenges/adventofcode/service"
)

const (
	day3       = "../programing-challenges/adventofcode/day3/input"
	day3Sample = "../programing-challenges/adventofcode/day3/sampleInput"
)

func EncounterTreesForDifferentPattern() int {
	type xyPattern struct {
		x, y int
	}
	patterns := []xyPattern{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	result := 1
	for _, p := range patterns {
		result *= EncounterTrees(p.x, p.y)
	}
	return result
}

func EncounterTrees(x, y int) int {
	trees := 0
	data := loadInputDay3()
	counterX := 0
	for i := 0; i < len(data); i += y {
		if counterX >= len(data[i]) {
			counterX -= len(data[i])
		}
		if string(data[i][counterX]) == "#" {
			trees++
		}
		counterX += x
	}
	return trees
}

func loadInputDay3() []string {
	var result []string
	file, err := service.OpenFile(day3)
	if err != nil {
		panic("couldn't open file with data : " + err.Error())
	}
	defer func() {
		_ = file.Close()
	}()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		result = append(result, sc.Text())
	}
	if err = sc.Err(); err != nil {
		panic("couldn't open file with data : " + err.Error())
	}
	return result
}
