package main

// Advent Of Code 2022: Day02
//
// https://adventofcode.com/
// https://adventofcode.com/2022/day/2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

const (
    Paper    int8 = 0
    Scissors int8 = 1
    Rock     int8 = 2

    Lose int8 = 0
    Draw int8 = 3
    Win  int8 = 6

    ScoreRock     int8 = 1
    ScorePaper    int8 = 2
    ScoreScissors int8 = 3

    X int8 = Rock
    Y int8 = Paper
    Z int8 = Scissors
)

type Evaluator interface {
    evaluate(game string, score *int32)
}

type GameEvaluator struct {
    score     int32
    evaluator Evaluator
}

type GameEvaluatorPartOne struct{}

type GameEvaluatorPartTwo struct{}

func main() {
    fmt.Println("Advent Of Code 2022: Day02")
    start := time.Now()
    var dataFilePath = "day02/data/elves.that.want.to.play.games"
    executePart1(dataFilePath)
    executePart2(dataFilePath)
    elapsed := time.Since(start)
    fmt.Printf("Time elapsed %s\n", elapsed)
    fmt.Println("End of AOC Day02")
}

func executePart1(filePath string) {
    var evaluator = GameEvaluator{score: 0, evaluator: GameEvaluatorPartOne{}}
    executePart(filePath, evaluator.evaluator, &evaluator.score)
    fmt.Printf("End of script (1), sum of scores: %d\n", evaluator.score)
}

func executePart2(filePath string) {
    var evaluator = GameEvaluator{score: 0, evaluator: GameEvaluatorPartTwo{}}
    executePart(filePath, evaluator.evaluator, &evaluator.score)
    fmt.Printf("End of script (1), sum of scores: %d\n", evaluator.score)
}

func executePart(filePath string, evaluator Evaluator, score *int32) {
    file := openFileReader(filePath)
    scanAndCompileScoreOfGames(file, score, evaluator)
    closeFile(file)
}

func openFileReader(filePath string) *os.File {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    return file
}

func closeFile(file *os.File) {
    defer file.Close()
}

func scanAndCompileScoreOfGames(reader io.Reader, score *int32, evaluator Evaluator) {
    scanner := bufio.NewScanner(reader)
    for scanner.Scan() {
        gameText := scanner.Text()
        evaluator.evaluate(gameText, score)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func (GameEvaluatorPartOne) evaluate(game string, score *int32) {
    firstHand, secondHand := parseGameHands(game)
    *score += int32(getScoreFromHand(secondHand) + getScoreOfGame(secondHand, firstHand))
}

func (GameEvaluatorPartTwo) evaluate(game string, score *int32) {
    firstHand, secondHand := parseGameHands(game)
    secondHand = getHandFromEndGameIndicator(secondHand, firstHand)
    *score += int32(getScoreFromHand(secondHand) + getScoreOfGame(secondHand, firstHand))
}

func parseGameHands(game string) (int8, int8) {
    var hands = strings.Split(game, " ")
    var firstHand = getHandFromStr(hands[0])
    var secondHand = getHandFromStr(hands[1])
    return firstHand, secondHand
}

func getHandFromStr(str string) int8 {
    var hand int8
    switch str {
    case "X", "A":
        hand = Rock
    case "Y", "B":
        hand = Paper
    case "Z", "C":
        hand = Scissors
    }
    return hand
}

func getHandFromEndGameIndicator(endGameIndicator int8, handAgainst int8) int8 {
    var handForIndicatedEndGame int8
    switch endGameIndicator {
    case X: // Should lose
        handForIndicatedEndGame = ((handAgainst + 2) % 3)
    case Y: // Should draw
        handForIndicatedEndGame = handAgainst
    case Z: // Should win
        handForIndicatedEndGame = (handAgainst + 1) % 3
    }
    return handForIndicatedEndGame
}

func getScoreFromHand(hand int8) int8 {
    var score int8
    switch hand {
    case Rock:
        score = ScoreRock
    case Scissors:
        score = ScoreScissors
    case Paper:
        score = ScorePaper
    }
    return score
}

func getScoreOfGame(firstHand int8, otherHand int8) int8 {
    difference := firstHand - otherHand
    var score int8
    switch math.Abs(float64(difference)) {
    case 0:
        score = Draw
    case 1:
        if difference > 0 {
            score = Win
        } else {
            score = Lose
        }
    case 2:
        if difference > 0 {
            score = Lose
        } else {
            score = Win
        }
    }
    return score
}
