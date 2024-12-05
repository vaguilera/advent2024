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
		v, v2 := checkUpdate(up)
		if v != -1 {
			fixed := fixSequence(up, v, v2)
			mid := fixed[len(fixed)/2]
			count += mid

		}
	}
	fmt.Printf("Part 2: %d\n", count)

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

func checkUpdate(update []int) (int, int) {
	for cpos := 0; cpos < len(update); cpos++ {
		cnum := update[cpos]
		for i := cpos + 1; i < len(update); i++ {
			if !findPair(cnum, update[i]) {
				if findPair(update[i], cnum) {
					return i, cpos
				}
			}

		}
	}
	return -1, -1
}

func findPair(cnum, up int) bool {
	for _, pair := range dicc {
		if pair[0] == cnum && pair[1] == up {
			//	fmt.Printf("cnum :%d, up: %d TRUE\n", cnum, up)
			return true
		}
	}
	//fmt.Printf("cnum :%d, up: %d FALSE\n", cnum, up)
	return false
}

func fixSequence(up []int, i, i2 int) []int {
	for {
		//	fmt.Printf("%v %d\n", up, i)
		tmp := up[i]
		up[i] = up[i2]
		up[i2] = tmp

		if i, i2 = checkUpdate(up); i == -1 {
			break
		}
	}
	//fmt.Printf("fixed %v %d\n", up, i)
	return up
}
