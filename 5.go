package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readByLineStdin()
	processStack := true
	var stacks [][]byte
	re := regexp.MustCompile("[0-9]+")
	for _, line := range lines {
		if line == "" {
			continue
		}
		if processStack == true && strings.ContainsAny(line, "123456789") {
			// fmt.Println("contains ", line)
			// re := regexp.MustCompile("[0-9]+")
			// x := re.FindAllString(line, -1)
			// n = len(x)
			// stacks = make([][]byte, n)
			processStack = false
			continue
		}
		if processStack {
			for i := 1; i < len(line); i += 4 {

				if line[i] != ' ' {
					pos := (i - 1) / 4
					for len(stacks) <= pos {
						stacks = append(stacks, []byte{})
					}
					stacks[pos] = append([]byte{line[i]}, stacks[pos]...)
				}
			}
			// fmt.Println(stacks)
		} else {
			// fmt.Println("todo ", line)
			nums_str := re.FindAllString(line, -1)
			nums := convertInt(nums_str)
			from := nums[1] - 1
			to := nums[2] - 1
			cnt := nums[0]
			if cnt > 0 {
				// for cnt > 0 {
				// // fmt.Println("from", from, "to", to, "cnt", cnt)
				l := len(stacks[from])
				// stacks[to] = append(stacks[to], stacks[from][l-1])
				// stacks[from] = stacks[from][:(l - 1)]
				// cnt--
				stacks[to] = append(stacks[to], stacks[from][l-cnt:]...)
				stacks[from] = stacks[from][:(l - cnt)]
			}
			// fmt.Println("stacks", stacks)
		}
	}
	ans := ""
	for _, stack := range stacks {
		ans += string(stack[len(stack)-1])
	}
	fmt.Println(ans)
}

func convertInt(nums_str []string) []int {
	var nums []int
	for _, num := range nums_str {
		i, _ := strconv.Atoi(num)
		nums = append(nums, i)
	}
	return nums
}

func readByLine(filename string) []string {
	lines := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
		return lines
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func readByLineStdin() []string {
	lines := []string{}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
