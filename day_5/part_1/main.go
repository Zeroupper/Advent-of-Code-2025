package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
}

func main() {
	file, err := os.Open("day_5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var freshRanges []Range
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		nums := strings.Split(line, "-")
		start, _ := strconv.Atoi(nums[0])
		end, _ := strconv.Atoi(nums[1])
		freshRanges = append(freshRanges, Range{start, end})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	freshCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		for _, r := range freshRanges {
			if num >= r.Start && num <= r.End {
				freshCount++
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(freshCount)
}
