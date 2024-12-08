package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vaguilera/advent2024/utils"
)

var final int

func main() {
	f, _ := utils.LoadFile()

	sum := 0
	for _, cl := range f {
		nums := []int{}
		splited := strings.Split(cl, ": ")
		res := utils.Atoi(splited[0])
		final = res
		numsStr := strings.Split(splited[1], " ")
		for _, n := range numsStr {
			nums = append(nums, utils.Atoi(n))
		}

		r := testRec(nums[0], nums[1:])
		if r {
			sum += final
		}
	}
	fmt.Printf("Result: %d\n", sum)
}

func testRec(acc int, nums []int) bool {
	if len(nums) == 0 {
		return acc == final
	}
	if testRec(acc*nums[0], nums[1:]) {
		return true
	}
	if testRec(acc+nums[0], nums[1:]) {
		return true
	}

	accStr := strconv.Itoa(acc)
	n0 := strconv.Itoa(nums[0])

	return testRec(utils.Atoi(accStr+n0), nums[1:])
}
