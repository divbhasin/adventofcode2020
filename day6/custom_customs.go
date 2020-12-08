package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func cleanMap(m map[string]int, total int) map[string]int {
	tmpMap := make(map[string]int)
	for k, v := range m {
		if v == total {
			tmpMap[k] = v
		}
	}

	return tmpMap
}

func parseAns(fileName string) []map[string]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	curr_ans := make(map[string]int)

	var all_ans []map[string]int

	num_ppl := 0

	for scanner.Scan() {
		ans := scanner.Text()
		if ans == "" {
			tmp_map := cleanMap(curr_ans, num_ppl)
			all_ans = append(all_ans, tmp_map)
			curr_ans = make(map[string]int)
			num_ppl = 0
		} else {
			options := strings.Split(ans, "")
			for _, o := range options {
				if val, ok := curr_ans[o]; !ok {
					curr_ans[o] = 1
				} else {
					curr_ans[o] = val + 1
				}
			}
			num_ppl += 1
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
