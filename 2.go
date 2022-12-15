package main

import (
	"bufio"
	"fmt"
	"strings"

	// "log"
	"os"
	// "strconv"
)

func main() {
	win := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}
	draw := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
	lose := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}
	cost := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	score := 0
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		words := strings.Split(line, " ")
		play := ""
		if words[1] == "X" {
			play = lose[words[0]]
			score += 0
		} else if words[1] == "Y" {
			play = draw[words[0]]
			score += 3
		} else {
			play = win[words[0]]
			score += 6
		}
		score += cost[play]
	}
	fmt.Println(score)

}
