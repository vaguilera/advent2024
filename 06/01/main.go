package main

import (
	"fmt"

	"github.com/vaguilera/advent2024/utils"
)

type pos struct {
	x, y int
}

var mapa utils.Map
var visited map[pos]bool

func main() {
	visited = make(map[pos]bool)
	cdir := pos{x: 0, y: -1}
	f, _ := utils.LoadFile()
	mapa = utils.NewMap(f)

	x, y := findguard(mapa.M)
	visited[pos{x, y}] = true

	fmt.Printf("x: %d, y: %d\n", x, y)

	var newx, newy int
	for {
		newx = x + cdir.x
		newy = y + cdir.y
		if newx < 0 || newx > len(mapa.M[0])-1 {
			break
		}
		if newy < 0 || newy > len(mapa.M)-1 {
			break
		}

		if mapa.M[newy][newx] == '#' {
			if cdir.y == -1 {
				cdir.y = 0
				cdir.x = 1
			} else if cdir.x == 1 {
				cdir.x = 0
				cdir.y = 1
			} else if cdir.y == 1 {
				cdir.y = 0
				cdir.x = -1
			} else {
				cdir.y = -1
				cdir.x = 0
			}
		} else {
			x = newx
			y = newy
			visited[pos{x, y}] = true
		}
	}

	fmt.Printf("part 1: %d\n", len(visited))

}

func findguard(m []string) (int, int) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			if m[y][x] == '^' {
				return x, y
			}
		}
	}

	panic("No guard found")
}
