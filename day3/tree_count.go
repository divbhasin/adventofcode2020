package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseMap(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var board [][]string
	for scanner.Scan() {
		row_str := scanner.Text()
		row := strings.Split(row_str, "")
		board = append(board, row)
	}

	return board
}

func computeTrees(terrain [][]string, slope [2]int) int {
	m, n := len(terrain), len(terrain[0])
	j := 0
	trees := 0

	for i := 0; i < m; i += slope[0] {
		item := terrain[i][j%n]
		if item == "#" {
			trees += 1
		}

		j += slope[1]
	}

	return trees
}

func main() {
	terrain := parseMap("input.txt")

	// Part 1
	slope := [2]int{1, 3}
	trees := computeTrees(terrain, slope)
	fmt.Println(trees)

	// Part 2
	slopes := [5][2]int{[2]int{1, 1}, [2]int{1, 3}, [2]int{1, 5}, [2]int{1, 7}, [2]int{2, 1}}
	prod := 1

	for _, slop := range slopes {
		prod *= computeTrees(terrain, slop)
	}

	fmt.Println(prod)
}
