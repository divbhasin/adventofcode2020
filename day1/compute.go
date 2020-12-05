package main

import "fmt"

func readFile(name string) {

}

func main() {
	nums = readFile("input.txt")
	m = make(map[int]bool)

	for i, elem := range nums {
		if val, ok := m[2020-elem]; ok {
			fmt.Println((2020 - elem) * elem)
		}
	}
	fmt.Println("vim-go")
}
