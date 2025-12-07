package days

import (
	"aoc2025/internal"
	"strconv"
	"strings"
)

func DayFivePartOne(input string) int {
	raw_input := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n\n")
	ranges_raw := strings.Split(raw_input[internal.INDEX_ZERO], "\n")
	ingredients := strings.Split(raw_input[internal.INDEX_ONE], "\n")
	counter := 0

	// Convert the ranges from string to [2]int
	ranges := make([][2]int, len(ranges_raw))
	for idx, val := range ranges_raw {
		rng := strings.Split(val, "-")
		start, _ := strconv.Atoi(rng[internal.INDEX_ZERO])
		end, _ := strconv.Atoi(rng[internal.INDEX_ONE])
		ranges[idx] = [2]int{int(start), int(end)}
	}

	for _, val := range ingredients {
		num, _ := strconv.Atoi(val)
		if internal.SearchInRange(ranges, int(num)) {
			counter++
		}
	}
	return counter
}

func DayFivePartTwo(input string) int {
	raw_input := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n\n")
	ranges_raw := strings.Split(raw_input[internal.INDEX_ZERO], "\n")
	counter := 0

	// Convert the ranges from string to [2]int
	ranges := make([][2]int, len(ranges_raw))
	for idx, val := range ranges_raw {
		rng := strings.Split(val, "-")
		start, _ := strconv.Atoi(rng[internal.INDEX_ZERO])
		end, _ := strconv.Atoi(rng[internal.INDEX_ONE])
		ranges[idx] = [2]int{int(start), int(end)}
	}
	merged_range := internal.MergeRangesSorted(ranges)

	for _, val := range merged_range {
		counter += (val[internal.INDEX_ONE] - val[internal.INDEX_ZERO] + internal.INDEX_ONE)
	}
	return counter
}
