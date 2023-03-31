package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getElfCaloriesHeap(fileName string) *MaxHeap {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	calsHeap := &MaxHeap{}
	heap.Init(calsHeap)
	var totalCals float64 = 0.0

	for scanner.Scan() {
		lineStr := scanner.Text()
		lineStr = strings.TrimSpace(lineStr)

		if len(lineStr) != 0 {
			currCals, _ := strconv.ParseFloat(lineStr, 64)
			totalCals += currCals
		} else {
			heap.Push(calsHeap, int(totalCals))
			totalCals = 0.0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return calsHeap
}

func getTopFromHeap(calsHeap *MaxHeap, numValues int) []int {
	output := []int{}
	for i := 0; i < numValues; i++ {
		a := heap.Pop(calsHeap).(int)
		output = append(output, a)
	}
	return output
}

func main() {
	calsHeap := getElfCaloriesHeap("input.txt")
	top := getTopFromHeap(calsHeap, 3)
	fmt.Println("The Elf with the maximum number of calories has ", top, " calories!")
}
