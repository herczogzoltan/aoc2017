package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/herczogzoltan/aoc/pkg/aoc"
)

var FilePath string
var Day int
var Part int

func init() {
	flag.StringVar(&FilePath, "FilePath", "input.txt", "Input file Path")
	flag.IntVar(&Day, "Day", 1, "Number of the day(1 to 25)")
	flag.IntVar(&Part, "Part", 1, "Assignment part (1 or 2)")

	flag.Parse()
}

func evaluate(day int, part int, input string) int {
	switch {
	case day == 1 && part == 1:
		return aoc.Day1{
			Input: input,
		}.FirstPart()
	case day == 1 && part == 2:
		return aoc.Day1{
			Input: input,
		}.SecondPart()
	}

	return 0
}

func main() {
	inputStr, err := ioutil.ReadFile(FilePath)

	if err != nil {
		fmt.Errorf("ReadFileError: %v\n", err)
	}

	if os.Getenv("DEBUG") != "" {
		fmt.Printf("Input string: %v\n", string(inputStr))
		fmt.Printf("Day: %v\n", Day)
		fmt.Printf("Part: %v\n", Part)
	}

	fmt.Printf("Result: %v", evaluate(Day, Part, string(inputStr)))
}
