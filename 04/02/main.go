package main

import (
	"fmt"
	"github.com/vaguilera/advent2024/utils"
)

func main() {
	f, _ := utils.LoadFile()
	mapa := utils.NewMap(f)

	l := len(mapa.M[0])
	count := 0

	for y := 1; y < len(mapa.M)-1; y++ {
		for x := 1; x < l-1; x++ {
			if mapa.M[y][x] == 'A' {
				b1 := mapa.GetD1Block(x-1, y-1, 3)
				if b1 != "MAS" && b1 != "SAM" {
					continue
				}
				b1 = mapa.GetD2Block(x+1, y-1, 3)
				if b1 != "MAS" && b1 != "SAM" {
					continue
				}
				count++
			}
		}
	}

	fmt.Printf("Part 2: %d\n", count)

}
