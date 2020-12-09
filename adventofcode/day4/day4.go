package day4

import (
	"bufio"
	"fmt"
	"github.com/mkruczek/programing-challenges/adventofcode/service"
	"regexp"
	"strconv"
	"strings"
)

const (
	day4             = "../programing-challenges/adventofcode/day4/input"
	day4Sample       = "../programing-challenges/adventofcode/day4/sampleInput"
	day4SamplePartII = "../programing-challenges/adventofcode/day4/sampleInputPartII"
)

//figure out how do it with regex

var allowedEcl = [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

type document struct {
	byr *int    //Birth Year
	iyr *int    //Issue Year
	eyr *int    //Expiration Year
	hgt *string //Height
	hcl *string //Hair Color
	ecl *string //Eye Color
	pid *string //Passport ID
	cid *int    //Country ID - not require
}

func (d *document) isValid() bool {
	return d.byr != nil && d.iyr != nil && d.eyr != nil && d.hgt != nil && d.hcl != nil && d.ecl != nil && d.pid != nil
}

func (d *document) isValidPartII() bool {
	isByrValid := d.byr != nil && 1920 <= *d.byr && *d.byr <= 2002
	isIyrValid := d.iyr != nil && 2010 <= *d.iyr && *d.iyr <= 2020
	isEyrValid := d.eyr != nil && 2020 <= *d.eyr && *d.eyr <= 2030
	isHgtValid := validateHgt(d.hgt)
	isEclValid := validateEcl(d.ecl)
	isPidValid := validatePid(d.pid)
	isHclValid := validateHcl(d.hcl)

	return isByrValid && isIyrValid && isEyrValid && isHgtValid && isEclValid && isPidValid && isHclValid
}

func HelpWithDocuments() int {
	validCounter := 0
	for _, d := range loadInputDay4() {
		if d.isValidPartII() {
			validCounter++
		}
	}
	return validCounter
}

func validateHcl(hcl *string) bool {
	if hcl == nil {
		return false
	}
	isHclValid, err := regexp.MatchString("#[0-9a-f]{6}", *hcl)
	if err != nil {
		fmt.Printf("error during validation document.Hcl for: %v\n", *hcl)
	}
	return isHclValid
}

func validatePid(pid *string) bool {
	if pid == nil {
		return false
	}
	isPidValid, err := regexp.MatchString("^[0-9]{9}$", *pid)
	if err != nil {
		fmt.Printf("error during validation document.Pid for: %v\n", *pid)
		return false
	}
	return isPidValid
}

func validateHgt(hgt *string) bool {
	if hgt == nil {
		return false
	}
	helper := *hgt
	unit := helper[len(helper)-2:]
	switch unit {
	case "cm":
		i, err := strconv.Atoi(helper[:len(helper)-2])
		if err != nil {
			return false
		}
		return 150 <= i && i <= 193
	case "in":
		i, err := strconv.Atoi(helper[:len(helper)-2])
		if err != nil {
			return false
		}
		return 59 <= i && i <= 76
	default:
		return false
	}
}

func validateEcl(ecl *string) bool {
	if ecl == nil {
		return false
	}
	for _, a := range allowedEcl {
		if strings.EqualFold(a, *ecl) {
			return true
		}
	}
	return false
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
			doc.pid = &singleValue[1]
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
