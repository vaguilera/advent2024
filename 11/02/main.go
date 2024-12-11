package main

import (
	"fmt"
	"math"
)

func main() {
	input := make(map[int]int)
	input[0] = 1
	input[5601550] = 1
	input[3914] = 1
	input[852] = 1
	input[50706] = 1
	input[68] = 1
	input[6] = 1
	input[645371] = 1

	for i := 0; i < 75; i++ {
		input = blink(input)
	}
	count := 0
	for _, v := range input {
		count += v
	}
	fmt.Printf("len: %d \n", count)
}

func blink(input map[int]int) map[int]int {
	res := make(map[int]int)
	for k, v := range input {
		pr := processNumber(k)

		for _, vpr := range pr {
			if _, ok := res[vpr]; ok {
				res[vpr] += v
			} else {
				res[vpr] = v
			}
		}
	}

	return res
}

func processNumber(v int) []int {
	if v == 0 {
		return []int{1}
	}

	digits := int(math.Log10(float64(v))) + 1

	if digits%2 == 0 {
		divisor := int(math.Pow(10, float64(digits/2)))
		firstHalf := v / divisor
		secondHalf := v % divisor
		return []int{firstHalf, secondHalf}
	}

	return []int{v * 2024}

}
