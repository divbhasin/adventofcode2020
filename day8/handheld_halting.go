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

func execInstructions(ins []string, acc *int) int {
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
				return pc + num
			}
			pc += (num - 1)
		}

		pc += 1
	}

	return pc
}

func main() {
	ins := getInstructions("input.txt")
	acc := 0
	execInstructions(ins, &acc)

	fmt.Printf("Part 1: %d\n", acc)

	// Part 2
	// We can brute-force to check which instruction is preventing halting. In
	// particular, we can iterate through the instructions, change exactly one
	// nop to a jmp or jmp to a nop, and then run execInstructions. If the pc
	// can go past the last instruction, we have found the instruction we needed
	// to halt the program.

	for i, _ := range ins {
		curr_acc := 0
		prev := ins[i]

		if strings.HasPrefix(ins[i], "jmp") {
			ins[i] = "nop" + ins[i][3:]
		} else if strings.HasPrefix(ins[i], "nop") {
			ins[i] = "jmp" + ins[i][3:]
		}

		if !strings.HasPrefix(ins[i], "acc") {
			pc := execInstructions(ins, &curr_acc)
			if pc == len(ins) {
				fmt.Printf("Part 2: %d\n", curr_acc)
			}
		}

		ins[i] = prev
	}
}
