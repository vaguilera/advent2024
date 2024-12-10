package main

import (
	"fmt"
	"github.com/vaguilera/advent2024/utils"
)

type walker struct {
	x      int
	y      int
	end    bool
	ox, oy int
}

type pos struct {
	x, y int
}

var walkers []walker
var mapa utils.Map

func main() {
	f, _ := utils.LoadFile()
	mapa = utils.NewMap(f)
	startPoints := getStartPoints(mapa)

	for _, sp := range startPoints {
		walkers = append(walkers, walker{sp.x, sp.y, false, sp.x, sp.y})
	}

	visited := make(map[pos][]pos)
	for {
		if len(walkers) == 0 {
			break
		}
		cw := walkers[len(walkers)-1]
		walkers = walkers[:len(walkers)-1]
		if cw.end {
			v, ok := visited[pos{cw.ox, cw.oy}]
			if !ok {
				visited[pos{cw.ox, cw.oy}] = []pos{{cw.x, cw.y}}
			} else {
				found := false
				for _, p := range v {
					if p.x == cw.x && p.y == cw.y {
						found = true
						break
					}
				}
				if !found {
					visited[pos{cw.ox, cw.oy}] = append(v, pos{cw.x, cw.y})
				}

			}
			continue
		}
		processWalker(cw)

	}
	count := 0
	for _, positions := range visited {
		count += len(positions)
	}
	fmt.Println(count)
}

func getStartPoints(mapa utils.Map) []walker {
	res := []walker{}
	for y, cl := range mapa.M {
		for x, cc := range cl {
			if cc == '0' {
				res = append(res, walker{x, y, false, x, y})
			}
		}
	}
	return res
}

func processWalker(w walker) {
	up := walker{w.x, w.y - 1, false, w.ox, w.oy}
	down := walker{w.x, w.y + 1, false, w.ox, w.oy}
	right := walker{w.x + 1, w.y, false, w.ox, w.oy}
	left := walker{w.x - 1, w.y, false, w.ox, w.oy}

	nw, moved := moveWalker(w, up)
	if moved {
		walkers = append(walkers, nw)
	}

	nw, moved = moveWalker(w, down)
	if moved {
		walkers = append(walkers, nw)
	}

	nw, moved = moveWalker(w, right)
	if moved {
		walkers = append(walkers, nw)
	}
	nw, moved = moveWalker(w, left)
	if moved {
		walkers = append(walkers, nw)
	}

}

func moveWalker(w walker, npos walker) (walker, bool) {
	if npos.x < 0 || npos.x >= len(mapa.M[0]) || npos.y < 0 || npos.y >= len(mapa.M) {
		return w, false
	}
	if mapa.M[npos.y][npos.x] < '0' || mapa.M[npos.y][npos.x] > '9' {
		return w, false
	}

	if mapa.M[npos.y][npos.x]-mapa.M[w.y][w.x] == 1 {
		if mapa.M[npos.y][npos.x] == '9' {
			npos.end = true
		}
		return npos, true
	}

	return w, false

}
