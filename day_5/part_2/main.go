package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
	valid      bool
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
		freshRanges = append(freshRanges, Range{start, end, true})
	}

	freshCount := 0

	var intervals []Range

	for _, r := range freshRanges {
		isNewInterval := true
		newStart := r.Start
		newEnd := r.End

		for j, i := range intervals {
			if !i.valid {
				continue
			}
			if newStart >= i.Start && newStart <= i.End || newEnd >= i.Start && newEnd <= i.End {
				newStart = int(math.Min(float64(newStart), float64(i.Start)))
				newEnd = int(math.Max(float64(newEnd), float64(i.End)))
				isNewInterval = false
				intervals[j].valid = false

			}
			if newStart < i.Start && newEnd > i.End {
				intervals[j].valid = false
			}
		}

		if isNewInterval {
			intervals = append(intervals, r)
		} else {
			intervals = append(intervals, Range{newStart, newEnd, true})
		}
	}

	for _, i := range intervals {
		if !i.valid {
			continue
		}
		count := i.End - i.Start + 1
		freshCount += count
	}

	fmt.Println(freshCount)
}
