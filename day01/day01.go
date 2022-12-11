package main

// Advent Of Code: Day01
//
// https://adventofcode.com/
// https://adventofcode.com/2022/day/1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
    fmt.Println("Advent Of Code 2022: Day01")

    start := time.Now()

    readFile, err := os.Open("day01/data/food.elves")

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    var maxCaloriesCarried int64 = 0
    var sumColoriesCarried int64 = 0

    for fileScanner.Scan() {
        fileLine := fileScanner.Text()

        if len(fileLine) > 0 {
            calories, err := strconv.ParseInt(fileLine, 10, 32)
            if err == nil {
                sumColoriesCarried += calories
        }
        } else {
            if sumColoriesCarried > maxCaloriesCarried {
                maxCaloriesCarried = sumColoriesCarried
            }
            sumColoriesCarried = 0
        }
    }
    readFile.Close()

    end := time.Now()
    elapsed := end.Sub(start)

    fmt.Println(fmt.Sprintf("Max calories carried: %d", maxCaloriesCarried))
    fmt.Printf("Time elapsed: %d ns", elapsed.Nanoseconds())
}
