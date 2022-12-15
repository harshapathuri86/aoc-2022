package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readByLineStdin()
	marker := 14
	for _, line := range lines {
		l := len(line)
		i := 0
		for ; i <= (l - marker); i++ {
			broken := false
			for j := 0; j <= marker-1; j++ {
				if broken {
					break
				}
				for k := 0; k < j; k++ {
					if broken {
						break
					}
					if line[i+j] == line[i+k] {
						broken = true
					}
				}
			}
			if !broken {
				break
			}
		}
		fmt.Println(i + marker)
	}
}

func readByLineStdin() []string {
	lines := []string{}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
