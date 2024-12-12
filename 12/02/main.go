package main

import (
	"fmt"
	"github.com/vaguilera/advent2024/utils"
	"slices"
)

type zone struct {
	area, perim int
	name        byte
}

type walker struct {
	x, y  int
	carea int
}

type casilla struct {
	walls int
	area  int
	name  byte
}

func main() {
	f, _ := utils.LoadFile()
	mapa := utils.NewMap(f)

	var visited [][]casilla
	for range mapa.M {
		lin := slices.Repeat([]casilla{{-1, 0, 0}}, len(mapa.M[0]))
		visited = append(visited, lin)
	}
	carea := 1
	areas := map[int]zone{}
	for y, cl := range visited {
		for x, cv := range cl {
			if cv.walls == -1 {
				walkers := []walker{{x, y, carea}}
				perimeter := 0
				for {
					if len(walkers) == 0 {
						break
					}
					w := walkers[0]
					walkers = walkers[1:]
					cc, _ := mapa.Get(w.x, w.y)
					nwalls := getWalls(&mapa, w.x, w.y)
					visited[w.y][w.x].walls = len(nwalls)
					visited[w.y][w.x].area = w.carea
					visited[w.y][w.x].name = cc

					perimeter += len(nwalls)

					c, ok := mapa.Get(w.x-1, w.y)
					if ok && visited[w.y][w.x-1].walls == -1 {
						if c == cc {
							if !findWalker(walkers, w.x-1, w.y) {
								walkers = append(walkers, walker{
									x:     w.x - 1,
									y:     w.y,
									carea: w.carea,
								})
							}
							wallsN := getWalls(&mapa, w.x-1, w.y)
							perimeter -= findMatchingWalls(nwalls, wallsN)
						}
					}

					c, ok = mapa.Get(w.x+1, w.y)
					if ok && visited[w.y][w.x+1].walls == -1 {
						if c == cc {
							if !findWalker(walkers, w.x+1, w.y) {
								walkers = append(walkers, walker{
									x:     w.x + 1,
									y:     w.y,
									carea: w.carea,
								})
							}
							wallsN := getWalls(&mapa, w.x+1, w.y)
							perimeter -= findMatchingWalls(nwalls, wallsN)
						}
					}
					c, ok = mapa.Get(w.x, w.y-1)
					if ok && visited[w.y-1][w.x].walls == -1 {
						if c == cc {
							if !findWalker(walkers, w.x, w.y-1) {
								walkers = append(walkers, walker{
									x:     w.x,
									y:     w.y - 1,
									carea: w.carea,
								})
							}
							wallsN := getWalls(&mapa, w.x, w.y-1)
							perimeter -= findMatchingWalls(nwalls, wallsN)
						}
					}
					c, ok = mapa.Get(w.x, w.y+1)
					if ok && visited[w.y+1][w.x].walls == -1 {
						if c == cc {
							if !findWalker(walkers, w.x, w.y+1) {
								walkers = append(walkers, walker{
									x:     w.x,
									y:     w.y + 1,
									carea: w.carea,
								})
							}
							wallsN := getWalls(&mapa, w.x, w.y+1)
							perimeter -= findMatchingWalls(nwalls, wallsN)
						}
					}
					ar, ok := areas[carea]
					if ok {
						ar.area += 1
						ar.perim = perimeter
						areas[carea] = ar

					} else {
						areas[carea] = zone{
							area:  1,
							name:  cc,
							perim: perimeter,
						}
					}

				}
				carea++
			}
		}
	}
	accum := 0
	for _, v := range areas {
		accum += v.area * v.perim
	}
	fmt.Printf("%d\n", accum)
}

func findWalker(walkers []walker, x, y int) bool {
	for _, w := range walkers {
		if w.x == x && w.y == y {
			return true
		}
	}
	return false
}

func getWalls(mapa *utils.Map, x, y int) string {

	walls := ""
	cc, _ := mapa.Get(x, y)
	c, ok := mapa.Get(x-1, y)
	if !ok || c != cc {
		walls += "L"
	}
	c, ok = mapa.Get(x, y-1)
	if !ok || c != cc {
		walls += "T"
	}

	c, ok = mapa.Get(x+1, y)
	if !ok || c != cc {
		walls += "R"
	}

	c, ok = mapa.Get(x, y+1)
	if !ok || c != cc {
		walls += "D"
	}
	return walls
}

func findMatchingWalls(current, neig string) int {
	res := 0
	for _, v := range current {
		for _, w := range neig {
			if v == w {
				res++
			}
		}
	}
	return res
}
