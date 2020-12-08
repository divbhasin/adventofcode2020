package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseAns(fileName string) []map[string]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	curr_ans := make(map[string]int)

	var all_ans []map[string]int

	for scanner.Scan() {
		ans := scanner.Text()
		if ans == "" {
			all_ans = append(all_ans, curr_ans)
			curr_ans = make(map[string]int)
		} else {
			options := strings.Split(ans, "")
			for _, o := range options {
				if _, ok := curr_ans[o]; !ok {
					curr_ans[o] = 0
				}
			}
		}
	}

	return all_ans
}

func main() {
	group_ans := parseAns("input.txt")
	total := 0

	for _, m := range group_ans {
		total += len(m)
	}

	fmt.Println(total)
}
