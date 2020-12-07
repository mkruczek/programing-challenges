package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

}

func findSeed(k int32, f []int32, c []int32) []int32 {
	/*
	 * Write your code here.
	 */

	return []int32{1, 2, 3}
}

func timeConversion(s string) string {

	if strings.HasPrefix(s, "12") && strings.HasSuffix(s, "PM") {
		return s[:len(s)-2]
	}

	if strings.HasPrefix(s, "12") && strings.HasSuffix(s, "AM") {
		return "00" + s[2:len(s)-2]
	}

	split := strings.Split(s, ":")
	v, _ := strconv.Atoi(split[0])
	if strings.HasSuffix(s, "PM") {
		v += 12
		split[0] = strconv.Itoa(v)
	} else if strings.HasPrefix(split[0], "0") {
		split[0] = "0" + strconv.Itoa(v)
	} else {
		split[0] = strconv.Itoa(v)
	}

	for i := 1; i < len(split); i++ {
		split[0] += ":" + split[i]
	}
	return split[0][:len(split[0])-2]
}

func birthdayCakeCandles(candles []int32) int32 {
	var max int32
	var m = make(map[int32]int32)
	for _, num := range candles {
		if num >= max {
			max = num
			m[max] = m[max] + 1
		}
	}
	return m[max]
}

func miniMaxSum(arr []int32) {
	var min, max int64

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	for i := 0; i < len(arr)-1; i++ {
		min += int64(arr[i])
		max += int64(arr[len(arr)-1-i])
	}
	fmt.Printf("%v %v", min, max)
}

func simpleArraySum(ar []int32) int32 {
	r := int32(0)
	for _, num := range ar {
		r += int32(num)
	}
	return r
}

func staircase(n int32) {
	p := int(n)
	for i := 0; i < int(n); i++ {
		p--
		for j := 0; j < p; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < (i + 1); j++ {
			fmt.Printf("%s", "#")
		}
		fmt.Printf("\n")
	}
}

func plusMinus(arr []int32) {
	var p, m, z float64
	l := float64(len(arr))
	for _, num := range arr {
		switch {
		case num == 0:
			z++
		case num > 0:
			p++
		case num < 0:
			m++
		}
	}
	fmt.Printf("%.6f\n%.6f\n%.6f\n", p/l, m/l, z/l)
}

func diagonalDifference(arr [][]int32) int32 {
	var fd, sd int32
	for i := 0; i < len(arr); i++ {
		fd += arr[i][i]
		sd += arr[len(arr)-i-1][i]
	}
	if result := fd - sd; result < 0 {
		return -result
	} else {
		return result
	}
}

func aVeryBigSum(ar []int64) int64 {
	var result int64
	for _, num := range ar {
		result += num
	}
	return result
}

func compareTriplets(a []int32, b []int32) []int32 {
	result := []int32{0, 0}
	for i, _ := range a {
		if a[i] > b[i] {
			result[0]++
		} else if a[i] < b[i] {
			result[1]++
		}
	}
	return result
}
