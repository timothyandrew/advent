package year2025

import (
	"fmt"

	"github.com/timothyandrew/advent/util"
)

type Position struct {
	x, y int
}

func adjacents(pos Position, rows, cols int) []Position {
	candidates := []Position{
		{x: pos.x - 1, y: pos.y},
		{x: pos.x + 1, y: pos.y},
		{x: pos.x, y: pos.y - 1},
		{x: pos.x, y: pos.y + 1},
		{x: pos.x + 1, y: pos.y - 1},
		{x: pos.x + 1, y: pos.y + 1},
		{x: pos.x - 1, y: pos.y - 1},
		{x: pos.x - 1, y: pos.y + 1},
	}

	results := []Position{}

	for _, c := range candidates {
		if c.x >= rows || c.y >= cols {
			continue
		}

		if c.x < 0 || c.y < 0 {
			continue
		}

		results = append(results, c)
	}

	return results
}

func (s Solutions) DayFourPartOne() error {
	m, err := util.FilenameToMap("2025/input/4/puzzle.txt")
	if err != nil {
		return err
	}

	results := make(map[Position]bool)

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] != '@' {
				continue
			}

			p := Position{x: i, y: j}
			adjCells := adjacents(p, len(m), len(m[0]))
			adjRolls := 0

			for _, a := range adjCells {
				if m[a.x][a.y] == '@' {
					adjRolls++
				}
			}

			if adjRolls < 4 {
				results[p] = true
			}
		}
	}

	fmt.Println(len(results))

	return nil
}

func (s Solutions) DayFourPartTwo() error {
	m, err := util.FilenameToMap("2025/input/4/puzzle.txt")
	if err != nil {
		return err
	}

	results := make(map[Position]bool)
	for {
		prevResultCount := len(results)

		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[0]); j++ {
				if m[i][j] != '@' {
					continue
				}

				p := Position{x: i, y: j}
				adjCells := adjacents(p, len(m), len(m[0]))
				adjRolls := 0

				for _, a := range adjCells {
					if m[a.x][a.y] == '@' {
						adjRolls++
					}
				}

				if adjRolls < 4 {
					results[p] = true
					m[i][j] = '.'
				}
			}
		}

		if len(results) == prevResultCount {
			break
		}
	}

	fmt.Println(len(results))

	return nil
}
