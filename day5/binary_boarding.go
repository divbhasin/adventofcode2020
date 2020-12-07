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

func binVisit(lo int, hi int, lo_idx int, hi_idx int, pass string, f byte) int {
	for i := lo_idx; i < hi_idx; i++ {
		mid := (lo + hi) / 2
		if pass[i] == f {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return (lo + hi) / 2
}

func calcRowCol(pass string) (int, int) {
	// parse row
	r := binVisit(0, 127, 0, 7, pass, 'F')
	c := binVisit(0, 7, 7, 10, pass, 'L')

	return r, c
}

func calcIds(passes []string) []int {
	var ids []int

	for _, pass := range passes {
		r, c := calcRowCol(pass)
		id := r*8 + c
		ids = append(ids, id)
	}

	return ids
}

func main() {
	boardingPasses := parsePasses("input.txt")

	// part 1
	ids := calcIds(boardingPasses)
	max_id := 0

	for _, id := range ids {
		if id > max_id {
			max_id = id
		}
	}

	fmt.Println(max_id)

	// part 2
	id_map := make(map[int]int)
	for _, id := range ids {
		id_map[id] = 0
	}

	for _, id := range ids {
		if _, ok := id_map[id+2]; ok {
			if _, ok := id_map[id+1]; !ok {
				fmt.Println(id + 1)
			}
		} else if _, ok := id_map[id-2]; ok {
			if _, ok := id_map[id-1]; !ok {
				fmt.Println(id - 1)
			}
		}
	}
}
