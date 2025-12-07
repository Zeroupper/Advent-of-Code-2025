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

	timelineCount := 0
	startingX := 0
	startingY := 0
	for y, row := range grid {
		if y == maxRows-1 {
			fmt.Printf("Row %d: %v\n", y, row)
			continue
		}
		for x, cell := range row {
			if cell == "." {
				continue
			}
			if cell == "S" {
				startingX = x
				startingY = y
				grid[y+1][x] = "|"
			}
			if cell == "|" {
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
			}
		}
		fmt.Printf("Row %d: %v\n", y, row)
	}

	fmt.Printf("Starting X: %d, Y: %d\n", startingX, startingY)
	fmt.Println("==========================")
	memo := make(map[string]int)
	timelineCount = countTimeline(grid, startingX, startingY, memo)

	fmt.Printf("Timeline count: %d", timelineCount)
}

func countTimeline(grid [][]string, x, y int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", x, y)
	if val, exists := memo[key]; exists {
		return val
	}

	if (x < 0 || x > len(grid[0])-1) && y != len(grid)-1 {
		return 0
	}
	if grid[y][x] != "|" && grid[y][x] != "S" {
		return 0
	}
	if y == len(grid)-1 && grid[y][x] == "|" {
		return 1
	}

	var result int
	if grid[y+1][x] == "^" {
		result = countTimeline(grid, x-1, y+1, memo) + countTimeline(grid, x+1, y+1, memo)
	} else {
		result = countTimeline(grid, x, y+1, memo)
	}

	memo[key] = result
	return result
}
