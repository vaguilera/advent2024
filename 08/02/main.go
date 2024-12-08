package main

import (
	"fmt"

	"github.com/vaguilera/advent2024/utils"
)

type pos struct {
	x, y int
}

var freq map[rune][]pos
var visited map[pos]bool
var mapa utils.Map

func main() {
	f, _ := utils.LoadFile()
	mapa = utils.NewMap(f)

	visited = make(map[pos]bool)
	freq = make(map[rune][]pos)

	for y, cl := range mapa.M {
		for x, cc := range cl {
			if cc != '.' {
				freq[cc] = append(freq[cc], pos{x, y})
			}
		}
	}

	for _, f := range freq {
		for ca, a := range f {
			for i := 0; i < len(f); i++ {
				if i != ca {
					procesPair(a, f[i])
				}

			}

		}
	}

	v := len(visited)
	for _, f := range freq {
		for _, a := range f {
			if !visited[pos{a.x, a.y}] {
				v++
			}

		}
	}
	fmt.Printf("Result: %d\n", v)

}

func procesPair(a, b pos) {
	dist := getDist(a, b)
	antipos := pos{a.x, a.y}
	for {
		antipos = pos{antipos.x + dist.x, antipos.y + dist.y}
		if antipos.x < 0 || antipos.x > len(mapa.M[0])-1 || antipos.y < 0 || antipos.y > len(mapa.M)-1 {
			return
		}
		visited[antipos] = true
	}
}

func getDist(a, b pos) pos {
	x := a.x - b.x
	y := a.y - b.y

	return pos{x, y}
}
