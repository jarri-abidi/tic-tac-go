package draw

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"image/color"
)

type vector = pixel.Vec

func v(x, y float64) vector {
	return pixel.V(x, y)
}

// Board draws a tic-tac-toe board of squares where
// length is length of sides of squares,
// thickness is thickness of lines on board
// and lineColor is color of lines on board
func Board(imd *imdraw.IMDraw, length, thickness float64, lineColor color.Color) {
	imd.Color = lineColor
	Square(imd, length, 0, 0, thickness)
	Square(imd, length, 0, length, thickness)
	Square(imd, length, 0, 2*length, thickness)
	Square(imd, length, length, 0, thickness)
	Square(imd, length, length, length, thickness)
	Square(imd, length, length, 2*length, thickness)
	Square(imd, length, 2*length, 0, thickness)
	Square(imd, length, 2*length, length, thickness)
	Square(imd, length, 2*length, 2*length, thickness)
}

// Square draws an individual square where
// length is the length of sides of square,
// thickness is the thickness of border of square,
// and x,y are the position of bottom left corner of square
func Square(imd *imdraw.IMDraw, length, x, y, thickness float64) {
	imd.Push(v(x, y))
	imd.Push(v(x, y+length))
	imd.Push(v(x+length, y))
	imd.Push(v(x+length, y+length))
	imd.Rectangle(thickness)
}

// O draws an O (circle) mark where
// c is center position of the circle
// radius is the length from center of circle to its border,
// thickness is the thickness of border of circle,
// and color is the color of the circle
func O(imd *imdraw.IMDraw, c vector, radius, thickness float64, color color.Color) {
	imd.Color = color
	imd.Push(c)
	imd.Circle(radius, thickness)
}

// X draws an X (cross) mark where
// c is center position of the cross
// length is the length of each diagonal in the cross,
// thickness is the thickness of the cross,
// and color is the color of the cross
func X(imd *imdraw.IMDraw, c vector, length, thickness float64, color color.Color) {
	imd.Color = color
	l := length / 2
	imd.Push(v(c.X-l, c.Y-l), v(c.X+l, c.Y+l))
	imd.Line(thickness)
	imd.Push(v(c.X+l, c.Y-l), v(c.X-l, c.Y+l))
	imd.Line(thickness)
}

// Line draws a line across marked squares (to show victory) where
// c1 is center position of the square at one end of the line
// c2 is center position of the square at other end of the line
// o1 is potental offset with which to extend the line at one end
// o2 is potental offset with which to extend the line at other end
// thickness is the thickness of the line,
// and color is the color of the line
func Line(imd *imdraw.IMDraw, c1, c2, o1, o2 vector, thickness float64, color color.Color) {
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
