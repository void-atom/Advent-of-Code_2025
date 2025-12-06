package days

import (
	"aoc2025/internal"
	"math"
	"strconv"
	"strings"
)

// Formula: invalid = r*10^(len(r))+r
func DayTwoPartOne(input string) int {
	input_data := strings.Split(input, ",")
	INDEX_ZERO := 0
	INDEX_ONE := 1
	var invalid_numbers []int

	for _, value := range input_data {
		range_value := strings.Split(value, "-")
		r_start, _ := strconv.Atoi(range_value[INDEX_ZERO])
		r_end, _ := strconv.Atoi(range_value[INDEX_ONE])
		loop_idx_start := 0
		loop_idx_end := -1

		// fmt.Println("Line:", line, r_start, r_end)

		loop_idx_start, _ = strconv.Atoi(range_value[INDEX_ZERO][0 : len(range_value[INDEX_ZERO])/2])
		if len(range_value[INDEX_ONE])%2 == 0 {
			loop_idx_end, _ = strconv.Atoi(range_value[INDEX_ONE][0 : len(range_value[INDEX_ONE])/2])
		} else {
			loop_idx_end, _ = strconv.Atoi(range_value[INDEX_ONE][0 : len(range_value[INDEX_ONE])/2+1])
		}

		// fmt.Println("Indexes:", loop_idx_start, loop_idx_end)

		for i := loop_idx_start; i <= loop_idx_end; i++ {
			probable_invalid := i*int(math.Pow(10, float64(len(strconv.Itoa(i))))) + i
			// fmt.Println("Probable valid:", i, (len(strconv.Itoa(i))), i*int(math.Pow(10, float64(len(strconv.Itoa(i))))), probable_invalid)

			if probable_invalid >= r_start && probable_invalid <= r_end {
				invalid_numbers = append(invalid_numbers, probable_invalid)
				// fmt.Println(10^(len(strconv.Itoa(i))), probable_invalid)

			}
		}
	}

	return internal.SumIntSlice(invalid_numbers)
}

func DayTwoPartTwo(input string) int {
	input_data := strings.Split(input, ",")
	answer := 0

	INDEX_ZERO := 0
	INDEX_ONE := 1

	// input_data = []string{"824824821-824824827"}
	for _, num_range := range input_data {
		scan_range := strings.Split(num_range, "-")

		start, _ := strconv.Atoi(scan_range[INDEX_ZERO])
		end, _ := strconv.Atoi(scan_range[INDEX_ONE])

		for i := start; i <= end; i++ {
			to_sum := internal.RepeatedSubstring(strconv.Itoa(i))
			if to_sum {
				// fmt.Println(i)
				answer += i
			}

		}

	}
	return answer
}
