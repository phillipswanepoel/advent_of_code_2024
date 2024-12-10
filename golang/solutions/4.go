package solutions

import (
	"fmt"
	"os"
	"strings"
)

var row_len, col_len = 0, 0

const xmas_len = 3

var directions = [8]string{"up", "down", "left", "right", "upleft", "upright", "downleft", "downright"}

func Solve_4() {
	input, _ := os.ReadFile("data/4/input")

	chars := [][]rune{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		runes := []rune(s)
		chars = append(chars, runes)
	}

	row_len = len(chars)
	col_len = len(chars[0])

	var count = 0

	for i, r := range chars {
		for j, c := range r {
			if c == 'X' || c == 'S' {
				for _, d := range directions {
					w := WindowFunc{chars, i, j}
					word, err := w.getWindow(d)
					if err != nil {
						continue
					}
					if word == "XMAS" || word == "SAMX" {
						count++
					}
				}

			}
		}
	}

	fmt.Println("PART ONE ANSWER: ")
	fmt.Println(count / 2)

	// PART TWO!
	count2 := 0

	for i, r := range chars {
		for j, c := range r {
			if c == 'A' && (i > 0) && (j > 0) {
				w1 := WindowFunc{chars, i - 1, j - 1}
				w2 := WindowFunc{chars, i - 1, j + 1}
				word1, err := w1.getWindow("downright")
				if err != nil {
					continue
				}
				word2, err := w2.getWindow("downleft")
				if err != nil {
					continue
				}
				if (word1 == "MAS" || word1 == "SAM") && (word2 == "MAS" || word2 == "SAM") {
					count2++
				}
			}
		}
	}

	fmt.Println("PART TWO ANSWER: ")
	fmt.Println(count2)
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
