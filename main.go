package main

import (
	"fmt"
	"github.com/JarriAbidi/tictacgo/draw"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
)

const (
	O            = "O"
	X            = "X"
	boardLength  = float64(600)
	squareLength = boardLength / 3
	thickness    = boardLength / 60
	crossLength  = squareLength / 2
	circleRadius = squareLength * 3 / 10
	lineOffset   = squareLength * 2 / 5
	firstOffset  = squareLength / 2
	secondOffset = squareLength * 3 / 2
	thirdOffset  = squareLength * 5 / 2
)

var (
	zv          = pixel.ZV
	zs          = square{center: zv, corner: zv, state: ""}
	squareColor = colornames.Black
	circleColor = colornames.Red
	crossColor  = colornames.Blue
	over        bool
	turn        string
	winner      string
)

type vector = pixel.Vec

type square struct {
	center vector
	corner vector
	state  string
}

func (s *square) string() string {
	return fmt.Sprintf("center: %s, corner: %s, state: %s", s.center.String(), s.corner.String(), s.state)
}

func v(x, y float64) vector {
	return pixel.V(x, y)
}

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Tic Tac Go",
		Bounds: pixel.R(0, 0, boardLength, boardLength),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	board := newGame(imd, O)

	for !win.Closed() {
		win.Clear(colornames.White)
		imd.Draw(win)
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			processClick(imd, board, getNearestSquare(board, win.MousePosition()))
		}
		win.Update()
	}
}

func newGame(imd *imdraw.IMDraw, firstTurn string) [][]square {
	go draw.Board(imd, squareLength, thickness, squareColor)
	turn = firstTurn
	over = false
	winner = ""
	board := [][]square{
		{
			{center: v(firstOffset, firstOffset), corner: v(-lineOffset, -lineOffset)},
			{center: v(firstOffset, secondOffset), corner: v(-lineOffset, 0)},
			{center: v(firstOffset, thirdOffset), corner: v(-lineOffset, lineOffset)},
		}, {
			{center: v(secondOffset, firstOffset), corner: v(0, -lineOffset)},
			{center: v(secondOffset, secondOffset), corner: v(0, 0)},
			{center: v(secondOffset, thirdOffset), corner: v(0, lineOffset)},
		}, {
			{center: v(thirdOffset, firstOffset), corner: v(lineOffset, -lineOffset)},
			{center: v(thirdOffset, secondOffset), corner: v(lineOffset, 0)},
			{center: v(thirdOffset, thirdOffset), corner: v(lineOffset, lineOffset)},
		},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j].state = ""
		}
	}
	return board
}

func processClick(imd *imdraw.IMDraw, board [][]square, clickedSquare *square) {
	var s1, s2 square
	if !over && clickedSquare.state == "" {
		if turn == O {
			draw.O(imd, clickedSquare.center, circleRadius, thickness, circleColor)
			clickedSquare.state = O
			turn = X
		} else if turn == X {
			draw.X(imd, clickedSquare.center, crossLength, thickness, crossColor)
			clickedSquare.state = X
			turn = O
		}
		if over, winner, s1, s2 = checkWinner(board); over {
			fmt.Printf("The winner is %s!\n", winner)
			var c color.Color 
			if winner == O {
				c = circleColor
			} else {
				c = crossColor
			}
			draw.Line(imd, s1.center, s2.center, s1.corner, s2.corner, thickness, c)
		}
	}
}

func getNearestSquare(board [][]square, click vector) *square {
	minDistance := math.MaxFloat64
	var nearest *square
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			square := &board[i][j]
			result := click.Sub(square.center)
			diff := math.Abs(result.X) + math.Abs(result.Y)
			if minDistance > diff {
				minDistance = diff
				nearest = square
			}
		}
	}
	return nearest
}

func checkWinner(board [][]square) (bool, string, square, square) {
	if over, winner, s1, s2 := checkVerticalWin(board); over {
		return true, winner, s1, s2
	}
	if over, winner, s1, s2 := checkHorizontalWin(board); over {
		return true, winner, s1, s2
	}
	if over, winner, s1, s2 := checkDiagonalWin(board); over {
		return true, winner, s1, s2
	}
	return false, "", zs, zs
}

func checkHorizontalWin(board [][]square) (bool, string, square, square) {
	i := 1
	j := 2
	for j >= 0 {
		if board[i][j].state == "" {
			j--
			continue
		}
		if board[i][j].state == O &&
			board[i-1][j].state == O &&
			board[i+1][j].state == O {
			return true, O, board[0][j], board[2][j]
		}
		if board[i][j].state == X &&
			board[i-1][j].state == X &&
			board[i+1][j].state == X {
			return true, X, board[0][j], board[2][j]
		}
		j--
	}
	return false, "", zs, zs
}

func checkVerticalWin(board [][]square) (bool, string, square, square) {
	i := 2
	j := 1
	for i >= 0 {
		if board[i][j].state == "" {
			i--
			continue
		}
		if board[i][j].state == O &&
			board[i][j-1].state == O &&
			board[i][j+1].state == O {
			return true, O, board[i][0], board[i][2]
		}
		if board[i][j].state == X &&
			board[i][j-1].state == X &&
			board[i][j+1].state == X {
			return true, X, board[i][0], board[i][2]
		}
		i--
	}
	return false, "", zs, zs
}

func checkDiagonalWin(board [][]square) (bool, string, square, square) {
	i, j := 1, 1
	if board[i][j].state == "" {
		return false, "", zs, zs
	}
	if board[i][j].state == O &&
		board[i-1][j-1].state == O &&
		board[i+1][j+1].state == O {
		return true, O, board[i-1][j-1], board[i+1][j+1]
	}
	if board[i][j].state == O &&
		board[i-1][j+1].state == O &&
		board[i+1][j-1].state == O {
		return true, O, board[i-1][j+1], board[i+1][j-1]
	}
	if board[i][j].state == X &&
		board[i-1][j-1].state == X &&
		board[i+1][j+1].state == X {
		return true, X, board[i-1][j-1], board[i+1][j+1]
	}
	if board[i][j].state == X &&
		board[i-1][j+1].state == X &&
		board[i+1][j-1].state == X {
		return true, X, board[i-1][j+1], board[i+1][j-1]
	}
	return false, "", zs, zs
}
