package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"fmt"
)

const (
	stateO string = "O"
	stateX string = "X"
	stateEmpty string = ""
)

var (
	zv = pixel.ZV
	squareLength = float64(170)
	crossLength = float64(90)
	circleRadius = float64(50)
	drawingThickness = float64(10)
	squareColor = colornames.Black
	circleColor = colornames.Red
	crossColor = colornames.Blue
	lineOffset = squareLength/3
	first = squareLength/2
	second = (squareLength/2)*3
	third = (squareLength/2)*5
	coordinates = [][]vector {
		{ v(first, first), v(first, second), v(first, third) },
		{ v(second, first), v(second, second), v(second, third) },
		{ v(third, first), v(third, second), v(third, third) },
	}
	offsets = [][]vector {
		{ v(-lineOffset, -lineOffset), v(-lineOffset, 0), v(-lineOffset, lineOffset) },
		{ v(0, -lineOffset), v(0, 0), v(0, lineOffset) },
		{ v(lineOffset, -lineOffset), v(lineOffset, 0), v(lineOffset, lineOffset) },
	}
	state = make(map[vector]string)
)

type vector = pixel.Vec 

func v(x, y float64) vector {
	return pixel.V(x, y)
}

func main() {
	pixelgl.Run(run)
}

func run() {
	// all of our code will be fired up from here
	cfg := pixelgl.WindowConfig{
		Title:  "Tic Tac Go",
		Bounds: pixel.R(0, 0, squareLength*3, squareLength*3),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	drawBoard(imd, squareLength, drawingThickness, squareColor)

	isOTurn := true
	gameOver := false
	winner := ""
	v1 := zv
	v2 := zv
	for !win.Closed() {
		win.Clear(colornames.White)
		imd.Draw(win)
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			drawPosition := getNearestSquare(win.MousePosition())
			if !gameOver && state[drawPosition] == stateEmpty {
				if isOTurn {
					drawO(imd, drawPosition, circleRadius, drawingThickness, circleColor)
					state[drawPosition] = stateO
					isOTurn = false
				} else {
					drawX(imd, drawPosition, crossLength, drawingThickness, crossColor)
					state[drawPosition] = stateX
					isOTurn = true
				}
				if gameOver, winner, v1, v2 = checkWinner(); gameOver {
					fmt.Printf("The winner is %s!\n", winner)
					x1 := int(v1.X)
					x2 := int(v2.X)
					y1 := int(v1.Y)
					y2 := int(v2.Y)
					var lineColor color.Color
					if isOTurn {
						lineColor = crossColor
					} else {
						lineColor = circleColor
					} 
					drawLine(imd, coordinates[x1][y1], coordinates[x2][y2], offsets[x1][y1], offsets[x2][y2], drawingThickness, lineColor)
				} 
			}
		}
		win.Update()
	}
}

func drawBoard(imd *imdraw.IMDraw, length, thickness float64, lineColor color.Color) {
	imd.Color = lineColor
	drawSquare(imd, length, 0, 0, thickness)
	drawSquare(imd, length, 0, length, thickness)
	drawSquare(imd, length, 0, 2*length, thickness)
	drawSquare(imd, length, length, 0, thickness)
	drawSquare(imd, length, length, length, thickness)
	drawSquare(imd, length, length, 2*length, thickness)
	drawSquare(imd, length, 2*length, 0, thickness)
	drawSquare(imd, length, 2*length, length, thickness)
	drawSquare(imd, length, 2*length, 2*length, thickness)
}

func drawSquare(imd *imdraw.IMDraw, length, x, y, thickness float64) {
	imd.Push(v(x, y))
	imd.Push(v(x, y+length))
	imd.Push(v(x+length, y))
	imd.Push(v(x+length, y+length))
	imd.Rectangle(thickness)
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			state[coordinates[i][j]] = stateEmpty
		}
	}
}

func drawO(imd *imdraw.IMDraw, c vector, radius, thickness float64, color color.Color) {
	imd.Color = color
	imd.Push(c)
	imd.Circle(radius, thickness)
}

func drawX(imd *imdraw.IMDraw, c vector, length, thickness float64, color color.Color) {
	imd.Color = color
	l := length/2
	imd.Push(v(c.X-l, c.Y-l), v(c.X+l, c.Y+l))
	imd.Line(thickness)
	imd.Push(v(c.X+l, c.Y-l), v(c.X-l, c.Y+l))
	imd.Line(thickness)
}

