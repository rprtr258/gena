package main

import (
	"image"
	"math"

	. "github.com/rprtr258/gena"
)

// https://en.wikipedia.org/wiki/Kite_%28geometry%29
type Kite [4]V2

type Triangle = [3]V2

func newKiteWithBaseSidesAndAngle(point V2, a, b, angle float64) Kite {
	left := Polar(a, angle/2)
	right := Polar(a, -angle/2)
	h := Sqrt(a*a + b*b - 2*a*b*Cos(PI-angle/2-math.Asin(Sin(angle/2)/b*a)))
	return Kite{
		point,
		point + left,
		// point + left + Polar(b, -angle/2),
		point + Polar(h, 0),
		point + right,
	}
}

func (k Kite) splitOnHeight() (left, right Triangle) {
	return Triangle{k[0], k[1], k[2]}, Triangle{k[0], k[2], k[3]}
}

func Compass(
	dc *Context,
	palette Pattern1D,
	c V2,
	numDiamonds int,
	a, b float64,
) {
	dc.SetColor(White)
	dc.Clear()
	dc.SetColor(Black)
	// dc.TransformSet(Rotate(3*PI/2))

	angle := 2 * PI / float64(numDiamonds)

	for i := range numDiamonds {
		triangle0, triangle1 := newKiteWithBaseSidesAndAngle(c, a, b, angle).splitOnHeight()
		phi := float64(i)*angle + 3*PI/2
		rotatedt1 := rotate(triangle0[:], c, phi)
		rotatedt2 := rotate(triangle1[:], c, phi)
		fill(dc, rotatedt1)
		draw(dc, rotatedt1)
		draw(dc, rotatedt2)
	}
}

func compass() *image.RGBA {
	dc := NewContext(Diag(300))
	Compass(dc, Plasma, Diag(150), 5, 50, 100)
	return dc.Image()
}
