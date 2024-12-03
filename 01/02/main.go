package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vaguilera/advent2024/utils"
)

func main() {
	text, _ := utils.LoadFile()

	listA := []int{}
	listB := []int{}

	for _, line := range text {
		nums := strings.Split(line, "   ")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		listA = append(listA, a)
		listB = append(listB, b)
	}

	final := 0
	for _, aValue := range listA {
		count := 0
		for _, bValue := range listB {
			if aValue == bValue {
				count++
			}
		}
		final += count * aValue
	}

	fmt.Printf("Result: %d\n", final)
}
