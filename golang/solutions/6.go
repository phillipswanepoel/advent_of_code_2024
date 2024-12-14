package solutions

import (
	"fmt"
	"os"
	"strings"
)

type cellType int

const (
	Empty cellType = iota // iota generates consecutive integers
	Obstacle
)

type pos struct {
	x, y int
}

func (p *pos) addPos(dir direction) {
	coords := delta[dir]
	p.x += coords.x
	p.y += coords.y
}

type direction uint8

const (
	Up direction = iota
	Right
	Down
	Left
)

func (d direction) turn() direction {
	return direction((int(d) + 1 + 4) % 4) // Add 4 to ensure no negative result before modulo.
}

type guard struct {
	current_pos pos
	dir         direction
}

func (g *guard) move() {
	g.current_pos.addPos(g.dir)
}

func (g *guard) turn() {
	g.dir = g.dir.turn()
}

type cell struct {
	cellType cellType
	char     rune
	dir      direction
	visited  bool
}

var delta = map[direction]pos{
	0: {x: 0, y: -1}, // up
	1: {x: 1, y: 0},  // right
	2: {x: 0, y: 1},  // down
	3: {x: -1, y: 0}, // left
}

func countMoves(grid map[pos]cell, g guard) int {
	for {
		old_pos := g.current_pos
		g.move()
		newCell, exists := grid[g.current_pos]
		if exists {
			switch newCell.cellType {
			case Obstacle:
				g.current_pos = old_pos
				g.turn()
				newCell.dir = g.dir

			case Empty:
				newCell.visited = true
				grid[g.current_pos] = newCell
			}
		} else {
			break
		}
	}

	sum := 0
	for key := range grid {
		cell := grid[key]
		if cell.visited {
			sum += 1
		}
	}
	return sum
}

func countLoops(grid map[pos]cell) int {
	visited := []pos{}
	for key, cell := range grid {
		if cell.visited {
			visited = append(visited, key)
		}
	}

	loops := 0
	for _, p := range visited {
		g1, g2 := readGrid("data/6/input")
		if doesLoop(g1, g2, p) {
			loops++
		}
	}

	return loops
}

func doesLoop(grid map[pos]cell, g guard, extraObs pos) bool {
	newObs := grid[extraObs]
	newObs.cellType = Obstacle
	grid[extraObs] = newObs

	for {
		old_pos := g.current_pos
		g.move()

		newCell, exists := grid[g.current_pos]
		if exists {
			if (newCell.dir == g.dir) && newCell.visited {
				return true
			}
			switch newCell.cellType {
			case Obstacle:
				g.current_pos = old_pos
				g.turn()

			case Empty:
				newCell.visited = true
				newCell.dir = g.dir
				grid[g.current_pos] = newCell
			}
		} else {
			return false
		}
	}
}

func readGrid(filepath string) (map[pos]cell, guard) {
	grid := make(map[pos]cell)
	g := guard{}
	input, _ := os.ReadFile(filepath)
	for i, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for j, c := range s {
			var t cellType
			switch string(c) {
			case ".":
				t = Empty
			case "#":
				t = Obstacle
			case "^":
				t = Empty
				g.dir = Up
				g.current_pos = pos{x: j, y: i}
			case ">":
				t = Empty
				g.dir = Right
				g.current_pos = pos{x: j, y: i}
			case "<":
				t = Empty
				g.dir = Left
				g.current_pos = pos{x: j, y: i}
			case "v":
				t = Empty
				g.dir = Down
				g.current_pos = pos{x: j, y: i}
			}
			grid[pos{x: j, y: i}] = cell{cellType: t, char: c}
		}
	}
	return grid, g
}

var (
	startGrid  map[pos]cell
	startGuard guard
)

func Solve_6() {
	grid, guard := readGrid("data/6/input")
	startGuard = guard
	startGrid = grid

	ans := countMoves(grid, guard)
	fmt.Println(ans)
	ans2 := countLoops(grid)
	fmt.Println(ans2)
}
