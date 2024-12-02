package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve_2() {
	input, _ := os.ReadFile("data/2/input")

	var reports [][]int
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var words = strings.Fields(s)
		var levels []int
		for _, w := range words {
			levelInt, _ := strconv.Atoi(w)
			levels = append(levels, levelInt)
		}
		reports = append(reports, levels)
	}

	var sum = 0
	for _, level := range reports {
		if isValid(level) {
			sum += 1
		}
	}
	fmt.Println(sum)

	var sum2 = 0
	for _, level := range reports {
		var valid = false
		for i := range level {
			var newLevel []int
			newLevel = append(newLevel, level[:i]...)
			newLevel = append(newLevel, level[i+1:]...)
			if isValid(newLevel) {
				valid = true
			}
		}

		if valid {
			sum2 += 1
		}
	}
	fmt.Println(sum2)

}

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
func validDistance(a int, b int) bool {
	var dist = abs(a - b)
	return (dist >= 1) && (dist <= 3)
}

func isValid(input []int) bool {
	firstSign := sign(input[1] - input[0])
	firstDist := validDistance(input[1], input[0])
	if !firstDist {
		return false
	}

	for i := 2; i < len(input); i++ {
		if sign(input[i]-input[i-1]) != firstSign {
			return false
		}
		if !(validDistance(input[i], input[i-1])) {
			return false
		}
	}

	return true
}
