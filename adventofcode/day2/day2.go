package day2

import (
	"bufio"
	"github.com/mkruczek/programing-challenges/adventofcode/service"
	"strconv"
	"strings"
)

const (
	day2       = "../programing-challenges/adventofcode/day2/input"
	day2Sample = "../programing-challenges/adventofcode/day2/sampleInput"
)

func TestPasswordsFirstRule() (validPassCounter int) {
	for _, pp := range loadInputDay2(day2) {
		if pp.isValidFirstRule() {
			validPassCounter++
		}
	}
	return validPassCounter
}

func TestPasswordsSecondRule() (validPassCounter int) {
	for _, pp := range loadInputDay2(day2) {
		if pp.isValidSecondRule() {
			validPassCounter++
		}
	}
	return validPassCounter
}

type passPolicy struct {
	letter string
	min    int
	max    int
	pass   string
}

func (pp passPolicy) isValidFirstRule() bool {
	counter := 0
	for _, s := range pp.pass {
		if strings.EqualFold(pp.letter, string(s)) {
			counter++
		}
	}
	return pp.min <= counter && counter <= pp.max
}

func (pp passPolicy) isValidSecondRule() bool {
	return strings.EqualFold(pp.letter, string(pp.pass[pp.min-1])) && !strings.EqualFold(pp.letter, string(pp.pass[pp.max-1])) ||
		!strings.EqualFold(pp.letter, string(pp.pass[pp.min-1])) && strings.EqualFold(pp.letter, string(pp.pass[pp.max-1]))
}

func newPassPolicyFromString(s string) passPolicy {
	//string pattern : min-max letter: password
	splited := strings.Split(s, " ")
	minValue, _ := strconv.Atoi(strings.Split(splited[0], "-")[0])
	maxValue, _ := strconv.Atoi(strings.Split(splited[0], "-")[1])
	return passPolicy{
		letter: splited[1][:1],
		min:    minValue,
		max:    maxValue,
		pass:   splited[2],
	}
}

func loadInputDay2(path string) []passPolicy {
	file, err := service.OpenFile(path)
	defer file.Close()
	if err != nil {
		panic("couldn't open file with data : " + err.Error())
	}
	result := []passPolicy{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		result = append(result, newPassPolicyFromString(sc.Text()))
	}
	if err = sc.Err(); err != nil {
		panic("couldn't open file with data : " + err.Error())
	}
	return result
}
