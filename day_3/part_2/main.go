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
		length := len(bank)

		maxBatteries := [12]int{}
		maxBatteryLength := len(maxBatteries)

		for i, char := range bank {
			value, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}
			for j, maxBattery := range maxBatteries {
				if i < length-(maxBatteryLength-j-1) {

					if value > maxBattery {
						maxBatteries[j] = value
						for k := j + 1; k < maxBatteryLength; k++ {
							maxBatteries[k] = 0
						}
						break
					}
				}
			}
		}

		str := ""
		for _, maxBattery := range maxBatteries {
			str += strconv.Itoa(maxBattery)
		}

		joltage, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		joltageSum += joltage
	}

	fmt.Println("joltageSum:", joltageSum)
}
