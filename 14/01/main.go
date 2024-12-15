package main

import (
	"fmt"
	"strings"

	"github.com/vaguilera/advent2024/utils"
)

const width = 101
const height = 103

type robot struct {
	x, y   int
	vx, vy int
}

func main() {
	robots := []robot{}
	f, _ := utils.LoadFile()

	for _, l := range f {
		r := robot{}
		parts := strings.Split(l[2:], " ")
		pos := strings.Split(parts[0], ",")
		r.x = utils.Atoi(pos[0])
		r.y = utils.Atoi(pos[1])

		vels := strings.Split(parts[1][2:], ",")
		r.vx = utils.Atoi(vels[0])
		r.vy = utils.Atoi(vels[1])
		robots = append(robots, r)
	}

	for cr := range robots {
		for i := 0; i < 100; i++ {
			newX := (robots[cr].x + robots[cr].vx)
			if newX < 0 {
				newX = width + newX
			} else {
				newX = newX % width
			}
			robots[cr].x = newX

			newY := (robots[cr].y + robots[cr].vy)
			if newY < 0 {
				newY = height + newY
			} else {
				newY = newY % height
			}
			robots[cr].y = newY
		}

	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, r := range robots {
		if r.x < width/2 && r.y < height/2 {
			q1++
			continue
		}
		if r.x > width/2 && r.y < height/2 {
			q2++
			continue
		}
		if r.x < width/2 && r.y > height/2 {
			q3++
			continue
		}
		if r.x > width/2 && r.y > height/2 {
			q4++
			continue
		}

	}
	fmt.Printf("%d %d %d %d\n", q1, q2, q3, q4)
	fmt.Printf("Res: %d\n", q1*q2*q3*q4)

}
