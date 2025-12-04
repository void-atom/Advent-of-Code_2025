package days

type Day struct {
	PartOne func(string) int
	PartTwo func(string) int
}

var DaysMap = map[int]Day{
	1: {DayOnePartOne, DayOnePartTwo},
	2: {DayTwoPartOne, DayTwoPartTwo},
	3: {DayThreePartOne, DayThreePartTwo},
}
