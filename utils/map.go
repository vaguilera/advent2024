package utils

type Map struct {
	m []string
}

func New(lines []string) Map {
	return Map{
		m: lines,
	}
}

func (m *Map) Get(x, y int) byte {
	if x < 0 || y < 0 {
		panic("X or Y less than 0")
	}
	if y > len(m.m) {
		panic("Y out of bounds")
	}
	if x > len(m.m[0]) {
		panic("X out of bounds")
	}

	return m.m[y][x]
}
