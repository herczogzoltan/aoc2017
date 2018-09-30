package aoc

type Assignments interface {
	FirstPart() int
	SecondPart() int
}

type Day struct {
	Input string
}
