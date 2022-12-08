package main

import "testing"

var sample [][]int = [][]int{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func TestFindScenicSpot(t *testing.T) {
	expected := 8
	actual := findScenicSpot(sample)
	if actual != expected {
		t.Errorf("%d is supposed to be %d\n", actual, expected)
	}
}

func TestFindTrees(t *testing.T) {
	expected := 21
	actual := findTrees(sample)
	if actual != expected {
		t.Errorf("%d is supposed to be %d\n", actual, expected)
	}
}
