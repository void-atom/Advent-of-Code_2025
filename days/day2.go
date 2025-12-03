package days

import (
	"aoc2025/internal"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Formula: invalid = r*10^(len(r))+r
func Day_two_part_one(input string) int {
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

	return internal.Sum_Int_Slice(invalid_numbers)
}

func Day_two_part_two(input string) int {
	input_data := strings.Split(input, ",")
	fmt.Println(input_data)
	return 2
}
