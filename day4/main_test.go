package main

import "testing"

func TestPart1(t *testing.T) {
	pairList := []Pair{
		{Range{3, 9}, Range{4, 12}},
		{Range{4, 12}, Range{3, 9}},
		{Range{2, 8}, Range{3, 7}},
		{Range{3, 7}, Range{2, 8}},
	}

	got := part2(pairList)
	want := 4

	if got != want {
		t.Error(got, want)
	}
}
