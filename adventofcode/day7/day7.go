package day7

import (
	"bufio"
	"fmt"
	"github.com/mkruczek/programing-challenges/adventofcode/service"
	"regexp"
	"strings"
)

const (
	day7       = "../programing-challenges/adventofcode/day7/input"
	day7Sample = "../programing-challenges/adventofcode/day7/sampleInput"
)

type bag struct {
	color string
	bags map[int]bag
}


func loadInputDay7() []bag {
	var result []bag
	file, err := service.OpenFile(day7)
	if err != nil {
		panic("couldn't open file with data : " + err.Error())
	}
	defer func() {
		_ = file.Close()
	}()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		result = append(result, convertLineToBag(sc.Text()))
	}
	if err = sc.Err(); err != nil {
		panic("couldn't open file with data : " + err.Error())
	}
	return result
}

func convertLineToBag(line string) bag {


	splited := strings.Split(line, " bags contain ")
	mainColour := splited[0]
	result := bag{color: mainColour}

	splited2 :=strings.Split(splited[1], ", ")

	if !strings.EqualFold(splited2[0], "no other bags."){
		result.bags = make(map[int]bag)
		for _, v := range splited2{
			amount := v[:1] //todo what if more then 9
			colour := v[1:]
			colour = strings.Split(colour, " bag")[0]
			result.bags[]
			fmt.Println(amount," :  ",colour)
		}
	}


	return bag{color: mainColour}
}

func goReSplit(text string, pattern string) []string {
	regex := regexp.MustCompile(pattern)
	result := regex.Split(text, -1)
	return result
}