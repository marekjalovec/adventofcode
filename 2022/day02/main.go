package main

import (
	"adventofcode/utils"
	"fmt"
)

const RockScore = 1
const PaperScore = 2
const ScissorsScore = 3

const LostScore = 0
const DrawScore = 3
const WonScore = 6

// FirstStarScenarios
// A for Rock, B for Paper, and C for Scissors
// X for Rock, Y for Paper, and Z for Scissors
var FirstStarScenarios = map[string]int{
	"A Z": ScissorsScore + LostScore,
	"B X": RockScore + LostScore,
	"C Y": PaperScore + LostScore,

	"A X": RockScore + DrawScore,
	"B Y": PaperScore + DrawScore,
	"C Z": ScissorsScore + DrawScore,

	"A Y": PaperScore + WonScore,
	"B Z": ScissorsScore + WonScore,
	"C X": RockScore + WonScore,
}

// SecondStarScenarios
// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
var SecondStarScenarios = map[string]int{
	"A X": ScissorsScore + LostScore,
	"A Y": RockScore + DrawScore,
	"A Z": PaperScore + WonScore,

	"B X": RockScore + LostScore,
	"B Y": PaperScore + DrawScore,
	"B Z": ScissorsScore + WonScore,

	"C X": PaperScore + LostScore,
	"C Y": ScissorsScore + DrawScore,
	"C Z": RockScore + WonScore,
}

func firstStar(lines []string) int {
	var score int
	for _, v := range lines {
		score += FirstStarScenarios[v]
	}
	return score
}

func secondStar(lines []string) int {
	var score int
	for _, v := range lines {
		score += SecondStarScenarios[v]
	}
	return score
}

func main() {
	var lines = utils.ReadFile("2022/day02/input.txt")

	fmt.Println("First Star:", firstStar(lines))
	fmt.Println("Second Star:", secondStar(lines))
}
