package main

import (
	"fmt"
	"sort"
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

	sort.Ints(listA)
	sort.Ints(listB)

	count := 0
	for i := range listA {
		count += utils.Abs(listA[i] - listB[i])
	}

	fmt.Printf("Result: %d\n", count)

}
