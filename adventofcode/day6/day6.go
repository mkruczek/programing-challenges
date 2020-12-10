package day6

import (
	"bufio"
	"github.com/mkruczek/programing-challenges/adventofcode/service"
	"strings"
)

const (
	day6       = "../programing-challenges/adventofcode/day6/input"
	day6Sample = "../programing-challenges/adventofcode/day6/sampleInput"
)

type group struct {
	answers []string
}

func (g *group) calculateYes() int {
	counterSet := make(map[string]bool)
	for _, a := range g.answers {
		for _, v := range strings.Split(a, "") {
			counterSet[v] = true
		}
	}
	return len(counterSet)
}

func CalculateAnswers() int {
	counter := 0
	for _, a := range loadInputDay6() {
		counter += a.calculateYes()
	}
	return counter
}

func loadInputDay6() []group {
	file, err := service.OpenFile(day6)
	if err != nil {
		panic("couldn't open file with data : " + err.Error())
	}
	defer func() {
		_ = file.Close()
	}()

	sc := bufio.NewScanner(file)
	var groupsString []string
	currentReadingGroup := ""
	for sc.Scan() {
		line := sc.Text()
		if len(strings.TrimSpace(line)) == 0 {
			groupsString = append(groupsString, currentReadingGroup)
			currentReadingGroup = ""
		} else {
			currentReadingGroup += line + " "
		}
	}
	groupsString = append(groupsString, currentReadingGroup) //added last doc because file in not ending with empty line

	result := []group{}
	for _, v := range groupsString {
		gr := convertLineToGroup(v)
		result = append(result, *gr)
	}

	if err = sc.Err(); err != nil {
		panic("couldn't open file with data : " + err.Error())
	}
	return result
}

func convertLineToGroup(l string) *group {
	result := group{}
	for _, v := range strings.Split(l, " ") {
		result.answers = append(result.answers, v)
	}
	return &result
}
