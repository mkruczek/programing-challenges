package day7

import (
	"bufio"
	"github.com/mkruczek/programing-challenges/adventofcode/service"
	"strconv"
	"strings"
)

const (
	myBag      = "shiny gold"
	day7       = "../programing-challenges/adventofcode/day7/input"
	day7Sample = "../programing-challenges/adventofcode/day7/sampleInput"
)

type bag struct {
	color  string
	amount int
	bags   []bag
}

func InProgress() interface{} {
	data := loadInputDay7()
	directly := bagCanContainsMyBagDirectly(myBag, data)
	return
}

func bagCanContainsMyBagDirectly(myBag string, data []bag) []bag {
	result := make([]bag, 0)
	for _, v := range data {
		for _, j := range v.bags {
			if j.color == myBag {
				result = append(result, v)
			}
		}
	}
	return result
}

func loadInputDay7() []bag {
	var result []bag
	file, err := service.OpenFile(day7Sample)
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
	splited2 := strings.Split(splited[1], ", ")
	if !strings.EqualFold(splited2[0], "no other bags.") {
		result.bags = []bag{}
		for _, v := range splited2 {
			tmpToSplit := strings.Index(v, " ")
			sAmount := v[:tmpToSplit]
			amount, _ := strconv.Atoi(sAmount)
			colour := v[(tmpToSplit + 1):]
			colour = strings.Split(colour, " bag")[0]
			result.bags = append(result.bags, bag{
				color:  colour,
				amount: amount,
				bags:   nil,
			})
		}
	}
	return result
}
