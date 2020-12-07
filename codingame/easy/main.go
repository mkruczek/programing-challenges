package easy

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func onboarding() {
	for {
		// enemy1: name of enemy 1
		var enemy1 string
		fmt.Scan(&enemy1)

		// dist1: distance to enemy 1
		var dist1 int
		fmt.Scan(&dist1)

		// enemy2: name of enemy 2
		var enemy2 string
		fmt.Scan(&enemy2)

		// dist2: distance to enemy 2
		var dist2 int
		fmt.Scan(&dist2)

		if dist1 < dist2 {
			fmt.Println(enemy1)
		} else {
			fmt.Println(enemy2)
		}

	}
}

func ASCIIART() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var L int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &L)

	var H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &H)

	scanner.Scan()
	T := scanner.Text()
	splited := strings.Split(T, "")

	for i := 0; i < H; i++ {
		scanner.Scan()
		ROW := scanner.Text()
		for _, v := range splited {
			letterPos := getLetterPos(v)
			endTrim := letterPos * L
			startTrim := endTrim - L
			print := ROW[startTrim:endTrim]
			fmt.Print(print)
		}
		fmt.Println("")
	}
}

func getLetterPos(l string) int {
	letters := map[string]int{
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9, "J": 10, "K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16, "Q": 17, "R": 18, "S": 19, "T": 20, "U": 21, "V": 22, "W": 23, "X": 24, "Y": 25, "Z": 26,
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
		"?": 27,
	}
	if result, ok := letters[l]; ok {
		return result
	}
	return letters["?"]
}
