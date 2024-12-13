package main

import (
	"fmt"
	"github.com/vaguilera/advent2024/utils"
	"math"
	"strings"
)

type dupla struct {
	x, y float64
}

type machine struct {
	ButtonA, ButtonB, Prize dupla
}

func main() {
	f, _ := utils.LoadFile()
	machines := []machine{}
	for i := 0; i < len(f); i++ {
		AX, AY := getButtonNumbers(f[i])
		i++
		BX, BY := getButtonNumbers(f[i])
		i++
		PX, PY := getButtonNumbers(f[i])
		i++
		machines = append(machines, machine{
			ButtonA: dupla{float64(AX), float64(AY)},
			ButtonB: dupla{float64(BX), float64(BY)},
			Prize:   dupla{float64(PX), float64(PY)},
		})
	}
	//count := 0
	//for i := range machines {
	solve(machines[0])
	//}

	//fmt.Println(count)
}

func getButtonNumbers(s string) (int, int) {
	sstr := strings.Split(s, ": ")
	nums := strings.Split(sstr[1], ", ")
	numX := utils.Atoi(nums[0][2:])
	numY := utils.Atoi(nums[1][2:])

	return numX, numY
}

func solve(m machine) int {
	x := float64(0)
	y := float64(0)
	results := []dupla{}
	for {
		y = (m.Prize.x - (m.ButtonA.x * x)) / m.ButtonB.x
		if y < 0 {
			break
		}
		if y == math.Trunc(y) {
			y2 := m.ButtonA.y*x + m.ButtonB.y*y
			if y2 == m.Prize.y {
				//fmt.Printf("x: %d y:%d\n", int(x), int(y))
				//mult := (m.Prize.x + 100) / x
				//mult2 := (m.Prize.y + 100) / y
				results = append(results, dupla{x, y})
				//return int(x), int(y)
			}
		}
		x++
		//if x > 100 {
		//	break
		//}
	}
	if len(results) == 0 {
		return 0
	}

	for _, res := range results {
		fmt.Printf("%d %d %d\n", int(res.x), int(res.y), int(res.x*3+res.y))
	}
	maxT := 1000000
	for _, res := range results {
		tokens := res.x*3 + res.y
		if int(tokens) < maxT {
			maxT = int(tokens)
		}
	}

	return maxT

}
