package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseBagMappings(fileName string) map[string][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	adj := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		bag_info := strings.Split(line, "contain")
		bag := bag_info[0]
		neighbour_info := strings.Split(bag_info[1], ",")
		for _, info := range neighbour_info {
			adj_bag := info[3 : strings.LastIndex(info, "bag")-1]
			fmt.Println(adj_bag)
			if _, ok := adj[bag]; !ok {
				adj[bag] = []string{adj_bag}
			} else {
				adj[bag] = append(adj[bag], adj_bag)
			}
		}
	}

	return adj
}

func visited(visited, adj, col) {
	visited[col] = true
	for _, adj_col := range adj[col] {

	}
}

func main() {
	adj := parseBagMappings("input.txt")

	// We have adj, now we can do DFS
	visited := []bool{}
	for key, _ := range adj {
		visited[key] = false
	}

	visit(visited, adj, "shiny gold")
	reachable := 0

	for _, val := range visited {
		if val == true {
			reachable += 1
		}
	}

	fmt.Println(reachable)
}
