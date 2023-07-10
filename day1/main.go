package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("./inp.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	first, second, third := 0, 0, 0
	currElfCalories := 0
	for scanner.Scan() {
		cl, err := strconv.Atoi(scanner.Text())

		if err != nil {
			// a newline, meaning new elf
			if currElfCalories > first {
				first, second, third = currElfCalories, first, second
			} else if currElfCalories > second {
				second, third = currElfCalories, second
			} else if currElfCalories > third {
				third = currElfCalories
			}
			currElfCalories = 0
			continue
		}

		currElfCalories += cl
	}

	print(first + second + third)
}
