package main

import (
	"fmt"
	"github.com/vaguilera/advent2024/utils"
	"strconv"
	"strings"
)

func main() {
	input := "0 5601550 3914 852 50706 68 6 645371"

	for i := 0; i < 25; i++ {
		input = blink(input)
	}
	fmt.Printf("len: %d \n", len(strings.Split(input, " ")))
}

func blink(input string) string {
	values := strings.Split(input, " ")

	res := []string{}
	for _, v := range values {
		if v == "0" {
			res = append(res, "1")
			continue
		}

		if len(v)%2 == 0 {
			res = append(res, removeLead(v[:len(v)/2]))
			res = append(res, removeLead(v[len(v)/2:]))
			continue
		}

		n := utils.Atoi(v)
		res = append(res, strconv.Itoa(n*2024))
	}
	return strings.Join(res, " ")
}

func removeLead(str string) string {
	var i int
	for i = range str {
		if str[i] != '0' {
			break
		}
	}
	return str[i:]
}
