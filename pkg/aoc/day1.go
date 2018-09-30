package aoc

import (
	"fmt"
	"strconv"
)

type Day1 Day

func (d Day1) FirstPart() int {
	var sum int
	firstNum, _ := strconv.Atoi(string(d.Input[0]))
	lastNum, _ := strconv.Atoi(string(d.Input[len(d.Input)-1]))
	fmt.Println("firstNum:", firstNum)
	fmt.Println("lastNum:", lastNum)

	if firstNum == lastNum {
		sum += firstNum
	}

	for i := 1; i < len(d.Input); i++ {
		currentNumber, _ := strconv.Atoi(string(d.Input[i]))
		prevNumber, _ := strconv.Atoi(string(d.Input[i-1]))

		if prevNumber == currentNumber {
			sum += prevNumber
		}

	}
	fmt.Println("sum:", sum)

	return sum
}

func (d Day1) SecondPart() int {
	var sum int
	inputLen := len(d.Input)
	step := (inputLen / 2)

	for i := 0; i < (inputLen / 2); i++ {
		currentNumber, _ := strconv.Atoi(string(d.Input[i]))
		nextNumber, _ := strconv.Atoi(string(d.Input[i+step]))

		if nextNumber == currentNumber {
			sum += (currentNumber + nextNumber)
		}
	}

	return sum
}
