package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/vaguilera/advent2024/utils"
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
	count := 0
	premios := 0
	for i := range machines {
		res := solve(machines[i])
		if res != 0 {
			premios++
			count += res
		} else {
			fmt.Printf("%#v\n", machines[i])
			os.Exit(0)
		}
	}

	fmt.Println(count)
	fmt.Println(premios)
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
				results = append(results, dupla{x, y})
			}
		}
		x++
		if x > 100 {
			break
		}
	}
	if len(results) == 0 {
		return 0
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
