package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileCounts := make(map[string]map[string]int)
	files := os.Args[1:]
	fmt.Println(files)
	if len(files) == 0 {
		countLines(os.Stdin, fileCounts)
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			fileCounts[f.Name()] = counts
			countLines(f, fileCounts)
			f.Close()
		}
	}
	for fileName, item := range fileCounts {
		for line, n := range item {

			if n > 1 {
				fmt.Printf("%s\t%s\t%d\n", fileName, line, n)
			}
		}

	}
}

func countLines(f *os.File, fileCounts map[string]map[string]int) {
	//fmt.Println(f.Name())
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, ok := fileCounts[f.Name()][input.Text()]; ok {
			fileCounts[f.Name()][input.Text()]++
		} else {
			fileCounts[f.Name()][input.Text()] = 1
		}
	}
}
