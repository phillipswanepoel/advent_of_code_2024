package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type order struct {
	before int
	after  int
}
type update struct {
	pages []int
}

func Solve_5() {
	orders, updates := readUpdates("data/5/test")

	for _, order := range orders {
		fmt.Printf("%d|%d", order.before, order.after)
		fmt.Println()
	}

	for _, update := range updates {
		fmt.Println(update)
	}

	// fmt.Println("PART ONE ANSWER: ")
	// fmt.Println(solvePartOne(chars) / 2)

	// fmt.Println("PART TWO ANSWER: ")
	// fmt.Println(solvePartTwo(chars))
}

func readUpdates(filepath string) ([]order, []update) {
	orders := []order{}
	updates := []update{}

	input, _ := os.ReadFile(filepath)
	mode := true
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		if strings.TrimSpace(s) == "" {
			mode = false
			continue
		}
		if mode {
			var n1, n2 int
			fmt.Sscanf(s, "%d|%d", &n1, &n2)
			orders = append(orders, order{before: n1, after: n2})
		} else {
			nums_strings := strings.Split(s, ",")
			var numbers []int
			for _, n := range nums_strings {
				num, err := strconv.Atoi(n)
				if err != nil {
					fmt.Println("Error parsing number:", err)
					continue
				}
				numbers = append(numbers, num)
			}
			updates = append(updates, update{pages: numbers})

		}
	}

	return orders, updates
}
