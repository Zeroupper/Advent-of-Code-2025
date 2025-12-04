package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day_4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, len(line))
		for i, ch := range line {
			row[i] = string(ch)
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rollPaperCount := 0

	rows := len(grid)
	columns := len(grid[0])

	canWeRemoveRolls := true
	howManyDidWeRemove := 0

	for canWeRemoveRolls {
		rollPaperCount += howManyDidWeRemove
		howManyDidWeRemove = 0
		for y, row := range grid {
			for x := range row {
				if grid[y][x] != "@" {
					continue
				}
	
				numberOfAdjacents := 0
	
				for _, yD := range []int{-1, 0, 1} {
					for _, xD := range []int{-1, 0, 1} {
						
						if xD == 0 && yD == 0 {
							// fmt.Printf("Not counting itself\n")
							continue
						}
						if xD+x > columns-1 || xD+x < 0 {
							// fmt.Printf("x should be in the columns\n")
							continue
						}
						if yD+y > rows-1 || yD+y < 0 {
							// fmt.Printf("y should be in the rows\n")
							continue
						}
						if grid[y+yD][x+xD] == "@" {
							// fmt.Printf("Calculate neighbour with @ [%d][%d]\n",y+yD,x+xD)
							numberOfAdjacents++
						}
					}
				}
				// fmt.Printf("Found adjacents at coordinate: [%d,%d] na:%d \n", y, x, numberOfAdjacents)
				if numberOfAdjacents < 4 {
					howManyDidWeRemove++
					grid[y][x] = "x"
				}
			}
		}
		if howManyDidWeRemove == 0 {
			canWeRemoveRolls = false
		}
	}

	fmt.Println(rollPaperCount)
}
