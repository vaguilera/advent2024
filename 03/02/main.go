package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/vaguilera/advent2024/utils"
)

func main() {
	f, _ := utils.LoadFile()
	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)

	//input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	matches := re.FindAllString(strings.Join(f, ""), -1)

	count := 0
	enabled := true
	for _, match := range matches {
		if match == "do()" {
			enabled = true
			continue
		}
		if match == "don't()" {
			enabled = false
			continue
		}
		numbers := strings.Split(match[4:len(match)-1], ",")
		x := utils.Atoi(numbers[0])
		y := utils.Atoi(numbers[1])
		if enabled {
			count += x * y
		}
	}

	fmt.Printf("Part 1: %d\n", count)

}
