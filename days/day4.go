package days

import (
	"aoc2025/internal"
	"strings"
)

func DayFourPartOne(input string) int {
	input_grid := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	pickable_cells := 0

	for y, line := range input_grid {
		for x, char := range line {

			if char == '@' && internal.CountTotalNeighbouringRolls(&input_grid, x, y) < 4 {
				pickable_cells++
			}
		}
	}
	return pickable_cells
}

func DayFourPartTwo(input string) int {
	old_grid := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	new_grid := make([]string, len(old_grid))
	copy(new_grid, old_grid)

	flag := true
	total := 0

	for flag {
		pickable_cells := 0

		for y, line := range old_grid {
			for x, char := range line {

				if char == '@' && internal.CountTotalNeighbouringRolls(&old_grid, x, y) < 4 {
					line := new_grid[y]
					new_grid[y] = line[:x] + "x" + line[x+1:]
					pickable_cells++
				}
			}
		}
		// fmt.Println(pickable_cells)
		// for _,line :=range new_grid
		if pickable_cells <= 0 {
			flag = false
		}
		total += pickable_cells
		copy(old_grid, new_grid)

	}

	return total
}
