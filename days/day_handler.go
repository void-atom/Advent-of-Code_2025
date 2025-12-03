package days

type Day struct {
	PartOne func(string) int
	PartTwo func(string) int
}

var Days_map = map[int]Day{
	1: {Day_one_part_one, Day_one_part_two},
	2: {Day_two_part_one, Day_two_part_two},
}
