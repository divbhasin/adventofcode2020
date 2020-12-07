package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parsePasses(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var passes []string

	for scanner.Scan() {
		passes = append(passes, scanner.Text())
	}

	return passes
}

func binVisit(lo int, hi int, lo_idx int, hi_idx int, pass string) int {
	res := 0

	for i := lo_idx; i < hi_idx; i++ {
		mid := (lo + hi) / 2
		if lo >= hi {
			res = mid
		} else {
			if pass[i] == 'F' {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}

	return res
}

func calcRowCol(pass string) (int, int) {
	// parse row
	r := binVisit(0, 127, 0, 7, pass)
	c := binVisit(0, 7, 7, 10, pass)

	return r, c
}

func calcIds(passes []string) []int {
	var ids []int

	for _, pass := range passes {
		r, c := calcRowCol(pass)
		fmt.Println(pass, r, c)
		id := r*8 + c
		ids = append(ids, id)
	}

	return ids
}

func main() {
	boardingPasses := parsePasses("input.txt")
	ids := calcIds(boardingPasses)
	max_id := 0

	for _, id := range ids {
		if id > max_id {
			max_id = id
		}
	}

	fmt.Println(max_id)
}