func drawLine(imd *imdraw.IMDraw, c1, c2, o1, o2 vector, thickness float64, color color.Color) {
	imd.Color = color
	if c1.X == c2.X {
		imd.Push(c1.Add(v(0, o1.Y)), c2.Add(v(0, o2.Y)))
	} else if c1.Y == c2.Y {
		imd.Push(c1.Add(v(o1.X, 0)), c2.Add(v(o2.X, 0)))
	} else {
		imd.Push(c1.Add(o1), c2.Add(o2))
	}
	imd.Line(thickness)
}

func getNearestSquare(click vector) vector {
	minDistance := math.MaxFloat64
	var nearestX float64
	var nearestY float64
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			square := coordinates[i][j] 
			result := click.Sub(square)
			diff := math.Abs(result.X) + math.Abs(result.Y)
			if minDistance > diff {
				minDistance = diff
				nearestX = square.X
				nearestY = square.Y
			}
		}
	}
	return v(nearestX, nearestY)
}

func checkWinner() (bool, string, vector, vector) {
	// check horizontal win
	if gameOver, winner := checkHorizontalWin(1, 2); gameOver {	return true, winner, v(0,2), v(2,2) }
	if gameOver, winner := checkHorizontalWin(1, 1); gameOver { return true, winner, v(0,1), v(2,1) }
	if gameOver, winner := checkHorizontalWin(1, 0); gameOver { return true, winner, v(0,0), v(2,0) }
	// check vertical win
	if gameOver, winner := checkVerticalWin(0, 1); gameOver { return true, winner, v(0,0), v(0,2) }
	if gameOver, winner := checkVerticalWin(1, 1); gameOver { return true, winner, v(1,0), v(1,2) }
	if gameOver, winner := checkVerticalWin(2, 1); gameOver { return true, winner, v(2,0), v(2,2) }
	// check diagonal win
	if gameOver, winner, v1, v2 := checkDiagonalWin(1, 1); gameOver { return true, winner, v1, v2 }
	return false, "", zv, zv
}

func checkHorizontalWin(i, j int) (bool, string) {
	if state[coordinates[i][j]] == stateEmpty {
		return false, ""
	} 
	if state[coordinates[i][j]] == stateO && 
	state[coordinates[i-1][j]] == stateO &&
	state[coordinates[i+1][j]] == stateO {
		return true, "O"
	} 
	if state[coordinates[i][j]] == stateX && 
	state[coordinates[i-1][j]] == stateX &&
	state[coordinates[i+1][j]] == stateX {
		return true, "X"
	}
	return false, ""
}

func checkVerticalWin(i, j int) (bool, string) {
	if state[coordinates[i][j]] == stateEmpty {
		return false, ""
	} 
	if state[coordinates[i][j]] == stateO && 
	state[coordinates[i][j-1]] == stateO &&
	state[coordinates[i][j+1]] == stateO {
		return true, "O"
	} 
	if state[coordinates[i][j]] == stateX && 
	state[coordinates[i][j-1]] == stateX &&
	state[coordinates[i][j+1]] == stateX {
		return true, "X"
	}
	return false, ""
}

func checkDiagonalWin(i, j int) (bool, string, vector, vector) {
	if state[coordinates[i][j]] == stateEmpty {
		return false, "", zv, zv
	} 
	if state[coordinates[i][j]] == stateO && 
	state[coordinates[i-1][j-1]] == stateO &&
	state[coordinates[i+1][j+1]] == stateO {
		return true, "O", v(float64(i-1), float64(j-1)), v(float64(i+1), float64(j+1))
	} 
	if state[coordinates[i][j]] == stateO && 
	state[coordinates[i-1][j+1]] == stateO &&
	state[coordinates[i+1][j-1]] == stateO {
		return true, "O", v(float64(i-1), float64(j+1)), v(float64(i+1), float64(j-1))
	}
	if state[coordinates[i][j]] == stateX && 
	state[coordinates[i-1][j-1]] == stateX &&
	state[coordinates[i+1][j+1]] == stateX {
		return true, "X", v(float64(i-1), float64(j-1)), v(float64(i+1), float64(j+1))
	} 
	if state[coordinates[i][j]] == stateX && 
	state[coordinates[i-1][j+1]] == stateX &&
	state[coordinates[i+1][j-1]] == stateX {
		return true, "X", v(float64(i-1), float64(j+1)), v(float64(i+1), float64(j-1))
	}
	return false, "", zv, zv
}