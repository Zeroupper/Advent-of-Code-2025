package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	var ranges []Range
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		nums := strings.Split(line, "-")
		start, _ := strconv.Atoi(nums[0])
		end, _ := strconv.Atoi(nums[1])
		ranges = append(ranges, Range{start, end})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := []Range{ranges[0]}
	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]
		if r.Start <= last.End {
			if r.End > last.End {
				last.End = r.End
			}
		} else {
			merged = append(merged, r)
		}
	}

	total := 0
	for _, m := range merged {
		total += m.End - m.Start + 1
	}

	fmt.Println(total)
}