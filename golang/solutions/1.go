package solutions

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

// THIS SOLUTION COPY PASTED FROM REDDIT FOR REFERENCE
// FULL CREDIT TO: https://github.com/mnml

func Solve_1() {
	input, _ := os.ReadFile("data/1/input")

	var list1, list2 []int
	counts2 := map[int]int{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var n1, n2 int
		fmt.Sscanf(s, "%d %d", &n1, &n2)
		list1, list2 = append(list1, n1), append(list2, n2)
		counts2[n2]++
	}

	slices.Sort(list1)
	slices.Sort(list2)

	part1, part2 := 0, 0
	for i := range list1 {
		part1 += abs(list2[i] - list1[i])
		part2 += list1[i] * counts2[list1[i]]
	}
	fmt.Println(part1)
	fmt.Println(part2)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}
