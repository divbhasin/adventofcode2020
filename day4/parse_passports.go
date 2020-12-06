package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func addToPassport(pass map[string]string, line string) {
	fieldVals := strings.Split(line, " ")
	for _, fieldVal := range fieldVals {
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

	scanner := bufio.NewScanner(file)

	currPassport := make(map[string]string)
	var passports []map[string]string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			passports = append(passports, currPassport)
			currPassport = make(map[string]string)
		} else {
			addToPassport(currPassport, line)
		}
	}

	passports = append(passports, currPassport)

	return passports
}

func validateField(key string, val string) bool {
	if key == "byr" {
		return regexp.MatchString("^19[2-9|0][0-9]|200[0-2]$", val)
	} else if key == "iyr" {
		return regexp.MatchString("^201[0-9]|2020$", val)
	} else if key == "eyr" {
		return regexp.MatchString("^202[0-9]|2030$", val)
	}
}

func howManyValid(passports []map[string]string) int {
	reqKeys := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	valid := 0
	for _, pport := range passports {
		currValid := true

		for _, key := range reqKeys {
			if val, ok := pport[key]; !ok {
				currValid = false
				break
			} else {
				currValid = validateField(key, val)
				if currValid == false {
					break
				}
			}
		}

		if currValid {
			valid++
		}
	}

	return valid
}

func main() {
	passports := parsePassports("input.txt")
	numValid := howManyValid(passports)
	fmt.Println(numValid)
}
