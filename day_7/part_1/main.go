package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day_7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]string

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, len(line))
		for i, ch := range line {
			row[i] = string(ch)
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print grid to verify
	maxRows := len(grid)

	splitterCount := 0
	for y, row := range grid {
		if y == maxRows - 1 {
			continue
		}
		for x, cell := range row {
			if cell == "." { continue }
			if cell == "S" || cell == "|" {
				if grid[y+1][x] != "^" {
					grid[y+1][x] = "|"
				}
			}
			if grid[y][x] == "|" && grid[y+1][x] == "^" {
				if grid[y+1][x+1] != "^" {
					grid[y+1][x+1] = "|"
				}
				if grid[y+1][x-1] != "^" {
					grid[y+1][x-1] = "|"
				}
				splitterCount++
			}
		}
		fmt.Printf("%v\n", row)
	}
	fmt.Printf("Splitter count: %d", splitterCount)
}
