package main

import (
	"fmt"

	"github.com/vaguilera/advent2024/utils"
)

type pos struct {
	x, y, dx, dy int
}

var mapa utils.Map

func main() {

	f, _ := utils.LoadFile()
	mapa = utils.NewMap(f)

	gx, gy := findguard(mapa.M)
	count := 0
	for cy, cl := range mapa.M {
		for cx, cc := range cl {
			if cc == '.' {
				mapa.M[cy] = utils.ReplaceAtIndex(mapa.M[cy], '#', cx)
				if isInfiniteLoop(gx, gy) {
					count++
				}
				mapa.M[cy] = utils.ReplaceAtIndex(mapa.M[cy], '.', cx)
			}
		}
	}

	fmt.Printf("Infinite loops: %d\n", count)

}

func isInfiniteLoop(guardX, guardY int) bool {
	visited := make(map[pos]bool)
	visited[pos{guardX, guardY, 0, -1}] = true
	cdir := pos{x: 0, y: -1}
	cx, cy := guardX, guardY

	var newx, newy int
	for {
		newx = cx + cdir.x
		newy = cy + cdir.y
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
			if visited[pos{cx, cy, cdir.x, cdir.y}] {
				return true
			}
			visited[pos{cx, cy, cdir.x, cdir.y}] = true
		} else {
			cx = newx
			cy = newy
			if visited[pos{cx, cy, cdir.x, cdir.y}] {
				return true
			}
			visited[pos{cx, cy, cdir.x, cdir.y}] = true
		}
	}
	return false
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
