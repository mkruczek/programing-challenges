package day5

import (
	"bufio"
	"fmt"
	"github.com/mkruczek/programing-challenges/adventofcode/service"
	"sort"
	"strings"
)

const (
	baseForSeatId = 8
	day5          = "../programing-challenges/adventofcode/day5/input"
)

type seat struct {
	row    string
	column string
	seatID int
}

func newSeat(row, column string) *seat {
	result := &seat{
		row:    row,
		column: column,
	}
	result.calculateSeatId()
	return result
}

func (s *seat) calculateSeatId() {
	s.seatID = s.calculateRow()*baseForSeatId + s.calculateColumn()
}

func (s *seat) calculateRow() int {
	splited := strings.Split(s.row, "")
	rows := getStarterRows()
	for _, r := range splited {
		switch r {
		case "F":
			rows = rows[:len(rows)/2]
		case "B":
			rows = rows[len(rows)/2:]
		default:
			fmt.Printf("wrong row for : %v\n", s)
		}
	}
	return rows[0]
}

func (s *seat) calculateColumn() int {
	splited := strings.Split(s.column, "")
	rows := getStarterColumns()
	for _, r := range splited {
		switch r {
		case "L":
			rows = rows[:len(rows)/2]
		case "R":
			rows = rows[len(rows)/2:]
		default:
			fmt.Printf("wrong column for : %v\n", s)
		}
	}
	return rows[0]
}

func getStarterRows() []int {
	var result []int
	for i := 0; i < 128; i++ {
		result = append(result, i)
	}
	return result
}

func getStarterColumns() []int {
	var result []int
	for i := 0; i < 8; i++ {
		result = append(result, i)
	}
	return result
}

func ReturnHighestSeatId() int {
	h := 0
	for _, s := range loadInputDay5() {
		if s.seatID > h {
			h = s.seatID
		}
	}
	return h
}

func FindMySeat() int {
	seats := loadInputDay5()
	sort.Slice(seats, func(i, j int) bool {
		return seats[i].seatID < seats[j].seatID
	})

	for i := 1; i < len(seats); i++ {
		if seats[i].seatID-seats[i-1].seatID != 1 {
			return seats[i].seatID - 1
		}
	}
	return 0
}

func loadInputDay5() []*seat {
	var result []*seat
	file, err := service.OpenFile(day5)
	if err != nil {
		panic("couldn't open file with data : " + err.Error())
	}
	defer func() {
		_ = file.Close()
	}()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		row := sc.Text()[:7]
		column := sc.Text()[7:]
		result = append(result, newSeat(row, column))
	}
	if err = sc.Err(); err != nil {
		panic("couldn't open file with data : " + err.Error())
	}
	return result
}
