package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day_6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]string
	var lists [][]int
	var operations []string

	for scanner.Scan() {
		line := scanner.Text()

		var row []string
		for _, ch := range line {
			if ch == '*' || ch == '+' {
				operations = append(operations, strings.Fields(line)...)
				break
			}
			row = append(row, string(ch))
		}

		if len(row) > 0 {
			grid = append(grid, row)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rowCount := len(grid)
	colCount := len(grid[0])

	var list []int
	for x := 0; x < colCount; x++ {
		num := ""

		for y := 0; y < rowCount; y++ {
			num += grid[y][x]
		}

		if strings.TrimSpace(num) == "" {
			lists = append(lists, list)
			list = []int{}
			continue
		}

		val, _ := strconv.Atoi(strings.TrimSpace(num))
		list = append(list, val)
		if x == colCount-1 {
			lists = append(lists, list)
		}
	}
	// fmt.Printf("Operations: %v\n", operations)
	// fmt.Printf("%v\n", lists)

	totalSum := 0

	for x, lst := range lists {
		partSum := 0
		for _, val := range lst {
			switch operations[x] {
			case "*":
				if partSum == 0 {
					partSum = 1
				}
				partSum *= val
			case "+":
				partSum += val
			}
		}
		totalSum += partSum
	}

	fmt.Println(totalSum)
}
