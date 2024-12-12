package utils

type Map struct {
	M []string
}

func NewMap(lines []string) Map {
	return Map{
		M: lines,
	}
}

func (m *Map) Get(x, y int) (byte, bool) {
	if x < 0 || y < 0 {
		return 0, false
	}
	if y >= len(m.M) {
		return 0, false
	}
	if x >= len(m.M[0]) {
		return 0, false
	}

	return m.M[y][x], true
}

func (m *Map) GetXBlock(x, y, c int) string {
	return m.M[y][x : x+c]
}

func (m *Map) GetYBlock(x, y, c int) string {
	res := ""
	for i := 0; i < c; i++ {
		res += m.M[y+i][x : x+1]
	}
	return res
}

func (m *Map) GetD1Block(x, y, c int) string {
	res := ""
	for i := 0; i < c; i++ {
		res += m.M[y+i][x+i : x+1+i]
	}
	return res
}

func (m *Map) GetD2Block(x, y, c int) string {
	res := ""
	for i := 0; i < c; i++ {
		res += m.M[y+i][x-i : x-i+1]
	}
	return res
}
