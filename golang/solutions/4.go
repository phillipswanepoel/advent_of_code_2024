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

func (w *WindowFunc) Up(r int) (rune, error) {
	if w.row-r < 0 {
		return 0, fmt.Errorf("out of bounds")
	}
	return w.chars[w.row-r][w.col], nil
}

func (w *WindowFunc) Down(r int) (rune, error) {
	if w.row+r >= col_len {
		return 0, fmt.Errorf("out of bounds")
	}
	return w.chars[w.row+r][w.col], nil
}

func (w *WindowFunc) Left(r int) (rune, error) {
	if w.col-r < 0 {
		return 0, fmt.Errorf("out of bounds")
	}
	return w.chars[w.row][w.col-r], nil
}

func (w *WindowFunc) DownLeft(r int) (rune, error) {
	if w.row+r >= col_len {
		return 0, fmt.Errorf("out of bounds")
	}
	if w.col-r < 0 {
		return 0, fmt.Errorf("out of bounds")
	}
	return w.chars[w.row+r][w.col-r], nil
}

func (w *WindowFunc) UpLeft(r int) (rune, error) {
	if w.row-r < 0 {
		return 0, fmt.Errorf("out of bounds")
	}
	if w.col-r < 0 {
		return 0, fmt.Errorf("out of bounds")
	}
	return w.chars[w.row-r][w.col-r], nil
}

func (w *WindowFunc) Right(r int) (rune, error) {
	if w.col+r >= row_len {
		return 0, fmt.Errorf("out of bounds")
	}
	return w.chars[w.row][w.col+r], nil
}

func (w *WindowFunc) DownRight(r int) (rune, error) {
	if w.row+r >= col_len {
		return 0, fmt.Errorf("out of bounds")
	}
	if w.col+r >= row_len {
		return 0, fmt.Errorf("out of bounds")
	}
	return w.chars[w.row+r][w.col+r], nil
}

func (w *WindowFunc) UpRight(r int) (rune, error) {
	if w.row-r < 0 {
		return 0, fmt.Errorf("out of bounds")
	}
	if w.col+r >= row_len {
		return 0, fmt.Errorf("out of bounds")
	}
	return w.chars[w.row-r][w.col+r], nil
}

func (w *WindowFunc) getWindow(direction string) (string, error) {
	out := ""
	for r := 0; r < xmas_len; r++ {
		var char rune
		var err error
		switch direction {
		case "up":
			char, err = w.Up(r)
			if err != nil {
				return "", fmt.Errorf("out of bounds: %v", err)
			}
		case "down":
			char, err = w.Down(r)
			if err != nil {
				return "", fmt.Errorf("out of bounds: %v", err)
			}
		case "left":
			char, err = w.Left(r)
			if err != nil {
				return "", fmt.Errorf("out of bounds: %v", err)
			}
		case "right":
			char, err = w.Right(r)
			if err != nil {
				return "", fmt.Errorf("out of bounds: %v", err)
			}
		case "upleft":
			char, err = w.UpLeft(r)
			if err != nil {
				return "", fmt.Errorf("out of bounds: %v", err)
			}
		case "upright":
			char, err = w.UpRight(r)
			if err != nil {
				return "", fmt.Errorf("out of bounds: %v", err)
			}
		case "downleft":
			char, err = w.DownLeft(r)
			if err != nil {
				return "", fmt.Errorf("out of bounds: %v", err)
			}
		case "downright":
			char, err = w.DownRight(r)
			if err != nil {
				return "", fmt.Errorf("out of bounds: %v", err)
			}
		}
		out += string(char)
	}

	return out, nil
}
