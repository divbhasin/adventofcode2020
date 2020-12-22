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
		bag_info := strings.Split(line, " contain ")
		bag := bag_info[0][0 : strings.LastIndex(bag_info[0], "bags")-1]
		if bag_info[1] != "no other bags." {
			neighbour_info := strings.Split(bag_info[1], ",")

			for _, info := range neighbour_info {
				adj_bag := strings.TrimSpace(info[2 : strings.LastIndex(info, "bag")-1])
				if _, ok := adj[bag]; !ok {
					adj[bag] = []string{adj_bag}
				} else {
					adj[bag] = append(adj[bag], adj_bag)
				}
			}
		}
	}

	return adj
}

func visit(visited map[string]bool, adj map[string][]string, col string, count *int) {
	if col == "shiny gold" {
		*count += 1
		return
	}

	for _, adj_col := range adj[col] {
		if !visited[adj_col] {
			visit(visited, adj, adj_col, count)
		}
	}
}

func main() {
	adj := parseBagMappings("input.txt")

	// We have adj, now we can do DFS
	visited := make(map[string]bool)
	for key, _ := range adj {
		visited[key] = false
	}

	reachable := 0
	for k, _ := range adj {
		visited[k] = true
		visit(visited, adj, k, &reachable)
	}
	fmt.Println(reachable)
}
