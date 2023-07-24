package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var strategyMap = map[string]string{
	"X": "lose",
	"Y": "draw",
	"Z": "win",
}

func pt2() int {
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
		opPlay, strategy := round[0], round[1]
		opPlay = opMap[opPlay]

		pts := calculatePointsPt2(opPlay, strategy)
		totalPoints += pts
	}

	return totalPoints

}

func calculatePointsPt2(opPlay string, strategy string) int {
	if strategyMap[strategy] == "lose" {
		return 0 + pointsMap[beatMap[opPlay]]
	}
	if strategyMap[strategy] == "draw" {
		return 3 + pointsMap[opPlay]
	}
	return 6 + pointsMap[beatMap[beatMap[opPlay]]]
}
