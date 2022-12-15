package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	top3 := []int{0, 0, 0}
	sum := 0

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() || sum != 0 {
		line := sc.Text()
		if line != "" {
			v, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalln(err)
			}
			sum += v
			continue
		}
		for i := range top3 {
			if sum <= top3[i] {
				continue
			}
			if i < 2 {
				copy(top3[i+1:], top3[i:])
			}
			top3[i] = sum
			break
		}
		sum = 0
	}
	total := 0
	for _, s := range top3 {
		total += s
	}
	fmt.Println(total, top3)
}
