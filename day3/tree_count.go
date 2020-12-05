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

func computeTrees(terrain [][]string) int {
	m, n := len(terrain), len(terrain[0])
	j := 0
	trees := 0

	for i := 0; i < m; i++ {
		item := terrain[i][j%n]
		if item == "#" {
			trees += 1
		}

		j += 3
	}

	return trees
}

func main() {
	terrain := parseMap("input.txt")
	trees := computeTrees(terrain)
	fmt.Println(trees)
}
