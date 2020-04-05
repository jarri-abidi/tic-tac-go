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

func TestCheckWinner(t *testing.T) {
	testTable := []testCase{
		{"O win bottom to top", [][]square{
			{square{state: o}, square{state: x}, square{state: x}},
			{square{state: x}, square{state: o}, square{state: x}},
			{square{state: ""}, square{state: ""}, square{state: o}},
		}, true, o, square{state: o}, square{state: o}},
		{"O win top to bottom", [][]square{
			{square{state: x}, square{state: x}, square{state: o}},
			{square{state: ""}, square{state: o}, square{state: x}},
			{square{state: o}, square{state: ""}, square{state: ""}},
		}, true, o, square{state: o}, square{state: o}},
		{"X win bottom to top", [][]square{
			{square{state: x}, square{state: o}, square{state: o}},
			{square{state: o}, square{state: x}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: x}},
		}, true, x, square{state: x}, square{state: x}},
		{"X win top to bottom", [][]square{
			{square{state: ""}, square{state: o}, square{state: x}},
			{square{state: o}, square{state: x}, square{state: ""}},
			{square{state: x}, square{state: ""}, square{state: o}},
		}, true, x, square{state: x}, square{state: x}},
		{"O win on top", [][]square{
			{square{state: o}, square{state: ""}, square{state: o}},
			{square{state: x}, square{state: x}, square{state: o}},
			{square{state: ""}, square{state: x}, square{state: o}},
		}, true, o, square{state: o}, square{state: o}},
		{"O win in middle", [][]square{
			{square{state: x}, square{state: o}, square{state: ""}},
			{square{state: x}, square{state: o}, square{state: x}},
			{square{state: ""}, square{state: o}, square{state: ""}},
		}, true, o, square{state: o}, square{state: o}},
		{"O win on bottom", [][]square{
			{square{state: o}, square{state: x}, square{state: ""}},
			{square{state: o}, square{state: ""}, square{state: ""}},
			{square{state: o}, square{state: x}, square{state: x}},
		}, true, o, square{state: o}, square{state: o}},
		{"X win on top", [][]square{
			{square{state: o}, square{state: o}, square{state: x}},
			{square{state: o}, square{state: ""}, square{state: x}},
			{square{state: ""}, square{state: ""}, square{state: x}},
		}, true, x, square{state: x}, square{state: x}},
		{"X win in middle", [][]square{
			{square{state: ""}, square{state: x}, square{state: ""}},
			{square{state: ""}, square{state: x}, square{state: o}},
			{square{state: ""}, square{state: x}, square{state: o}},
		}, true, x, square{state: x}, square{state: x}},
		{"X win on bottom", [][]square{
			{square{state: x}, square{state: o}, square{state: ""}},
			{square{state: x}, square{state: ""}, square{state: ""}},
			{square{state: x}, square{state: o}, square{state: o}},
		}, true, x, square{state: x}, square{state: x}},
		{"O win on right", [][]square{
			{square{state: o}, square{state: o}, square{state: o}},
			{square{state: x}, square{state: x}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, true, o, square{state: o}, square{state: o}},
		{"O win in middle", [][]square{
			{square{state: x}, square{state: x}, square{state: ""}},
			{square{state: o}, square{state: o}, square{state: o}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, true, o, square{state: o}, square{state: o}},
		{"O win on left", [][]square{
			{square{state: x}, square{state: x}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: o}, square{state: o}, square{state: o}},
		}, true, o, square{state: o}, square{state: o}},
		{"X win on right", [][]square{
			{square{state: x}, square{state: x}, square{state: x}},
			{square{state: o}, square{state: ""}, square{state: o}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, true, x, square{state: x}, square{state: x}},
		{"X win in middle", [][]square{
			{square{state: o}, square{state: o}, square{state: ""}},
			{square{state: x}, square{state: x}, square{state: x}},
			{square{state: ""}, square{state: ""}, square{state: o}},
		}, true, x, square{state: x}, square{state: x}},
		{"X win on left", [][]square{
			{square{state: o}, square{state: o}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: x}, square{state: x}, square{state: x}},
		}, true, x, square{state: x}, square{state: x}},
		{"No win on empty", [][]square{
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, false, "", square{state: ""}, square{state: ""}},
		{"No win on tie", [][]square{
			{square{state: o}, square{state: o}, square{state: x}},
			{square{state: x}, square{state: x}, square{state: o}},
			{square{state: o}, square{state: x}, square{state: o}},
		}, false, "", square{state: ""}, square{state: ""}},
	}
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			result, winner, start, end := checkWinner(testCase.Board)
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

func TestCheckVerticalWin(t *testing.T) {
	testTable := []testCase{
		{"O win on right", [][]square{
			{square{state: o}, square{state: o}, square{state: o}},
			{square{state: x}, square{state: x}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, true, o, square{state: o}, square{state: o}},
		{"O win in middle", [][]square{
			{square{state: x}, square{state: x}, square{state: ""}},
			{square{state: o}, square{state: o}, square{state: o}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, true, o, square{state: o}, square{state: o}},
		{"O win on left", [][]square{
			{square{state: x}, square{state: x}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: o}, square{state: o}, square{state: o}},
		}, true, o, square{state: o}, square{state: o}},
		{"X win on right", [][]square{
			{square{state: x}, square{state: x}, square{state: x}},
			{square{state: o}, square{state: ""}, square{state: o}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, true, x, square{state: x}, square{state: x}},
		{"X win in middle", [][]square{
			{square{state: o}, square{state: o}, square{state: ""}},
			{square{state: x}, square{state: x}, square{state: x}},
			{square{state: ""}, square{state: ""}, square{state: o}},
		}, true, x, square{state: x}, square{state: x}},
		{"X win on left", [][]square{
			{square{state: o}, square{state: o}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: x}, square{state: x}, square{state: x}},
		}, true, x, square{state: x}, square{state: x}},
		{"No win on empty", [][]square{
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, false, "", square{state: ""}, square{state: ""}},
		{"No win on tie", [][]square{
			{square{state: o}, square{state: o}, square{state: x}},
			{square{state: x}, square{state: x}, square{state: o}},
			{square{state: o}, square{state: x}, square{state: o}},
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
	testTable := []testCase{
		{"O win on top", [][]square{
			{square{state: o}, square{state: ""}, square{state: o}},
			{square{state: x}, square{state: x}, square{state: o}},
			{square{state: ""}, square{state: x}, square{state: o}},
		}, true, o, square{state: o}, square{state: o}},
		{"O win in middle", [][]square{
			{square{state: x}, square{state: o}, square{state: ""}},
			{square{state: x}, square{state: o}, square{state: x}},
			{square{state: ""}, square{state: o}, square{state: ""}},
		}, true, o, square{state: o}, square{state: o}},
		{"O win on bottom", [][]square{
			{square{state: o}, square{state: x}, square{state: ""}},
			{square{state: o}, square{state: ""}, square{state: ""}},
			{square{state: o}, square{state: x}, square{state: x}},
		}, true, o, square{state: o}, square{state: o}},
		{"X win on top", [][]square{
			{square{state: o}, square{state: o}, square{state: x}},
			{square{state: o}, square{state: ""}, square{state: x}},
			{square{state: ""}, square{state: ""}, square{state: x}},
		}, true, x, square{state: x}, square{state: x}},
		{"X win in middle", [][]square{
			{square{state: ""}, square{state: x}, square{state: ""}},
			{square{state: ""}, square{state: x}, square{state: o}},
			{square{state: ""}, square{state: x}, square{state: o}},
		}, true, x, square{state: x}, square{state: x}},
		{"X win on bottom", [][]square{
			{square{state: x}, square{state: o}, square{state: ""}},
			{square{state: x}, square{state: ""}, square{state: ""}},
			{square{state: x}, square{state: o}, square{state: o}},
		}, true, x, square{state: x}, square{state: x}},
		{"No win on empty", [][]square{
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, false, "", square{state: ""}, square{state: ""}},
		{"No win on tie", [][]square{
			{square{state: o}, square{state: o}, square{state: x}},
			{square{state: x}, square{state: x}, square{state: o}},
			{square{state: o}, square{state: x}, square{state: o}},
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

func TestCheckDiagonalWin(t *testing.T) {
	testTable := []testCase{
		{"O win bottom to top", [][]square{
			{square{state: o}, square{state: x}, square{state: x}},
			{square{state: x}, square{state: o}, square{state: x}},
			{square{state: ""}, square{state: ""}, square{state: o}},
		}, true, o, square{state: o}, square{state: o}},
		{"O win top to bottom", [][]square{
			{square{state: x}, square{state: x}, square{state: o}},
			{square{state: ""}, square{state: o}, square{state: x}},
			{square{state: o}, square{state: ""}, square{state: ""}},
		}, true, o, square{state: o}, square{state: o}},
		{"No win on O vertical win", [][]square{
			{square{state: x}, square{state: x}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: o}, square{state: o}, square{state: o}},
		}, false, "", square{state: ""}, square{state: ""}},
		{"X win bottom to top", [][]square{
			{square{state: x}, square{state: o}, square{state: o}},
			{square{state: o}, square{state: x}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: x}},
		}, true, x, square{state: x}, square{state: x}},
		{"X win top to bottom", [][]square{
			{square{state: ""}, square{state: o}, square{state: x}},
			{square{state: o}, square{state: x}, square{state: ""}},
			{square{state: x}, square{state: ""}, square{state: o}},
		}, true, x, square{state: x}, square{state: x}},
		{"No win on X vertical win", [][]square{
			{square{state: o}, square{state: o}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: x}, square{state: x}, square{state: x}},
		}, false, "", square{state: ""}, square{state: ""}},
		{"No win on empty", [][]square{
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
			{square{state: ""}, square{state: ""}, square{state: ""}},
		}, false, "", square{state: ""}, square{state: ""}},
		{"No win on tie", [][]square{
			{square{state: o}, square{state: o}, square{state: x}},
			{square{state: x}, square{state: x}, square{state: o}},
			{square{state: o}, square{state: x}, square{state: o}},
		}, false, "", square{state: ""}, square{state: ""}},
	}
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			result, winner, start, end := checkDiagonalWin(testCase.Board)
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
