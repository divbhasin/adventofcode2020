package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	// part 1
	m := make(map[int]bool)

	for _, elem := range nums {
		if _, ok := m[2020-elem]; ok {
			fmt.Println((2020 - elem) * elem)
		}

		m[elem] = true
	}

	// part 2
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		j := 0
		k := len(nums) - 1

		target := 2020 - nums[i]

		for j < k {
			if nums[j]+nums[k] == target {
				fmt.Println(nums[j] * nums[k] * nums[i])
				j++
				k--
			} else if nums[j]+nums[k] < target {
				j++
			} else {
				k--
			}
		}
	}
}
