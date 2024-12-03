package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vaguilera/advent2024/utils"
)

func main() {
	text, _ := utils.LoadFile()

	count := 0
	for _, line := range text {
		numsStr := strings.Split(line, " ")
		nums := []int{}
		for _, n := range numsStr {
			a, _ := strconv.Atoi(n)
			nums = append(nums, a)
		}

		if validSequence(nums) {
			count++
		}

	}

	fmt.Printf("Result: %d\n", count)

}

func validSequence(nums []int) bool {
	asc := true
	if nums[1] < nums[0] {
		asc = false
	}

	for i := 1; i < len(nums); i++ {
		if (nums[i-1] < nums[i]) && !asc {
			return false
		}
		if (nums[i-1] > nums[i]) && asc {
			return false
		}
		delta := utils.Abs(nums[i] - nums[i-1])
		if delta < 1 || delta > 3 {
			return false
		}
	}
	return true
}
