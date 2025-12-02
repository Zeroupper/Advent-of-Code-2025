package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day_1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var rotations []int
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}
		
		direction := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		
		if direction == 'L' {
			rotations = append(rotations, -value)
		} else {
			rotations = append(rotations, value)
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Calculate password
	position := 50
	password := 0
	
	for _, rotation := range rotations {
		position = (position + rotation) % 100
		if position < 0 {
			position += 100
		}
		if position == 0 {
			password++
		}
	}

	fmt.Println("Password:", password)
}