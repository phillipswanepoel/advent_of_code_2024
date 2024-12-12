package solutions

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
)

type order struct {
	before int
	after  int
}
type update struct {
	pages  []int
	graph  graph.Graph[int, int]
	sorted []int
}

func Solve_5() {
	orders, updates := readUpdates("data/5/input")

	for i := range updates {
		updates[i].graph = createGraph(orders, updates[i])
		updates[i].sorted, _ = graph.TopologicalSort(updates[i].graph)
	}

	ans_1 := solve5_part1(updates)
	fmt.Println("Part one answer: ")
	fmt.Println(ans_1)

	ans_2 := solve5_part2(updates)
	fmt.Println("Part two answer: ")
	fmt.Println(ans_2)
}

func solve5_part1(updates []update) int {
	sum := 0
	for _, u := range updates {
		if u.isValidUpdate() {
			middle := u.pages[len(u.pages)/2]
			sum += middle
		}
	}
	return sum
}

func solve5_part2(updates []update) int {
	sum := 0
	for _, u := range updates {
		if !(u.isValidUpdate()) {
			middle := u.sorted[len(u.sorted)/2]
			sum += middle
		}
	}
	return sum
}

func (u *update) isValidUpdate() bool {
	// Remove elements from sorted that aren't in update
	return slices.Equal(u.pages, u.sorted)
}

func createGraph(orders []order, update update) graph.Graph[int, int] {
	g := graph.New(graph.IntHash, graph.Directed(), graph.Acyclic(), graph.PreventCycles())

	for _, o := range orders {
		if slices.Contains(update.pages, o.before) && slices.Contains(update.pages, o.after) {
			_ = g.AddVertex(o.before)
			_ = g.AddVertex(o.after)
			_ = g.AddEdge(o.before, o.after)

		}
	}

	return g
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
