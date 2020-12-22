package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInstructions(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var ins []string

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		ins = append(ins, line)
	}

	return ins
}

func execInstructions(ins []string, acc *int) {

}

func main() {
	ins := getInstructions("input.txt")
	acc := 0
	execInstructions(ins, acc)

	fmt.Println(acc)
}
