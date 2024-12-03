package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func LoadFile() ([]string, int) {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	lines := 0
	for scanner.Scan() {
		text = append(text, scanner.Text())
		lines++
	}

	fmt.Printf("Lines read: %d\n", lines)
	return text, lines
}

func LoadFile2() ([]string, int) {
	file, err := os.Open("../input2.txt")
	if err != nil {
		log.Fatalf("failed to open input.txt")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	lines := 0
	for scanner.Scan() {
		text = append(text, scanner.Text())
		lines++
	}

	fmt.Printf("Lines read: %d\n", lines)
	return text, lines
}

func FileToNumbers(text []string) []int {
	var numbers []int

	for _, strnum := range text {
		i, _ := strconv.Atoi(strnum)
		numbers = append(numbers, i)
	}

	fmt.Printf("Numbers parsed: %d\n", len(numbers))
	return numbers
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinMax(x, y int) (int, int) {
	if x < y {
		return x, y
	}
	return y, x
}

func BitsToInt(s string) uint64 {
	if i, err := strconv.ParseUint(s, 2, 64); err != nil {
		panic(fmt.Sprintf("Error converting bits: %s", s))
	} else {
		return i
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func RemoveFromSlice(slice []int, i int) []int {
	newSlice := make([]int, len(slice)-1)
	copy(newSlice, slice[:i])
	copy(newSlice[i:], slice[i+1:])
	return newSlice
}

func Atoi(i string) int {
	r, _ := strconv.Atoi(i)
	return r
}
