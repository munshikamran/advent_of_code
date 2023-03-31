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

func getMaximumCaloriesFromFile(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalCals float64 = 0.0
	var maxCals float64 = 0.0

	for scanner.Scan() {
		lineStr := scanner.Text()
		lineStr = strings.TrimSpace(lineStr)

		if len(lineStr) != 0 {
			currCals, _ := strconv.ParseFloat(lineStr, 64)
			totalCals += currCals
		} else {
			maxCals = math.Max(maxCals, totalCals)
			totalCals = 0.0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return int(maxCals)
}

func main() {
	answer := getMaximumCaloriesFromFile("input.txt")
	fmt.Println("The Elf with the maximum number of calories has ", answer, " calories!")
}
