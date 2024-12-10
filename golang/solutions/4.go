package solutions

import (
	"fmt"
	"os"
	"strings"
)

var row_len, col_len = 0, 0

var xmas_len = 4

var directions = [8]string{"up", "down", "left", "right", "upleft", "upright", "downleft", "downright"}

func Solve_4() {
	chars := readInputToMatrix("data/4/input")
	row_len = len(chars)
	col_len = len(chars[0])

	fmt.Println("PART ONE ANSWER: ")
	fmt.Println(solvePartOne(chars) / 2)

	fmt.Println("PART TWO ANSWER: ")
	fmt.Println(solvePartTwo(chars))
}

func readInputToMatrix(filepath string) [][]rune {
	input, _ := os.ReadFile(filepath)
	chars := [][]rune{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		chars = append(chars, []rune(s))
	}
	return chars
}

func solvePartOne(chars [][]rune) int {
	count := 0
	for i, row := range chars {
		for j, char := range row {
			if !isStartingChar(char) {
				continue
			}
			count += countXMASPatterns(chars, i, j)
		}
	}
	return count
}

func isStartingChar(c rune) bool {
	return c == 'X' || c == 'S'
}

func countXMASPatterns(chars [][]rune, i, j int) int {
	count := 0
	w := WindowFunc{chars, i, j}

	for _, direction := range directions {
		word, err := w.getWindow(direction)
		if err != nil {
			continue
		}
		if isValidXMASPattern(word) {
			count++
		}
	}
	return count
}

func isValidXMASPattern(word string) bool {
	return word == "XMAS" || word == "SAMX"
}

func solvePartTwo(chars [][]rune) int {
	xmas_len = 3
	count := 0
	for i, row := range chars {
		for j, char := range row {
			if !isValidStartingPosition(char, i, j) {
				continue
			}
			if hasIntersectingPatterns(chars, i, j) {
				count++
			}
		}
	}
	return count
}

func isValidStartingPosition(char rune, i, j int) bool {
	return char == 'A' && i > 0 && j > 0
}

func hasIntersectingPatterns(chars [][]rune, i, j int) bool {
	w1 := WindowFunc{chars, i - 1, j - 1}
	w2 := WindowFunc{chars, i - 1, j + 1}

	word1, err1 := w1.getWindow("downright")
	if err1 != nil {
		return false
	}

	word2, err2 := w2.getWindow("downleft")
	if err2 != nil {
		return false
	}

	return isValidMASPattern(word1) && isValidMASPattern(word2)
}

func isValidMASPattern(word string) bool {
	return word == "MAS" || word == "SAM"
}

type WindowFunc struct {
	chars [][]rune
	row   int
	col   int
}

type Direction struct {
	row int
	col int
}

var dirs = map[string]Direction{
	"up":        {-1, 0},
	"down":      {1, 0},
	"left":      {0, -1},
	"right":     {0, 1},
	"upleft":    {-1, -1},
	"upright":   {-1, 1},
	"downleft":  {1, -1},
	"downright": {1, 1},
}

func (w *WindowFunc) getWindow(direction string) (string, error) {
	delta := dirs[direction]
	out := ""

	for r := 0; r < xmas_len; r++ {
		newRow := w.row + (delta.row * r)
		newCol := w.col + (delta.col * r)

		// Check bounds
		if newRow < 0 || newRow >= col_len || newCol < 0 || newCol >= row_len {
			return "", fmt.Errorf("out of bounds")
		}

		out += string(w.chars[newRow][newCol])
	}

	return out, nil
}
