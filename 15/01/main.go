package main

import (
	"fmt"

	"github.com/vaguilera/advent2024/utils"
)

type dupla struct {
	x, y int
}

var mapa []string
var inst string

func main() {
	mapa = []string{}
	inst = ""
	start := dupla{0, 0}
	f, _ := utils.LoadFile()
	parseFile(f)
	for y, l := range mapa {
		for x, cc := range l {
			if cc == '@' {
				start = dupla{x, y}
				mapa[y] = utils.ReplaceAtIndex(mapa[y], '.', x)
				break
			}
		}
	}

	for i := range inst {
		start = move(start, getMov(i))
	}
	printMapa(start)

	count := 0
	for y, l := range mapa {
		for x, cc := range l {
			if cc == 'O' {
				count += 100*y + x
			}
		}
	}

	fmt.Printf("Res: %d\n", count)

}

func printMapa(pos dupla) {
	for y, l := range mapa {
		for x, cc := range l {
			if x == pos.x && y == pos.y {
				fmt.Printf("%s", "@")
				continue
			}
			fmt.Printf("%s", string(cc))
		}
		fmt.Println()
	}
}

func parseFile(data []string) {
	mode := 0

	for _, l := range data {
		if l == "" {
			mode = 1
			continue
		}
		if mode == 0 {
			mapa = append(mapa, l)
			continue
		}
		inst += l
	}
}

func move(pos, dir dupla) dupla {
	npos := dupla{pos.x + dir.x, pos.y + dir.y}

	if mapa[npos.y][npos.x] == '#' {
		return pos
	}
	if mapa[npos.y][npos.x] == 'O' {
		if push(npos, dir) {
			mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], '.', npos.x)
			return npos
		}
		return pos

	}
	return npos
}

func push(pos, dir dupla) bool {
	npos := dupla{pos.x + dir.x, pos.y + dir.y}
	if mapa[npos.y][npos.x] == '#' {
		return false
	}
	if mapa[npos.y][npos.x] == '.' {
		mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], 'O', npos.x)
		return true
	}

	if push(npos, dir) {
		mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], 'O', npos.x)
		return true

	}
	return false
}

func getMov(cop int) dupla {
	switch inst[cop] {
	case '^':
		return dupla{0, -1}
	case '>':
		return dupla{1, 0}
	case '<':
		return dupla{-1, 0}
	case 'v':
		return dupla{0, 1}
	default:
		panic("invalid instruction")
	}
}
