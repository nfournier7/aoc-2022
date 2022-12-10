package main

// Advent Of Code 2022: Day03
//
// https://adventofcode.com/
// https://adventofcode.com/2022/day/3

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Advent Of Code 2022: Day03")
	start := time.Now()
	testData := "vJrwpWtwJgWrhcsFMMfFFhFp"
	duplicate := findDuplicateByte(testData)
	priority := getPriority(duplicate)
	fmt.Printf("duplicate: %s, priority: %d\n", string(duplicate), priority)
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed %s\n", elapsed)
	fmt.Println("End of AOC Day03")
}

type void struct{}

var Char void

func findDuplicateByte(rucksack string) rune {

	leftSectionByteSet := make(map[int32]void)
	righSectionByteSet := make(map[int32]void)

	var startIndex = 0
	var endIndex = len(rucksack) - 1

	for endIndex >= startIndex {

		leftChar := rune(rucksack[startIndex])
		rightChar := rune(rucksack[endIndex])

		leftSectionByteSet[leftChar] = Char
		leftSectionByteSet[rightChar] = Char

		fmt.Printf("left: %d, right: %d\n", leftChar, rightChar)

		_, isRightCharAlreadyInLeftSection := leftSectionByteSet[rightChar]
		_, isLeftCharAlreadyInRightSection := righSectionByteSet[leftChar]

		if endIndex == startIndex && isLeftCharAlreadyInRightSection && isRightCharAlreadyInLeftSection {
			fmt.Println(leftChar)
			return leftChar
		} else if isRightCharAlreadyInLeftSection {
			fmt.Println(rightChar)
			return rightChar
		} else if isLeftCharAlreadyInRightSection {
			fmt.Println(leftChar)
			return leftChar
		} else {
			fmt.Println("Should not reach")
		}

		startIndex += 1
		endIndex -= 1

		fmt.Printf("%d, %d\n\n", startIndex, endIndex)
		fmt.Println(leftSectionByteSet)
		fmt.Println(righSectionByteSet)
	}

	return rune('!')
}

func getPriority(character rune) int32 {
	return character - rune('a') + 1
}
