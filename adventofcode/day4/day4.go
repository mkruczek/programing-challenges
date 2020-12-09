package day4

import (
	"bufio"
	"fmt"
	"github.com/mkruczek/programing-challenges/adventofcode/service"
	"strconv"
	"strings"
)

const (
	day4       = "../programing-challenges/adventofcode/day4/input"
	day4Sample = "../programing-challenges/adventofcode/day4/sampleInput"
)

type document struct {
	byr *int    //Birth Year
	iyr *int    //Issue Year
	eyr *int    //Expiration Year
	hgt *string //Height
	hcl *string //Hair Color
	ecl *string //Eye Color
	pid *int    //Passport ID
	cid *int    //Country ID - not require
}

func (d *document) isValid() bool {
	return d.byr != nil && d.iyr != nil && d.eyr != nil && d.hgt != nil && d.hcl != nil && d.ecl != nil && d.pid != nil
}

func HelpWithDocuments() int {
	validCounter := 0
	for _, d := range loadInputDay4() {
		if d.isValid() {
			validCounter++
		}
	}
	return validCounter
}

func loadInputDay4() []document {
	file, err := service.OpenFile(day4)
	if err != nil {
		panic("couldn't open file with data : " + err.Error())
	}
	defer func() {
		_ = file.Close()
	}()

	sc := bufio.NewScanner(file)
	var documents []string
	currentReadingDoc := ""
	for sc.Scan() {
		line := sc.Text()
		if len(strings.TrimSpace(line)) == 0 {
			documents = append(documents, currentReadingDoc)
			currentReadingDoc = ""
		} else {
			currentReadingDoc += line + " "
		}
	}
	documents = append(documents, currentReadingDoc) //added last doc because file in not ending with empty line

	var result []document
	for _, l := range documents {
		doc, err := convertLineToDocument(l)
		if err != nil {
			fmt.Printf("error during convert to document [%s] : %v", l, err)
			continue
		}
		result = append(result, *doc)
	}

	if err = sc.Err(); err != nil {
		panic("couldn't open file with data : " + err.Error())
	}
	return result
}

func convertLineToDocument(l string) (*document, error) {
	doc := new(document)
	splited := strings.Split(l, " ")
	for _, s := range splited {
		singleValue := strings.Split(s, ":")
		kind := singleValue[0]
		switch kind {
		case "byr":
			i, err := strconv.Atoi(singleValue[1])
			if err != nil {
				return nil, err
			}
			doc.byr = &i
		case "iyr":
			i, err := strconv.Atoi(singleValue[1])
			if err != nil {
				return nil, err
			}
			doc.iyr = &i
		case "eyr":
			i, err := strconv.Atoi(singleValue[1])
			if err != nil {
				return nil, err
			}
			doc.eyr = &i
		case "hgt":
			doc.hgt = &singleValue[1]
		case "hcl":
			doc.hcl = &singleValue[1]
		case "ecl":
			doc.ecl = &singleValue[1]
		case "pid":
			i, err := strconv.Atoi(singleValue[1])
			if err != nil {
				return nil, err
			}
			doc.pid = &i
		case "cid":
			i, err := strconv.Atoi(singleValue[1])
			if err != nil {
				return nil, err
			}
			doc.cid = &i
		}
	}
	return doc, nil
}
