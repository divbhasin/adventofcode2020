package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInstructions(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var ins []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ins = append(ins, line)
	}

	return ins
}

func find(i int, arr []int) bool {
	for _, val := range arr {
		if val == i {
			return true
		}
	}

	return false
}

func execInstructions(ins []string, acc *int) {
	pc := 0
	var execd []int

	for pc < len(ins) {
		curr_ins := ins[pc]
		num, err := strconv.Atoi(curr_ins[4:])

		if err != nil {
			log.Fatal(err)
		}

		execd = append(execd, pc)

		if strings.HasPrefix(curr_ins, "acc") {
			*acc += num
		} else if strings.HasPrefix(curr_ins, "jmp") {
			if find(pc+num, execd) {
				return
			}
			pc += (num - 1)
		}

		pc += 1
	}
}

func main() {
	ins := getInstructions("input.txt")
	acc := 0
	execInstructions(ins, &acc)

	fmt.Println(acc)
}
