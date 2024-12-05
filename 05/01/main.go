package main

import (
	"fmt"
	"github.com/vaguilera/advent2024/utils"
	"strings"
)

var dicc [][2]int
var updates [][]int

func main() {
	f, _ := utils.LoadFile()

	parseInput(f)
	count := 0
	for _, up := range updates {
		v := checkUpdate(up)
		if v {
			mid := up[len(up)/2]
			count += mid
		}
	}
	fmt.Printf("Part 1: %d\n", count)

}

func parseInput(lines []string) {
	first := true
	for _, l := range lines {
		if first {
			if l == "" {
				first = false
				continue
			}
			val := strings.Split(l, "|")
			a := utils.Atoi(val[0])
			b := utils.Atoi(val[1])
			dicc = append(dicc, [2]int{a, b})
			continue
		}
		val := strings.Split(l, ",")
		up := []int{}
		for _, v := range val {
			up = append(up, utils.Atoi(v))
		}
		updates = append(updates, up)
	}
}

func checkUpdate(update []int) bool {
	for cpos := 0; cpos < len(update); cpos++ {
		cnum := update[cpos]
		for i := cpos + 1; i < len(update); i++ {
			if !findPair(cnum, update[i]) {
				if findPair(update[i], cnum) {
					return false
				}
			}

		}
	}
	return true
}

func findPair(cnum, up int) bool {
	for _, pair := range dicc {
		if pair[0] == cnum && pair[1] == up {
			//fmt.Printf("cnum :%d, up: %d TRUE\n", cnum, up)
			return true
		}
	}
	//fmt.Printf("cnum :%d, up: %d FALSE\n", cnum, up)
	return false
}
