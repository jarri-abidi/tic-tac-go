package main

import (
	"testing"
)

type testCase struct {
	Name       string
	Board      [][]square
	Result     bool
	Winner     string
	Start, End square
}

func TestCheckVerticalWin(t *testing.T) {
	testTable := []testCase {
		{"O win on right", [][]square{
			{square{state: O}, square{state: O}, square{state: O}},
			{square{state: X}, square{state: X}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, true, O, square{state: O}, square{state: O}},
		{"O win in middle", [][]square{
			{square{state: X}, square{state: X}, square{state: ""}},
			{square{state: O}, square{state: O}, square{state: O}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, true, O, square{state: O}, square{state: O}},
		{"O win on left", [][]square{
			{square{state: X}, square{state: X}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: O}, square{state: O}, square{state: O}},
		}, true, O, square{state: O}, square{state: O}},
		{"X win on right", [][]square{
			{square{state: X}, square{state: X}, square{state: X}},
			{square{state: O}, square{state: ""}, square{state: O}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, true, X, square{state: X}, square{state: X}},
		{"X win in middle", [][]square{
			{square{state: O}, square{state: O}, square{state: ""}},
			{square{state: X}, square{state: X}, square{state: X}},
			{square{state: ""}, square{state: ""}, square{state: O}},
		}, true, X, square{state: X}, square{state: X}},
		{"X win on left", [][]square{
			{square{state: O}, square{state: O}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: X}, square{state: X}, square{state: X}},
		}, true, X, square{state: X}, square{state: X}},
		{"No win on empty", [][]square{
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, false, "", square{state: ""}, square{state: ""}},
	}
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			result, winner, start, end := checkVerticalWin(testCase.Board)
			if result != testCase.Result {
				t.Errorf("Expected result to be %t got %t", testCase.Result, result)
			}
			if winner != testCase.Winner {
				t.Errorf("Expected winner to be %s got %s", testCase.Winner, winner)
			}
			if start != testCase.Start {
				t.Errorf("Expected start square to be %s got %s", testCase.Start.string(), start.string())
			}
			if end != testCase.End {
				t.Errorf("Expected end square to be %s got %s", testCase.End.string(), end.string())
			}
		})
	}
}

func TestCheckHorizontalWin(t *testing.T) {
	testTable := []testCase {
		{"O win on top", [][]square{
			{square{state: O}, square{state: ""}, square{state: O}},
			{square{state: X}, square{state: X}, square{state: O}},
			{square{state: ""}, square{state: X}, square{state: O}},
		}, true, O, square{state: O}, square{state: O}},
		{"O win in middle", [][]square{
			{square{state: X}, square{state: O}, square{state: ""}},
			{square{state: X}, square{state: O}, square{state: X}},
			{square{state: ""}, square{state: O}, square{state: ""}},
		}, true, O, square{state: O}, square{state: O}},
		{"O win on bottom", [][]square{
			{square{state: O}, square{state: X}, square{state: ""}},
			{square{state: O}, square{state: ""}, square{state: ""}},
			{square{state: O}, square{state: X}, square{state: X}},
		}, true, O, square{state: O}, square{state: O}},
		{"X win on top", [][]square{
			{square{state: O}, square{state: O}, square{state: X}},
			{square{state: O}, square{state: ""}, square{state: X}},
			{square{state: ""}, square{state: ""}, square{state: X}},
		}, true, X, square{state: X}, square{state: X}},
		{"X win in middle", [][]square{
			{square{state: ""}, square{state: X}, square{state: ""}},
			{square{state: ""}, square{state: X}, square{state: O}},
			{square{state: ""}, square{state: X}, square{state: O}},
		}, true, X, square{state: X}, square{state: X}},
		{"X win on bottom", [][]square{
			{square{state: X}, square{state: O}, square{state: ""}},
			{square{state: X}, square{state: ""}, square{state: ""}},
			{square{state: X}, square{state: O}, square{state: O}},
		}, true, X, square{state: X}, square{state: X}},
		{"No win on empty", [][]square{
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, false, "", square{state: ""}, square{state: ""}},
	}
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			result, winner, start, end := checkHorizontalWin(testCase.Board)
			if result != testCase.Result {
				t.Errorf("Expected result to be %t got %t", testCase.Result, result)
			}
			if winner != testCase.Winner {
				t.Errorf("Expected winner to be %s got %s", testCase.Winner, winner)
			}
			if start != testCase.Start {
				t.Errorf("Expected start square to be %s got %s", testCase.Start.string(), start.string())
			}
			if end != testCase.End {
				t.Errorf("Expected end square to be %s got %s", testCase.End.string(), end.string())
			}
		})
	}
}
