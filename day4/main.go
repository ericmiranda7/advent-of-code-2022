package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair [2]Range

func (p Pair) String() string {
	return fmt.Sprintf("%v, %v", p[0], p[1])
}

type Range [2]int

func (r Range) String() string {
	return fmt.Sprintf("[%d - %d]", r[0], r[1])
}

func main() {
	/*
		2. check if fully contained
		3. return count
	*/
	pairList := *readFile("./inp.txt")

	c := part1(pairList)
	println(c)

	c = part2(pairList)
	print(c)
}

func part1(pairList []Pair) int {
	c := 0
	for _, v := range pairList {
		if isFullyContained(v[0], v[1]) {
			c += 1
		}
	}
	return c
}

func part2(pairList []Pair) int {
	c := 0
	for _, v := range pairList {
		if isPartlyContained(v[0], v[1]) || isFullyContained(v[0], v[1]) {
			c += 1
		}
	}
	return c
}

func isPartlyContained(r1, r2 Range) bool {
	if (r1[0] >= r2[0] && r1[0] <= r2[1]) || (r1[1] >= r2[0] && r1[1] <= r2[1]) {
		return true
	}
	return false
}

func isFullyContained(r1, r2 Range) bool {
	if (r1[0] >= r2[0] && r1[1] <= r2[1]) || r2[0] >= r1[0] && r2[1] <= r1[1] {
		return true
	}
	return false
}

func readFile(fileName string) *[]Pair {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("no such file: ", err.Error())
	}

	scanner := bufio.NewScanner(f)

	pairList := &[]Pair{}

	for scanner.Scan() {
		str := scanner.Text()
		strRange := strings.Split(str, ",")

		rangeList := []Range{}
		for _, v := range strRange {
			l, _ := strconv.Atoi(strings.Split(v, "-")[0])
			u, _ := strconv.Atoi(strings.Split(v, "-")[1])
			rangeList = append(rangeList, Range{l, u})
		}

		*pairList = append(*pairList, Pair(rangeList))
	}

	return pairList
}
