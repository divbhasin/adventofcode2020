package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(name string) []int {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var nums []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	return nums
}

func main() {
	nums := readFile("input.txt")
	m := make(map[int]bool)

	for _, elem := range nums {
		if _, ok := m[2020-elem]; ok {
			fmt.Println((2020 - elem) * elem)
		}

		m[elem] = true
	}
}
