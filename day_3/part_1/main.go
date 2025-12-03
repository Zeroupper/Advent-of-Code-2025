package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day_3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var banks []string
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}
		
		banks = append(banks, line)
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Calculate password
	joltageSum := 0

	for _, bank := range banks {
		// fmt.Println(bank)
		largeBattery := 0
		smallBattery := 0
		for i, char := range bank {
			length := len(bank)
			value, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}

			if value > largeBattery {
				if(i < length-1) {
					largeBattery = value
					smallBattery = 0
					continue
				}
			}
			if value > smallBattery {
				smallBattery = value
			}
		}
		str := fmt.Sprintf("%d%d", largeBattery, smallBattery)
		joltage, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(joltage)
		joltageSum += joltage
	}

	fmt.Println("joltageSum:", joltageSum)
}