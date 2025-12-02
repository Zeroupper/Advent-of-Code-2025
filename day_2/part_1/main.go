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

func isInvalid(id int) bool {
	s := strconv.Itoa(id)
	length := len(s)
	
	// Skip odd length numbers
	if length % 2 != 0 {
		return false
	}
	
	mid := length / 2
	left := s[:mid]
	right := s[mid:]
	
	return left == right
}

func main() {
	file, err := os.Open("day_2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	
	parts := strings.Split(line, ",")
	ranges := make([]Range, 0, len(parts))
	
	for _, part := range parts {
		nums := strings.Split(part, "-")
		start, _ := strconv.Atoi(nums[0])
		end, _ := strconv.Atoi(nums[1])
		ranges = append(ranges, Range{start, end})
	}
	
	sum := 0
	for _, r := range ranges {
		for id := r.Start; id <= r.End; id++ {
			if isInvalid(id) {
				sum += id
			}
		}
	}
	
	fmt.Println("Sum:", sum)
}