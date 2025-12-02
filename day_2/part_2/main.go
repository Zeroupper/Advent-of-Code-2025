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

func isInvalid(id int, divisorMap map[int][]int) bool {
	s := strconv.Itoa(id)
	length := len(s)

	// Get divisors for this length (excluding 1 and the length itself for patterns)
	divisors := divisorMap[length]
	// Check each divisor to see if the string is made of repeating patterns
	for _, div := range divisors {
		pattern := s[:div]
		valid := true

		// Check if entire string is made of this pattern repeated
		for i := 0; i < length; i += div {
			end := i + div
			if s[i:end] != pattern {
				valid = false
				break
			}
		}

		if valid {
			return true
		}
	}

	return false
}

// New function from patch
func getAllDivisors(n int) []int {
	divisors := []int{}
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
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

	maxNum := 0
	for _, part := range parts {
		nums := strings.Split(part, "-")
		start, _ := strconv.Atoi(nums[0])
		end, _ := strconv.Atoi(nums[1])
		ranges = append(ranges, Range{start, end})
		if end > maxNum {
			maxNum = end
		}
	}

	// Create divisor map for all number lengths up to max
	maxLen := len(strconv.Itoa(maxNum))
	divisorMap := make(map[int][]int)

	for length := 2; length <= maxLen; length++ {
		divisors := getAllDivisors(length)
		divisorMap[length] = divisors
	}

	fmt.Println("divisorMap:", divisorMap)
	

	sum := 0
	for _, r := range ranges {
		for id := r.Start; id <= r.End; id++ {
			if isInvalid(id, divisorMap) {
				// fmt.Println("Invalid id: ", id)
				sum += id
			}
		}
	}

	fmt.Println("Sum:", sum)
}
