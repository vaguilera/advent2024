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

	// count := 0
	// i := 0
	// reader := bufio.NewReader(os.Stdin)
	// inst = ""
	// for {
	// 	char, _, _ := reader.ReadRune()
	// 	if char != 'a' && char != 'd' && char != 's' && char != 'w' {
	// 		continue
	// 	}
	// 	switch char {
	// 	case 'a':
	// 		inst = "<"
	// 	case 'd':
	// 		inst = ">"
	// 	case 's':
	// 		inst = "v"
	// 	case 'w':
	// 		inst = "^"
	// 	}
	// 	start = move(start, getMov(i))
	// 	printMapa(start)
	// 	//	time.Sleep(500 * time.Millisecond)
	// }

	for i := range inst {
		printMapa(start)
		start = move(start, getMov(i))

	}
	printMapa(start)

	// count := 0
	// for y, l := range mapa {
	// 	for x, cc := range l {
	// 		if cc == 'O' {
	// 			count += 100*y + x
	// 		}
	// 	}
	// }

	// fmt.Printf("Res: %d\n", count)

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
			cline := ""
			for _, cc := range l {
				switch cc {
				case '#':
					cline += "##"
				case '.':
					cline += ".."
				case 'O':
					cline += "[]"
				case '@':
					cline += "@."
				}
			}
			mapa = append(mapa, cline)
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
	if mapa[npos.y][npos.x] == '[' || mapa[npos.y][npos.x] == ']' {
		if push(npos, dir) {
			if dir.y == 1 || dir.y == -1 {
				if mapa[npos.y][npos.x] == '[' {
					mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], '.', npos.x)
					mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], '.', npos.x+1)
				} else {
					mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], '.', npos.x)
					mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], '.', npos.x-1)
				}
			} else {
				mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], '.', npos.x)
			}
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
		if dir.y == 1 || dir.y == -1 {
			if mapa[pos.y][pos.x] == '[' {
				if mapa[npos.y][npos.x+1] != '.' {
					return false
				}
			}
			if mapa[pos.y][pos.x] == ']' {
				if mapa[npos.y][npos.x-1] != '.' {
					return false
				}
			}
			mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], rune(mapa[pos.y][pos.x]), npos.x)
			if mapa[pos.y][pos.x] == '[' {
				mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], rune(mapa[pos.y][pos.x+1]), npos.x+1)
			} else {
				mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], rune(mapa[pos.y][pos.x-1]), npos.x-1)
			}
		} else {
			mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], rune(mapa[pos.y][pos.x]), npos.x)
		}
		return true
	}

	// if dir.y == 1 || dir.y == -1 {
	// 	if mapa[npos.y][npos.x] != mapa[pos.y][pos.x] {
	// 		return false
	// 	}
	// }

	if dir.y == 1 || dir.y == -1 {
		pair := 0
		if mapa[pos.y][pos.x] == ']' {
			pair = -1
		} else {
			pair = 1
		}
		if push(npos, dir) && push(dupla{npos.x + pair, npos.y}, dir) {
			mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], rune(mapa[pos.y][pos.x]), npos.x)
			return true

		}
	} else {
		if push(npos, dir) {
			mapa[npos.y] = utils.ReplaceAtIndex(mapa[npos.y], rune(mapa[pos.y][pos.x]), npos.x)
			return true

		}
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
