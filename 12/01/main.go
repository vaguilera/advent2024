package main

import (
	"fmt"
	"github.com/vaguilera/advent2024/utils"
	"slices"
)

type group struct {
	area, walls int
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
	for y, cl := range visited {
		for x, cv := range cl {
			if cv.walls == -1 {
				walkers := []walker{{x, y, carea}}
				for {
					if len(walkers) == 0 {
						break
					}
					w := walkers[0]
					walkers = walkers[1:]
					cc, _ := mapa.Get(w.x, w.y)
					visited[w.y][w.x].walls = getWalls(&mapa, w.x, w.y)
					visited[w.y][w.x].area = w.carea
					visited[w.y][w.x].name = cc

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
						}
					}

				}
				carea++
			}
		}
	}

	areas := map[int]group{}
	for _, cl := range visited {
		for _, cv := range cl {
			if a, ok := areas[cv.area]; ok {
				a.walls += cv.walls
				a.area++
				areas[cv.area] = a
			} else {
				areas[cv.area] = group{
					area:  1,
					walls: cv.walls,
					name:  cv.name,
				}
			}
		}
	}

	count := 0
	for _, v := range areas {
		fmt.Printf("%s: %d %d\n", string(v.name), v.area, v.walls)
		count += v.walls * v.area
	}
	fmt.Println(count)
}

func findWalker(walkers []walker, x, y int) bool {
	for _, w := range walkers {
		if w.x == x && w.y == y {
			return true
		}
	}
	return false
}

func getWalls(mapa *utils.Map, x, y int) int {

	walls := 4
	cc, _ := mapa.Get(x, y)
	c, ok := mapa.Get(x-1, y)
	if ok {
		if c == cc {
			walls--
		}
	}
	c, ok = mapa.Get(x, y-1)
	if ok {
		if c == cc {
			walls--
		}
	}

	c, ok = mapa.Get(x+1, y)
	if ok {
		if c == cc {
			walls--
		}
	}
	c, ok = mapa.Get(x, y+1)
	if ok {
		if c == cc {
			walls--
		}

	}
	return walls
}
