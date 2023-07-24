package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var beatMap = map[string]string{
	"rock":    "scissor",
	"paper":   "rock",
	"scissor": "paper",
}

var pointsMap = map[string]int{
	"rock":    1,
	"paper":   2,
	"scissor": 3,
}

var opMap = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissor",
}

var myMap = map[string]string{
	"X": "rock",
	"Y": "paper",
	"Z": "scissor",
}

func pt1() int {
	f, err := os.Open("./inp.txt")
	if err != nil {
		log.Fatal("Could not find input file with: ", err.Error())
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	totalPoints := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		round := strings.Split(lineText, " ")
		opPlay, myPlay := round[0], round[1]
		opPlay = opMap[opPlay]
		myPlay = myMap[myPlay]

		pts := calculatePoints(opPlay, myPlay)
		totalPoints += pts
	}

	return totalPoints

}

func calculatePoints(opPlay string, myPlay string) int {
	gesturePts := pointsMap[myPlay]

	if opPlay == myPlay {
		return 3 + gesturePts
	}
	if beatMap[opPlay] == myPlay {
		return gesturePts
	}
	// only scenario remaining is you won
	return 6 + gesturePts
}