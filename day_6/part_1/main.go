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
	var lists [][]int
	var operations []string

	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)
		if columns[0] == "*" || columns[0] == "+" {
			operations = columns
			break
		}
		var row []int
		for _, col := range columns {

			num, err := strconv.Atoi(col)
			if err == nil {
				row = append(row, num)
			}
		}
		lists = append(lists, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalSum := 0

	for j, operation := range operations {
		partSum := 0
		for _, lst := range lists {
			switch operation {
			case "*":
				if partSum == 0 {
					partSum = 1
				}
				partSum *= lst[j]
			case "+":
				partSum += lst[j]
			}
		}
		totalSum += partSum
	}

	fmt.Println(totalSum)
}
