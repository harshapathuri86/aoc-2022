package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	ans := 0
	cnt := 0
	priority := make(map[rune]int)
	for i := 'a'; i <= 'z'; i++ {
		priority[i] = int(i-'a') + 1
	}
	for i := 'A'; i <= 'Z'; i++ {
		priority[i] = int(i-'A') + 27
	}
	m1 := make(map[rune]bool)
	m2 := make(map[rune]bool)
	for sc.Scan() {
		cnt += 1
		line := sc.Text()
		if cnt%3 == 1 {
			for _, v := range line {
				m1[v] = true
			}
		} else if cnt%3 == 2 {
			for _, v := range line {
				m2[v] = true
			}
		} else {
			for _, v := range line {
				if m1[v] && m2[v] {
					ans += priority[v]
					break
				}
			}
			m1 = make(map[rune]bool)
			m2 = make(map[rune]bool)
		}
	}
	fmt.Println(ans)
}
