package main

import (
	"bufio"
	"os"
)

func main() {
	f, _ := os.Open("./inp.txt")
	s := bufio.NewScanner(f)

	p1prio := part1(s)
	p2prio := part2(s)

	print(p1prio)
	print(p2prio)
}

func part1(s *bufio.Scanner) int {
	totalPriority := 0
	for s.Scan() {
		rucksack := s.Text()
		itemSize := len(rucksack)
		comp1 := rucksack[:itemSize/2]
		comp2 := rucksack[itemSize/2:]
		duplicateItem := findDuplicate(comp1, comp2)
		prior := getPriority(duplicateItem)
		totalPriority += prior
	}

	return totalPriority
}

func part2(s *bufio.Scanner) int {
	totalPriority := 0

	for s.Scan() {
		r1 := s.Text()
		if !s.Scan() {
			break
		}
		r2 := s.Text()
		if !s.Scan() {
			break
		}
		r3 := s.Text()

		groupBadge := getBadge(r1, r2, r3)
		totalPriority += getPriority(groupBadge)
	}

	return totalPriority
}

func getBadge(r1, r2, r3 string) rune {
	return findIntersection(r1, r2, r3)
}

func getPriority(item rune) int {
	if int(item) <= 90 {
		return (int(item) - 65) + 27
	}
	return int(item) - 97 + 1
}

func findDuplicate(c1 string, c2 string) rune {
	c1map := map[rune]struct{}{}
	for _, item := range c1 {
		c1map[item] = struct{}{}
	}

	for _, item := range c2 {
		if _, ok := c1map[item]; ok {
			return item
		}
	}
	return 'x'
}

func findIntersection(c1, c2, c3 string) rune {
	c1map := map[rune]struct{}{}
	resSet := map[rune]struct{}{}
	for _, item := range c1 {
		c1map[item] = struct{}{}
	}
	for _, item := range c2 {
		if _, ok := c1map[item]; ok {
			resSet[item] = struct{}{}
		}
	}

	for _, item := range c3 {
		if _, ok := resSet[item]; ok {
			return item
		}
	}

	panic("out of syllabus")
}
