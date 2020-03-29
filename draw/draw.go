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

func Square(imd *imdraw.IMDraw, length, x, y, thickness float64) {
	imd.Push(v(x, y))
	imd.Push(v(x, y+length))
	imd.Push(v(x+length, y))
	imd.Push(v(x+length, y+length))
	imd.Rectangle(thickness)
}

func O(imd *imdraw.IMDraw, c vector, radius, thickness float64, color color.Color) {
	imd.Color = color
	imd.Push(c)
	imd.Circle(radius, thickness)
}

func X(imd *imdraw.IMDraw, c vector, length, thickness float64, color color.Color) {
	imd.Color = color
	l := length / 2
	imd.Push(v(c.X-l, c.Y-l), v(c.X+l, c.Y+l))
	imd.Line(thickness)
	imd.Push(v(c.X+l, c.Y-l), v(c.X-l, c.Y+l))
	imd.Line(thickness)
}

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
