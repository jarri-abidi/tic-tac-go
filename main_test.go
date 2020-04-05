package main

import (
	"testing"
)

func TestCheckWinner(t *testing.T) {
	testTable := []struct {
		Name       string
		Board      [][]square
		Result     bool
		Winner     string
		Start, End square
	}{
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

func TestGetNearestSquare(t *testing.T) {
	fo := firstOffset
	so := secondOffset
	to := thirdOffset
	board := [][]square{
		{
			{center: v(fo, fo)},
			{center: v(fo, so)},
			{center: v(fo, to)},
		}, {
			{center: v(so, fo)},
			{center: v(so, so)},
			{center: v(so, to)},
		}, {
			{center: v(to, fo)},
			{center: v(to, so)},
			{center: v(to, to)},
		},
	}
	testTable := []struct {
		Name   string
		Board  [][]square
		Click  vector
		Result square
	}{
		{"Click bottom left corner", board, v(0, 0), square{center: v(fo, fo)}},
		{"Click left corner", board, v(0, so), square{center: v(fo, so)}},
		{"Click top left corner", board, v(0, so*2), square{center: v(fo, to)}},
		{"Click bottom middle corner", board, v(so, 0), square{center: v(so, fo)}},
		{"Click center", board, v(so, so), square{center: v(so, so)}},
		{"Click top middle corner", board, v(so, so*2), square{center: v(so, to)}},
		{"Click bottom right corner", board, v(so*2, 0), square{center: v(to, fo)}},
		{"Click right corner", board, v(so*2, so), square{center: v(to, so)}},
		{"Click top right corner", board, v(so*2, so*2), square{center: v(to, to)}},
	}
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			result := getNearestSquare(testCase.Board, testCase.Click)
			if *result != testCase.Result {
				t.Errorf("Expected result to be %s got %s", testCase.Result.string(), result.string())
			}
		})
	}
}
