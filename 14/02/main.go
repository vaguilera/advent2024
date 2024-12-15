package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
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

	test := 82
	for {
		canvas := []string{}
		robots := []robot{}
		f, _ := utils.LoadFile()

		str := strings.Repeat(".", width)
		for i := 0; i < height; i++ {
			canvas = append(canvas, str)
		}

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
			for i := 0; i < test; i++ {
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

		for _, r := range robots {
			tmpStr := utils.ReplaceAtIndex(canvas[r.y], 'X', r.x)
			canvas[r.y] = tmpStr

		}

		generateImage(canvas, test)
		if test > 8000 {
			break
		}

		test++
	}

}

func generateImage(datos []string, id int) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if datos[y][x] == 'X' {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, color.Black)
			}
		}
	}
	name := strconv.Itoa(id)
	f, _ := os.Create(name + ".png")
	png.Encode(f, img)
}
