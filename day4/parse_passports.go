package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func addToPassport(pass *map[string]string, line string) {
	fieldVals := strings.Split(line, " ")
	for fieldVal := range fieldVals {
		temp := strings.Split(fieldVal, ":")
		field, val := temp[0], temp[1]

		pass[field] = val
	}
}

func parsePassports(fileName string) []map[string]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := scanner.NewScanner(file)

	var currPassport map[string]string
	var passports []map[string]string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, currPassport)
		}

		currPassport = addToPassport(&currPassport, line)
	}

	return passports
}

func main() {
	passports := parsePassports("input.txt")
	numValid := howManyValid(passports)
	fmt.Println(numValid)
}
