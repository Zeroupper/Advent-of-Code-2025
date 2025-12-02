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
	
	// Input file read, create rotation list from each line 
	for scanner.Scan() {
		line := scanner.Text()
		
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
		fullRotations := rotation / 100
		
		newPosition := (position + rotation) % 100
		if newPosition < 0 {
			newPosition += 100
		}
		if fullRotations < 0 {
			fullRotations = -fullRotations
		}
		if position != 0 && (rotation > 0 && newPosition <= position || rotation < 0 && (newPosition >= position || newPosition == 0)) {
			password++
		}
		password += fullRotations
		position = newPosition
	}

	fmt.Println("Password:", password)
}
