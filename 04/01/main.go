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
	for i := range mapa.M {
		for j := 0; j < l-3; j++ {
			block := mapa.GetXBlock(j, i, 4)
			if block == "XMAS" || block == "SAMX" {
				count++
			}
		}
	}

	for x := 0; x < l; x++ {
		for y := 0; y < len(mapa.M)-3; y++ {
			block := mapa.GetYBlock(x, y, 4)
			if block == "XMAS" || block == "SAMX" {
				count++
			}
		}
	}

	for x := 0; x < l-3; x++ {
		for y := 0; y < len(mapa.M)-3; y++ {
			block := mapa.GetD1Block(x, y, 4)
			if block == "XMAS" || block == "SAMX" {
				count++
			}
		}
	}

	for x := 3; x < l; x++ {
		for y := 0; y < len(mapa.M)-3; y++ {
			block := mapa.GetD2Block(x, y, 4)
			if block == "XMAS" || block == "SAMX" {
				count++
			}
		}
	}
	fmt.Printf("Part 1: %d\n", count)

}
