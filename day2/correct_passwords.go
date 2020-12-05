package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func split(line, sep string) (string, string) {
	str := strings.Split(line, sep)
	return str[0], str[1]
}

func verify(line string) bool {
	count_char, password := split(line, ":")
	rang, char := split(count_char, " ")
	lo_str, hi_str := split(rang, "-")
	lo, err1 := strconv.Atoi(lo_str)
	hi, err2 := strconv.Atoi(hi_str)

	if err1 != nil {
		log.Fatal(err1)
	}
	if err2 != nil {
		log.Fatal(err2)
	}

	count := 0
	for _, c := range password {
		if string(c) == char {
			count += 1
		}
	}

	return count >= lo && count <= hi
}

func parseInput(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	valid := 0
	for scanner.Scan() {
		res := verify(scanner.Text())
		if res {
			valid++
		}
	}

	return valid
}

func main() {
	ans := parseInput("input.txt")
	fmt.Println(ans)
}
